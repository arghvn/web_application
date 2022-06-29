package main

import (
	"net/http"
	"text/template"

)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        w.Header().Add("Content Type", "text/html")
		tmpl, err := template.New("test").pars(doc)
		if err == nil {
			context := COntext{
				[3]string{"Lemon", "Orange", "Apple"},
				"the title",
			}
			tmlp.Execute(w, context)
		}
// header for sending html
})
     http.ListenAndServe(":8000", nil)
}

const doc =  
// create doc type and routes
// static content
<!DOCTYPE html>
<html>
  <head><title>{{.title}}</title></head>
  <body>
  <h1>List of fruit</h1>
  <u1>
  {{range .Fruit}}
  <li>{{.}}</li>
  {{end}}
  </body>
</html>

// if we run it , we have Hello templates in localhost
// dynamic content 
// use {{.}} instead templates above and changing nil in error handling
// if we run , we have Hello/ in localhost
// if we after run search localhost:8000/GO we wil have Hello/GO

type context struct {
	Message string
}
// if we run we have Hello the message in localhost
// note eq means equal and ne means not equal
// for example eq 1 (0+1) (2-1) in result true
// gt and le are used for compare , gt means greater  than and le means less or equal

 type context struct {
	fruit [3]string
	title string
 }