package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func TimeHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t := time.Now()
		fmt.Fprintf(w, "%d%s%d", t.Hour(), "h", t.Minute())
	}
}
func entriesHandlerPost(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(w, "Something went bad")
			return
		}

		saveFile, err := os.OpenFile("./save.data", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		defer saveFile.Close()

		openfile := bufio.NewWriter(saveFile)
		if err == nil {
			fmt.Fprintf(openfile, "%s:%s\n", req.PostForm["author"][0], req.PostForm["entry"][0])
		}
		openfile.Flush()

		fmt.Fprintf(w, "%s: %s\n", req.PostForm["author"][0], req.PostForm["entry"][0])

		return
	}
}
func entriesHandlerGet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		file, err := os.Open("save.data")
		if err != nil {
			fmt.Println(err)
			fmt.Fprintln(w, "Something went bad")
			return
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			a := strings.FieldsFunc(scanner.Text(), Split)
			fmt.Fprintf(w, "%s\n", a[1])
		}

	}
}

func Split(r rune) bool {
	return r == ':' || r == '.'
}

func main() {
	http.HandleFunc("/", TimeHandler)
	http.HandleFunc("/add", entriesHandlerPost)
	http.HandleFunc("/entries", entriesHandlerGet)
	http.ListenAndServe(":4567", nil)
}
