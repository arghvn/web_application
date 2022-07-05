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

// Adding the bird REST API handlers
// As we can see, we will need to implement two APIs in order for this application to work:

// GET /bird - that will fetch the list of all birds currently in the system
// POST /bird - that will add an entry to our existing list of birds
// For this, we will write the corresponding handlers.

// Create a new file called bird_handlers.go, adjacent to the main.go file.

// First, we will add the definition of the Bird struct and initialize a common bird variable:

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird
// Next, define the handler to get all birds :

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	//Convert the "birds" variable to json
	birdListBytes, err := json.Marshal(birds)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of birds to the response
	w.Write(birdListBytes)
}
Next, the handler to create a new entry of birds :

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Bird
	bird := Bird{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the bird from the form info
	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	// Append our existing list of birds with a new entry
	birds = append(birds, bird)

	//Finally, we redirect the user to the original HTMl page
	// (located at `/assets/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
The last step, is to add these handler to our router, in order to enable them to be used by our application :

	// These lines are added inside the newRouter() function before returning r
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r



