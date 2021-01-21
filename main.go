package main

import (
	"flag"
	"fmt"
	"net/http"
	"text/template"

	"github.com/inlineboss/fwac/present"
	"github.com/inlineboss/fwac/proxy"

	_ "github.com/go-sql-driver/mysql"
)

var rootDir string
var rootPath string

func homeHandler(w http.ResponseWriter, r *http.Request) {

	prx := proxy.MakeProxy(r, rootDir)

	presenter := present.MakePresenter(prx)

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

func main() {

	root := flag.String("dir", "FAIL", "Root direcory")
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
