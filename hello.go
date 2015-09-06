package main

import (
	"fmt"
	"net/http"
	"github.com/mmorrell/util"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi theree swag!, I love %s!", r.URL.Path[1:])
	util.Tester()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8082", nil)
}