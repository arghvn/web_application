package controllers

import (
	"net/http"
	"os"
	"text/template"
	"bufio"
	"strings"
	"viewmodels"
	"github.com/gorilla/mux"
)

func register(templates *template.Template) {
	router := mux.NewRouter()

	hc := new(homecontroller)
	hc.template = templates.Lookup("home.html")
    router.HandleFunc("/home", hc.get)

	cc := new(CategoriesController)
	cc.template = templates.Lookup("categories.html")
	router.HandleFunc("/categories", cc.get)
    categorycontroller := new(categorycontroller)
    categorycontroller.template = templates.Lookup("products.html")
	router.HandleFunc("/categories/{id}", categorycontroller.get)

	http.Handle("/", router)

	http.Handlefunc("/", router)
        func(w http.RespnseWriter, req *http.Request) {
			requestedfile := req.URL.path[1:]
			template :=
			     templates.lookup(requestedfile + ".html")

				 var context interface{} = nil
                 switch requestedfile {

				 case "home":
					context = viewmodles.Gethome()
				 case "categories":
					context : viewmodels.GetCategories()
				 }
				 if template != nil {
					template.Execute(w, context)
				 } else {
					w.writeHeader(404)
				 }
				 }
		})
	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)



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

	f, err := os.Open(path)

	if err == nil {
		defer f.close()
		w.Header().Add("content type", contenttpye)
	}
	
// make use controller package in package main by import

// in output(ru as application) every thing is fine 
// for example in lemon juice we see the details
// and other choices



