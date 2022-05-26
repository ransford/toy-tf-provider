package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var fooVal = 3

func getFoo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "{\"foo\": %d}\n", fooVal)
}

func putFoo(w http.ResponseWriter, req *http.Request) {
	_newFoo := req.FormValue("foo")
	var newFoo int
	var err error
	if newFoo, err = strconv.Atoi(_newFoo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fooVal = newFoo
	getFoo(w, req)
}

func fooer(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getFoo(w, req)
	case "PUT":
		putFoo(w, req)
	}
}

func main() {
	http.HandleFunc("/foo", fooer)
	http.ListenAndServe(":8090", nil)
}
