package models

import "github.com/sohamkamani/birdpedia/summary/entities"

type productModel struct {
}

func (*productModel) Finall() ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select * from product")
		if err != nil {
			return nil, err2
		} else { 
			var product []entities.Product
			var product entities.Product
			for rows.Next() {
				rows.scan(&product.Id, &product.Name, &product.price, &product.Quantity, &product.Description)
				products = append(products, product)
			}
			return products, nil
	}
} 