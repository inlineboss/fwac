package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/inlineboss/fwac/present"
	"github.com/inlineboss/fwac/proxy"

	_ "github.com/go-sql-driver/mysql"
)

var rootDir string
var rootPath string

func templateParse(text string) (*template.Template, error) {
	var t *template.Template

	tmp := t.New("N")
	tmp.Parse(text)

	return t, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	prx := proxy.MakeProxy(r, rootDir)

	f, err := os.Stat(prx.FS.LongPath)
	if err == nil && !f.IsDir() {
		file, err := ioutil.ReadFile(prx.FS.LongPath)
		if err == nil {
			w.Write(file)
			return
		}
	}

	presenter := present.MakePresenter(prx)

	// page, err := template.New("Html").Funcs(template.FuncMap{
	// 	"minus": func(a, b int) int {
	// 		return a - b
	// 	},
	// }).Parse(templates.Html)

	page, err := template.New("home.html").Funcs(template.FuncMap{
		"minus": func(a, b int) int {
			return a - b
		},
	}).ParseFiles("templates/home.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = page.Execute(w, presenter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {

	root := flag.String("dir", "/Users/inlineboss/", "Root direcory")
	port := flag.Int("port", 8080, "Port")
	flag.Parse()

	if *root == "FAIL" {
		fmt.Println("Specify root direcory")
		return
	}
	rootDir = *root

	fmt.Println("Root: " + fmt.Sprint(rootDir))
	fmt.Println("Port: " + fmt.Sprint(*port))

	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":"+fmt.Sprint(*port), nil)

}
