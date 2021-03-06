package viewmodels

import (

)

type categories struct {
	Title string 
	Active string
	Categories []category
}

type category struct {
	ImageUrl string
	Title string 
	Description string
	IsOrientRight bool
	Id int
}

func GetCategories() categories {
	result := Categories{
		Title: "Lemonada stand society - shop",
		Active: "shop",
 	}
}

// fpr example for juice category :
juiceCategory := Category{
	Id: 1,
	ImageUrl: "lemon.png",
	Title: "juice and Mixes",
	Description: "a short brief",

IsOrientRight: false,

// IsOrientRight use for category in GO
}

supplycategory := category{
	Id: 2,
	ImageUrl: "Kiwi.png",
	Title: "cups, strawa, and other supplies",
	Description: "a short brief",
	IsOrientRight: true,
}

advertiseCategory := category{
	Id: 3,
	ImageUrl: "pineapple.png",
	Title: " signs and advertising",
	Description: "a short brief"
	IsOrientRight: false,
}

result.categories = []category{
	juiceCategory,
	supplycategory,
	advertiseCategory
}

return result