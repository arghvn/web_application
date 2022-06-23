package main

// lemonade application for buy juices
// creating a resource server for our application

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// w is an object that implements http.ResponseWriter , instead using object in req we use pointer to http.Request
		w.Write([]byte("Hello , im writting a shop"))
		// its take two argument ,pattern string and handler func(ResponseWriter, *Request), / use for route in project
	})
	http.ListenAndServe(":8000", nil)
}
