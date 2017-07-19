package hello

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{name}", hello)
	http.Handle("/", r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	u := mux.Vars(r)["name"]
	fmt.Fprintf(w, "Hello, %s!", u)
}
