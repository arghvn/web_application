package productcontroller

import "net/http"

func Index(responsehttp.ResponseWriter, request *http.Request) {
	tmp,_ := template.parsefile("views/product/index.html")
	tmp.Execute(response)
}

// keep on develop after learning  database