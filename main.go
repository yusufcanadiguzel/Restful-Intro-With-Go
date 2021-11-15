package main

import (
	"fmt"
	"restfulIntro/restful"
)

func main() {

	//Ürünleri listele
	products, _ := restful.GetAllProducts()

	for _, product := range products {
		fmt.Println(product.ProductName)
	}

	//Kategorileri listele
	categories, _ := restful.GetAllCategories()

	for _, category := range categories {
		fmt.Println(category.CategoryName)
	}

	//Ürün ekleme
	restful.AddProduct()
}
