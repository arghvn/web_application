package controllers

import (
	"net/http"
	"text/template"
	"viewmodels"
	"github.com/gorilla/mux"
	"strconv"
)

type CategoriesController struct {
	template *template.Template
}

func (this *CategoriesController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.vars(req)

	idRaw := vars["id"]

	id, err := strconv.Atoi(idRaw)
	if err == nil {

	
	vm := viewmodels.GetCategories(id)

	w.Header().Add("content Type", "text/html")
	this.template.Execute(w, vm)
} else {
	w.WriteHeader(404)
}
}

type CategoriesController struct {
	template *template.Template
}

func (this *CategoriesController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.Getproducts(!!!!!!!!)

	w.Header().Add("Content Type", "text/html")
	this.template.Execute(w, vm)
}

 // gorilla mux instalation
 // go get -u github.com/gorilla/mux
 // -u for last update