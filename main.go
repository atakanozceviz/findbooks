package main

import (
	"html"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/atakanozceviz/findbooks/controller"
	"github.com/atakanozceviz/findbooks/model"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", searchKeyword)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./view/favicon.ico")
	})
	log.Println("Serving on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func searchKeyword(w http.ResponseWriter, r *http.Request) {
	k := html.EscapeString(r.FormValue("keyword"))
	if k != "" {
		var books model.Books
		w.Header().Set("Content-Type:", "application/json;charset=utf-8")
		w.Write(controller.Search(&books, k).ToJson())
	}
}
