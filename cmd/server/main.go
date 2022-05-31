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

func getBar(w http.ResponseWriter, req *http.Request) {
	log.Printf("got bar = %d\n", barVal)
	fmt.Fprintf(w, "{\"bar\": %d}\n", barVal)
}

func putBar(w http.ResponseWriter, req *http.Request) {
	log.Println("in putBar")

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
	}
}

func main() {
	http.HandleFunc("/foo", fooer)
	http.ListenAndServe(":8090", nil)
}
