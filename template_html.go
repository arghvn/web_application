package main

import (
	"net/http"
	"text/template"

)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        w.Header().Add("Content Type", "text/html")
		templates := template.New("template")
		templates.New("test").parse(doc)
		templates.New("header").parse(header)
		templates.New("footer").parse(footer)
			context := COntext{
				[3]string{"Lemon", "Orange", "Apple"},
				"the title",
			}
			tmlpaltes.Lookup.Execute(w, context)
		}
// header for sending html
})
     http.ListenAndServe(":8000", nil)
}

const doc =  
// create doc type and routes
// static content
{{template "header" .title}} 
  <body>
  <h1>List of fruit</h1>
  <u1>
  {{range .Fruit}}
  <li>{{.}}</li>
  {{end}}
  </body>
  {{template "footer"}}




const header = 
<!DOCTYPE html>
<html>
  <head><title>{{.}}</title></head>


  const footer = 
  </html>


type context struct {
	Message string
}
 type context struct {
	fruit [3]string
	title string
 }


 // output in this section by searching localhost:8000 :
 // list of friut :
 // Lemon
 // Orange
 // Apple