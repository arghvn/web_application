package controllers

import (
	"controllers/util"
	"net/http"
	"text/template"
	"viewmodels"
)

// the package util made in controllers folder

type homeController struct {
	template *template.Template
}

func (this *homeController) get(w http.RespnseWriter, req *http.request) {
	vm := viewmodels.GetHome()

	w.Header().Add("content Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	// defer in GO ?
	this.template.Execute(responseWriter, vm)
}
