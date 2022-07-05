package main

// Starting an HTTP server

import (
	"fmt"
	"net/http"
	// fmt" has methods for formatted I/O operations (like printing to the console
	// The "net/http" library has methods to implement HTTP clients and servers
)

func main() {
	// The "HandleFunc" method accepts a path and a function as arguments
	// we can pass functions as arguments, and even treat them like variables in Go

	http.HandleFunc("/", handler)
	// After defining our server, we finally "listen and serve" on port 8080  The second
	// argument is the handler, which we will come to later on, but for now it is left as nil,
	// and the handler defined above (in "HandleFunc") is used.

	http.ListenAndServe(":8080", nil)
}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World!")
}

// fmt.Fprintf, unlike the other “printf” statements , takes a “writer” as its first argument.
// The second argument is the data that is piped into this writer.
// The output therefore appears according to where the writer moves it.
// In our case the ResponseWriter w writes the output as the response to the users request.

// You can now run this file / go run main.go and navigate to http://localhost:8080 in your browser, or by running the command :

// curl localhost:8080
// And see the output: “Hello World!”

// You have now successfully started an HTTP server in Go.

// Making routes :
// Our server is now running, but, you might notice that we get the same “Hello World!” response regardless of the route we hit,
//  or the HTTP method that we use.
//  To see this yourself, run the following curl commands, and observe the response that the server gives you :
// curl localhost:8080/some-other-route
// curl -X POST localhost:8080
// curl -X PUT localhost:8080/samething
// All three commands still give you “Hello World!”

// We would like to give our server a little more intelligence than this,
//  so that we can handle a variety of paths and methods. This is where routing comes into play.

// Although you can achieve this with Go’s net/http standard library,
//  there are other libraries out there that provide a more idiomatic and declarative way to handle http routing.

// Installing external libraries :
// We will be installing a few external libraries through this tutorial, where the standard libraries don’t provide the features that we want.
// When we install libraries, we need a way to ensure that other people who work on our code also have the same version of the library that we do.

// In order to do this, we can make use of the go get command,
//  which helps us install the library of our choice, and adds its version information to the go.mod and go.sum files.

// Let’s install our routing library:

// go get github.com/gorilla/mux
// This will add the gorilla/mux library to your project.

// Now, we can modify our code to make use of the functionality that this library provides .
