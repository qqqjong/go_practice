package main

import (
	"github.com/qqqjong/go_practice/src/ex25/app"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}