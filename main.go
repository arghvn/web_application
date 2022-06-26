package main

// lemonade application for buy juices
// creating a resource server for our application

import (
	"io/ioutil"
	"net/http"
	"strings"
	"os"
	"bufio"
)

func main() {
	http.Handle("/", new(MyHandler))
	// w is an object that implements http.ResponseWriter , instead using object in req we use pointer to http.Request
	// instead deleted command we use MyHandler struct
	// its take two argument ,pattern string and handler func(ResponseWriter, *Request), / use for route in project

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
	f, err == os.Open(path)
	// at first data and err have an erro that says this two variable declare but not used , for handle this erro we write the erro handling :
	// if err == nil {
	// 	w.Write(data)
	// } else {
	// 	w.WriteHeader(404)
	// 	w.Write([]byte("404 -" + http.StatusText(404)))

	if err == nil {
		bufferdReader := bufio.NewReader(f)
		var contentType string
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css "
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
			} else if strings.HasSuffix(path, ".mp4") {
		} else {
			contentType = "video/mp4"
		}
		// using headerfor response by headerfunction
		w.Header().Add("content Type", contentType)
		bufferdReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 -" + http.StatusText(404)))
	}
}

// output : 404 - notfound

// if we search localhost:8000/html/home.html in  browser , show a page include login , home , shop , stand locator ... related to site that we are coding

// by adding if/else for css , html , js / png ... and again if we search localhost:8000/html/home.html in  browser ,
// show a page include login , home , shop , stand locator ... related to site that we are coding , wonderful ... this site is userfriendly
