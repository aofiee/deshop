package main

import (
	"fmt"
	"net/http"

	"github.com/aofiee/deshop/discount"
	"github.com/gorilla/mux"
)

func main() {
	handleRequest()
}

func handleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/discounts/{country}", gameList)
	http.Handle("/", r)
	http.ListenAndServe(":1234", nil)
}

func gameList(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	response := discount.GetDiscountGameFrom(discount.API[param["country"]])
	fmt.Fprint(w, string(response))
}
