package main

import (
	"fmt"
	"net/http"

	"github.com/aofiee/deshop/eshop"
	"github.com/gorilla/mux"
)

func main() {
	handleRequest()
}

func handleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/discounts/{country}", gameList)
	r.HandleFunc("/new-release/{country}", newRelease)
	r.HandleFunc("/ranking/{country}", newRelease)
	http.Handle("/", r)
	http.ListenAndServe(":1234", nil)
}

func newRelease(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	response := eshop.GetEndpoint(param["country"], `new`)
	fmt.Fprint(w, string(response))
}

func gameList(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	response := eshop.GetEndpoint(param["country"], `sale`)
	fmt.Fprint(w, string(response))
}

func rankingList(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	response := eshop.GetEndpoint(param["country"], `ranking`)
	fmt.Fprint(w, string(response))
}
