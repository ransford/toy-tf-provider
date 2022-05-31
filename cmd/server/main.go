package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var barVal = 3

func getBar(w http.ResponseWriter, req *http.Request) {
	log.Printf("got bar = %d\n", barVal)
	fmt.Fprintf(w, "{\"bar\": %d}\n", barVal)
}

func putBar(w http.ResponseWriter, req *http.Request) {
	_newFoo := req.FormValue("foo")
	var newBar int
	var err error
	if newBar, err = strconv.Atoi(_newFoo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	barVal = newBar
	log.Printf("stored bar = %d\n", barVal)
	getBar(w, req)
}

func fooer(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getBar(w, req)
	case "PUT":
		putBar(w, req)
	}
}

func main() {
	http.HandleFunc("/foo", fooer)
	http.ListenAndServe(":8090", nil)
}
