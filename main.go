package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/lazy", LazyHandler)
	r.HandleFunc("/", RootHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "Category: %v\n", vars["category"])
	var tmplFile = "rsc/index.html.tpl"
	tmpl, err := template.New("index.html.tpl").ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	inputs := map[string]string{
		"timestamp": time.Now().String(),
	}
	err = tmpl.Execute(w, inputs)
	if err != nil {
		panic(err)
	}
}

func LazyHandler(w http.ResponseWriter, r *http.Request) {
	println("triggered lazy stuff")
	fmt.Fprint(w, "<p>hello world</p>")
    w.WriteHeader(http.StatusOK)
}
