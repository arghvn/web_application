package viewmodels

import (

)

type Home struct {
	Title string
	Active string
}

func GetHome|() Home {
	result := Home {
		Title: "Lemonda stand society",
		Active: "home",
	}
	return result
}

// we created viewmodels package and we can use it in main.go
// first we have to import it