package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var barVal = 13

type FooThing struct {
	Bar int `json:"bar"`
}

// getBar handles a GET request by returning the value of barVal in JSON format.
func getBar(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "{\"bar\": %d}\n", barVal)
}

// putBar handles a PUT request by overwriting barVal with the provided value.
func putBar(w http.ResponseWriter, req *http.Request) {
	var ft FooThing
	err := json.NewDecoder(req.Body).Decode(&ft)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	barVal = ft.Bar
	log.Printf("stored bar = %d\n", barVal)
	getBar(w, req)
}

func fooer(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getBar(w, req)
	case "PUT":
		putBar(w, req)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	http.HandleFunc("/foo", fooer)
	http.ListenAndServe(":8090", nil)
}
