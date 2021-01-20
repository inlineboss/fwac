package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/inlineboss/fwac/fs"
	"github.com/inlineboss/fwac/url"
)

var rootDir string
var rootPath string

type HTMLObject struct {
	URL struct {
		Road    url.Details
		Return  string
		Current string
	}

	Elems []fs.Element
}

func (htmlObjects *HTMLObject) resetPaths() {
	if (htmlObjects.URL.Current == rootDir) && (htmlObjects.URL.Return == rootDir) {
		return
	}

	htmlObjects.URL.Current = rootDir
	htmlObjects.URL.Return = rootDir
	htmlObjects.Elems = fs.ShowDir(rootDir)
	htmlObjects.URL.Road = 
}

func (htmlObjects *HTMLObject) setPath(path string) {
	if htmlObjects.URL.Current == path {
		return
	}

	htmlObjects.URL.Current = path
	htmlObjects.Elems = fs.ShowDir(path)

	d, _ := url.ExtractLast(path)
	htmlObjects.URL.Return = d.Path

	str := strings.TrimLeft(htmlObjects.URL.Current, rootDir)
	htmlObjects.URL.Road = url.ExtractLasts(str)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	var webObject HTMLObject

	if (r.URL.Path == "/") || (r.URL.Path == rootDir) {
		webObject.resetPaths()
	} else {
		webObject.setPath(r.URL.Path)
	}

	page, err := template.ParseFiles("templates/home.html")
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = page.Execute(w, webObject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about")
}

func setHandlers() {
	http.HandleFunc("/", homeHandler)
}

func cherr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	name := flag.String("d", "/Users/inlineboss", "specify root direcory")
	port := flag.Int("p", 8080, "specify port of receive")

	flag.Parse()

	fmt.Println("Root: " + fmt.Sprint(*name))
	fmt.Println("Port: " + fmt.Sprint(*port))

	rootDir = *name

	setHandlers()
	http.ListenAndServe(":"+fmt.Sprint(*port), nil)

}
