package controller

import (
	"net/http"
	"strconv"
	"text/template"
	"viewmodels"

	"github.com/gorilla/mux"
)

type product struct {
	template *template.Template
}

func (this *product) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	idRaw := vars["id"]

	id, err := strconv.Atoi(idRaw)
	if err == nil {
		vm := viewmodels.Getproduct(id)
		w.Header().Add("content type", "text/html")
		this.template.Execute(w, vm)
	} else {
		w.WriteHeader(404)
	}
}
