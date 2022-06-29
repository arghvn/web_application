package main

import (
	"html/template"
	"net/http"
)

func main() {
	templates := populateTemplates()

	http.HandlerFunc("/",
	    func(w http.ResponseWriter, req *http.Request){
			requestedFile := req.URL.path[1:]
			template := 
			templates.Lookup(requstedFile + ".html")
			if template != nil {
				template.Execute(w, nil)
			}else
			    w.WriteHeader(404)
		}
 )
    http.ListenAndServe(":8000", nil)
	http.ListenAndServe(":8000", http.FileServer(http.Dir("pub")))
}

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templatefolder, _ := os.open(basepath)
	defer templateFolder.close()

	templatepathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := mew([]string)
	for _, pathinfo := range templatePathRaw {
		if !pathinfo.IsDir() {
			*templatePaths = append(*templatePaths, 
			basePath + "/" + pathinfo.Name())
			
		}
	}
}

result.ParseFiles(*templatePaths...)

return result


// run :
// localhost:8000
//