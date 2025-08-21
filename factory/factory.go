package factory

import (
	"ecommerce/models"
	"fmt"
)

// ProductFactory type for creating products
type ProductFactory func(id, name string, price float64, stock int, extra interface{}) models.ProductCategory

var productFactories = map[int]ProductFactory{
	1: func(id, name string, price float64, stock int, extra interface{}) models.ProductCategory {
		warranty, ok := extra.(int)
		if !ok {
			warranty = 0
		}
		return models.NewElectronics(id, name, price, stock, warranty)
	},
	2: func(id, name string, price float64, stock int, extra interface{}) models.ProductCategory {
		sizes, ok := extra.([]string)
		if !ok {
			sizes = []string{"M"}
		}
		return models.NewClothing(id, name, price, stock, sizes)
	},
	3: func(id, name string, price float64, stock int, extra interface{}) models.ProductCategory {
		return models.NewBooks(id, name, price, stock)
	},
}

func CreateProduct(typeChoice int, id, name string, price float64, stock int, extra interface{}) models.ProductCategory {
	if factory, ok := productFactories[typeChoice]; ok {
		return factory(id, name, price, stock, extra)
	}
	fmt.Println("Invalid product type")
	return nil
}