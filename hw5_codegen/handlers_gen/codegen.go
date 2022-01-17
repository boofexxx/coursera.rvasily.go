package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

func FatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Api struct {
	URL    string `json:"url"`
	Auth   bool   `json:"auth"`
	Method string `json:"method"`
}

func (a Api) String() string {
	return fmt.Sprintf("URL: %s, Auth: %t, Method: %s", a.URL, a.Auth, a.Method)
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("usage: %s [input] [output]", os.Args[0])
	}
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
	FatalIfError(err)

	out, err := os.Create(os.Args[2])
	FatalIfError(err)
	defer out.Close()

	fmt.Fprintln(out, `package `+node.Name.Name)
	fmt.Fprintln(out)
	fmt.Fprintln(out, `import (`)
	// all the imports are here
	fmt.Fprintln(out, `)`)
	fmt.Fprintln(out)

	for _, n := range node.Decls {
		mf, ok := n.(*ast.FuncDecl)
		if !ok {
			log.Printf("SKIP %#T is not *ast.FunDecl\n", n)
			continue
		}
		// continue if the function declaration is not a method
		// or there is no comment
		if mf.Recv == nil {
			log.Printf("SKIP %T is not a method\n", mf)
			continue
		} else if mf.Doc == nil {
			log.Printf("SKIP %T doesn't have a comment\n", mf)
			continue
		}

		api := Api{}
		needCodegen := false
		for _, comment := range mf.Doc.List {
			needCodegen = needCodegen || strings.HasPrefix(comment.Text, "// apigen:api")
			if needCodegen {
				err := json.Unmarshal([]byte(comment.Text[len("// apigen:api"):]), &api)
				if err != nil {
					log.Printf("SKIP comment of %s: %s", mf.Name, err)
					continue
				}
				break
			}
		}
		if !needCodegen {
			log.Printf("SKIP method %#v, it doesnt have apigen mark\n", mf.Name)
			continue
		}

		// process comment
		fmt.Fprintln(out, api)
		recvType := ""
		recvName := ""
		for _, v := range mf.Recv.List {
			switch xv := v.Type.(type) {
			case *ast.StarExpr:
				recvType = "*"
				if si, ok := xv.X.(*ast.Ident); ok {
					recvType += si.Name
					recvName += si.Obj.Name
				}
			case *ast.Ident:
				recvType = xv.Name
				recvName = xv.Obj.Name
			}
		}
		if typeTable[recvType] {
			continue
		}
		typeTable[recvType] = true
		fmt.Fprintf(out, "func (%s %s) ServeHTTP(w http.ResponseWriter, r *http.Request) {\n", recvName, recvType)
		fmt.Fprintln(out, "}")
	}
}

var typeTable map[string]bool = make(map[string]bool)
