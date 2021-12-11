package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var output = flag.String("o", "", "file where you want the app to out")
var printFiles = flag.Bool("f", false, "make it true if you want print files as well")
var nestedLevel = flag.Int("n", 100, "how deep to go")

func main() {
	log.SetFlags(0)
	flag.Parse()
	out := os.Stdout
	var err error

	if *output != "" {
		out, err = os.Create(*output)
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	for _, path := range flag.Args() {
		fmt.Fprintf(out, "%s:\n", path)
		err = dirTree(out, path, *printFiles)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func dirTree(out io.Writer, path string, printFiles bool) (err error) {
	if len(path) > 2 {
		if path[0] == '.' && path[1] == os.PathSeparator {
			path = path[2:]
		}
	}
	if path[len(path)-1] == os.PathSeparator {
		path = path[:len(path)-1]
	}
	var r string
	r, err = formTree(path, path, printFiles, make(map[int]bool))
	fmt.Fprint(out, r)
	return err
}

// formTree returns string =3
func formTree(root, path string, printFiles bool, levels map[int]bool) (string, error) {
	if root != path {
		if findPathDepth(path[len(root)+1:])+2 > *nestedLevel {
			return "", nil
		}
	}
	s := ""
	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return s, err
	}
	var dirs []fs.FileInfo
	if !printFiles {
		for _, fi := range fis {
			if fi.IsDir() {
				dirs = append(dirs, fi)
			}
		}
		fis = dirs
	}
	graph := "├───"
	for i, fi := range fis {
		path := path + string(os.PathSeparator) + fi.Name()
		levels[findPathDepth(path[len(root)+1:])] = true
		var sep string
		for k, j := 0, findPathDepth(path[len(root)+1:]); k < j; k++ {
			if levels[k] {
				sep += "│\t"
			} else {
				sep += "\t"
			}
		}

		if i == len(fis)-1 {
			levels[findPathDepth(path[len(root)+1:])] = false
			graph = "└───"
		}
		if fi.IsDir() {
			s += fmt.Sprintf("%s%s%s\n", sep, graph, fi.Name())
			t, err := formTree(root, path, printFiles, levels)
			s += t
			if err != nil {
				// I should have returned it but I don't really know how to continue to work from here returning an error
				log.Print(err)
				continue
			}
		} else {
			s += fmt.Sprintf("%s%s%s (%s)\n", sep, graph, fi.Name(), formatSize(fi.Size()))
		}
	}
	return s, nil
}

// findPathDepth returns the number of os.PathSeparator in path
func findPathDepth(path string) (n int) {
	for _, r := range path {
		if r == os.PathSeparator {
			n++
		}
	}
	return
}

// formatSize formats the int64 to string size
func formatSize(s int64) string {
	if s == 0 {
		return "empty"
	}
	return strconv.FormatInt(s, 10) + "b"
}
