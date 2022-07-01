package controllers

import (
	"net/http"
	"text/template"
	"viewmodels"
)

type homeController struct {
	template *template.Template
}

func (this *homeController) get(w http.RespnseWriter, req *http.request) {
	vm := viewmodels.GetHome()

	w.Header().Add("content Type", "text/html")
	this.template.Execute(w, vm)
}
