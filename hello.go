package hello

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"html/template"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/users/{name}", hello)
	r.HandleFunc("/", index)
	http.Handle("/", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	if err := indexTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var indexTemplate = template.Must(template.ParseFiles("templates/index.html"))

func hello(w http.ResponseWriter, r *http.Request) {
	u := mux.Vars(r)["name"]
	fmt.Fprintf(w, "Hello, %s!", u)
}
