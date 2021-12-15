package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/valyala/fastjson"
)

// type User struct {
// 	Browsers []string
// 	Company  string
// 	Country  string
// 	Email    string
// 	Job      string
// 	Name     string
// 	Phone    string
// }

var Android = []byte("Android")
var MSIE = []byte("MSIE")

func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	var seenBrowsers [][]byte
	uniqueBrowsers := 0
	foundUsers := ""

	scanner := bufio.NewScanner(file)
	var p fastjson.Parser
	i := 0
	for scanner.Scan() {
		i++
		text := scanner.Bytes()
		user, err := p.ParseBytes(text)
		if err != nil {
			panic(err)
		}

		isAndroid := false
		isMSIE := false
		browsers := user.GetArray("browsers")
		if browsers == nil {
			// log.Println("cant cast browsers")
			continue
		}

		for _, browserRaw := range browsers {
			browser, err := browserRaw.StringBytes()
			if err != nil {
				// log.Println("cant cast browser to string")
				continue
			}
			if bytes.Contains(browser, Android) {
				isAndroid = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if bytes.Equal(item, browser) {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					buffer := make([]byte, len(browser))
					copy(buffer, browser)
					seenBrowsers = append(seenBrowsers, buffer)
					uniqueBrowsers++
				}
			}
			if bytes.Contains(browser, MSIE) {
				isMSIE = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if bytes.Equal(item, browser) {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					buffer := make([]byte, len(browser))
					copy(buffer, browser)
					seenBrowsers = append(seenBrowsers, buffer)
					uniqueBrowsers++
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		// I think we could replace it to bytes either. but I don't really think it would reduce number of allocations 
		email := strings.Replace(string(user.GetStringBytes("email")), "@", " [at] ", -1)
		foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i-1, string(user.GetStringBytes("name")), email)
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}