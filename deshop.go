package main

import (
	"fmt"
	"net/http"

	"github.com/aofiee/deshop/discount"
)

func main() {
	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/", gameList)
	http.ListenAndServe(":1234", nil)
}

func gameList(w http.ResponseWriter, r *http.Request) {
	response := discount.GetDiscountGameFrom(discount.NZ)
	fmt.Fprint(w, string(response))
}
