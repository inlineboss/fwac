package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/inlineboss/fwac/fs"
)

var ROOT_DIR string

type DetailDir struct {
	Name string
	Path string
}

func ExtractDirs(path string) []DetailDir {

	if len(path) == 0 {
		return nil
	}

	var details []DetailDir
	var name string

	for i := len(strings.Split(path, "/")) - 1; i > 0; i-- {
		name, path = ExtractLastDir(path)

		if name != "" || path != "" {
			details = append(details, DetailDir{Name: name, Path: path})
		}
	}

	if details == nil {
		return append(details, DetailDir{Name: path, Path: "/"})
	}

	return details
}

func ExtractLastDir(path string) (n string, p string) {
	if len(path) == 0 || path == "/" {
		return "", ""
	}

	lastPos := strings.LastIndexByte(path, '/')
	if lastPos == -1 || lastPos == 0 {
		return "", ""
	}

	n = path[lastPos+1:]
	p = path[0:lastPos]

	return n, p
}

type HTMLObject struct {
	URL struct {
		Road    []DetailDir
		Return  string
		Current string
	}

	Elems []fs.Element
}

func (htmlObjects *HTMLObject) resetPaths() {
	if (htmlObjects.URL.Current == ROOT_DIR) && (htmlObjects.URL.Return == ROOT_DIR) {
		return
	}

	htmlObjects.URL.Current = ROOT_DIR
	htmlObjects.URL.Return = ROOT_DIR
	htmlObjects.Elems = fs.ShowDir(ROOT_DIR)
	htmlObjects.URL.Road = []DetailDir{
		DetailDir{"/", ROOT_DIR},
	}
}

func (htmlObjects *HTMLObject) setPath(path string) {
	if htmlObjects.URL.Current == path {
		return
	}

	htmlObjects.URL.Current = path
	htmlObjects.Elems = fs.ShowDir(path)

	_, htmlObjects.URL.Return = ExtractLastDir(path)

	str := strings.TrimLeft(htmlObjects.URL.Current, ROOT_DIR)
	http.
		htmlObjects.URL.Road = ExtractDirs(str)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	var webObject HTMLObject

	if (r.URL.Path == "/") || (r.URL.Path == ROOT_DIR) {
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
	http.HandleFunc("/about/", aboutHandler)
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

	ROOT_DIR = *name

	setHandlers()
	http.ListenAndServe(":"+fmt.Sprint(*port), nil)

}
