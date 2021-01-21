package main

import (
	"flag"
	"fmt"
	"net/http"
	"text/template"

	"github.com/inlineboss/fwac/html"
	"github.com/inlineboss/fwac/proxy"

	_ "github.com/go-sql-driver/mysql"
)

var rootDir string
var rootPath string

func homeHandler(w http.ResponseWriter, r *http.Request) {

	prx := proxy.MakeProxy(r, rootDir)

	presenter := html.MakeWEBPresenter(prx)

	page, err := template.ParseFiles("templates/home.html")

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

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about")
}

func setHandlers() {
	http.HandleFunc("/", homeHandler)
}

func main() {

	rootDir = *flag.String("d", "/Users/inlineboss", "specify root direcory")
	port := flag.Int("p", 8080, "specify port of receive")

	flag.Parse()

	fmt.Println("Root: " + fmt.Sprint(rootDir))
	fmt.Println("Port: " + fmt.Sprint(*port))

	setHandlers()
	http.ListenAndServe(":"+fmt.Sprint(*port), nil)

}
