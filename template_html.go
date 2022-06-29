package main

import (
	"net/http"
	"text/template"
	"bufio"
	"strings"

)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        w.Header().Add("Content Type", "text/html")
		templates := template.New("template")
		templates.New("test").parse(doc)
		templates.New("header").parse(header)
		templates.New("footer").parse(footer)
			context := Context{
				[3]string{"Lemon", "Orange", "Apple"},
				"the title",
			}
			tmlpaltes.Lookup.Execute(w, context)
		}
// header for sending html
})
     http.ListenAndServe(":8000", nil)
}

func serveResource(w http.ResponseWriter, req *http.Request)
     path := "public" + req.URL.Path
	 var contentType string
	 if strings.HasSuffix(path, ".css"){
		contentType = "text/css"
	 } else if strings.HasSuffix(path, ".png"){
		else { 
			contentType = "text/plain"
			// text/plain use for captcha
		}
	 }



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