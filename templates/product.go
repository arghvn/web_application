package main

import (
	"viewmodels"
)

type products struct {
	title string
	active string
	products []products
}

func Getproducts(id int) products {
	var result products
	result.active = "shop"
	var shopname string
	switch id {
	case 1:
		shopname = "juice"
	case 2:
		shopname = "supply"
	case 3:
		shopname = "advertising-"

	}
	result.title = "lemon stand society -" + shopname + " shop"
    
	if id == 1 {
		result.products = getproductlist()
	
	lemonjuice := makelemonjuiceproduct()
	applejuice := makeapplejuiceproduct()
	watermelonjuice := makewatermelonjuiceproduct()
	kiwijuice := makekiwijuiceproduct()
	mangosteenjuice := makemangosteenjuice()
	orangejuice := makeorangejuiceproduct()
	pineapplejuice := makepineapplejuice()
	strawberryjuice := makestrawberryjuice()

	// result := []product{}
	result.products = []product{
		lemonjuice,
		applejuice,
		watermelonjuice,
		kiwijuice,
		mangosteenjuice,
		orangejuice,
		pineapplejuice,
		strawberryjuice,
	}

	return result
}

func Getproductlist() []product {

}

type productVM struct {
	title string
	active string
	productVM product
}

func Getproduct(id int) productVM {
	var result productVM
	productlist := getproductlist()
	var product product
	for _, p := range productlist {
		if p.id == id {
			product = id,
				break;
			}
		}
	}
}
