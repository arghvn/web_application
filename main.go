package main

// lemonade application for buy juices
// creating a resource server for our application

import (
	"io/ioutil"
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

type MyHandler struct {
	http.Handler
}

// HTTP handler functions in Go must have a very particular signature:
// func(http.ResponseWriter, *http.Request).
//  The web server implementation in the standard library needs to call these functions, so they must conform to a type declared in the http package, and that type only allows for those two parameters.
func (this *MyHandler) SrveHttp(w http.ResponseWriter, req *http.Request) {
	path := "public/" + req.URL.Path
	data, err := ioutil.ReadFile(string(path))
	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 -" + http.StatusText(404)))
	}
}
