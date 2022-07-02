package controllers

import (
	"net/http"
	"text/template"
	"viewmodels"
)

type CategoriesController struct {
	template *template.Template
}

func (this *CategoriesController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetCategories()

	w.Header().Add("content Type", "text/html")
	this.template.Execute(w, vm)
}


type CategoriesController struct {
	template *template.Template
}

func (this *CategoriesController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.Getproducts(!!!!!!!!)
}