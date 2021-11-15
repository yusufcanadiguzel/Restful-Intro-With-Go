package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	CategoryId  int     `json:"categoryId"`
	ProductName string  `json:"productName"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() {
	response, err := http.Get("http://localhost:3000/products")

	if err != nil {
		fmt.Println("Bir sorun oluştu: ", err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	fmt.Println(products)
}

func GetAllCategories()  {
	response, err := http.Get("http://localhost:3000/categories")

	if err != nil {
		fmt.Println("Bir sorun oluştu: ", err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var categories []Category
	json.Unmarshal(bodyBytes, &categories)
	fmt.Println(categories)
}

func AddProduct()  {
	var product Product = Product{CategoryId: 2, ProductName: "Jean", UnitPrice: 150.0}

	jsonProduct, _ := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products","application/json; charset = utf-8", bytes.NewBuffer(jsonProduct))

	if err != nil {
		fmt.Println("Bir sorun oluştu :", err)
	}

	defer response.Body.Close()
}

// ---------- JSON PLACE HOLDER MODEL ----------
// type ToDo struct {
// 	UserId    int    `json:"userId"`
// 	Id        int    `json:"id"`
// 	Title     string `json:"title"`
// 	Completed bool   `json:"completed"`
// }

// ---------- JSON PLACE HOLDER GET FUNC ----------
// func GetValue() {
// 	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

// 	if err != nil {
// 		fmt.Println("Bir sorun oluştu: ",err.Error())
// 	}

// 	defer response.Body.Close()

// 	bodyBytes, _ := ioutil.ReadAll(response.Body)

// 	//String değer okumak için
// 	stringBody := string(bodyBytes)
// 	fmt.Println(stringBody)

// 	//Modele uygun şekilde okumak için
// 	var todo ToDo
// 	json.Unmarshal(bodyBytes, &todo)
// 	fmt.Println(todo)
// }

// ---------- JSON PLACE HOLDER SET FUNC ----------
// func SetValue()  {

// 	var todo ToDo = ToDo{UserId: 1, Id: 1, Title: "Kod yazılıcak.", Completed: true}
// 	jsonTodo, _ := json.Marshal(todo)

// 	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset = utf-8", bytes.NewBuffer(jsonTodo))

// 	if err != nil {
// 		fmt.Println("Bir hata oluştu :", err.Error())
// 	}

// 	defer response.Body.Close()

// 	bodyBytes, _ := ioutil.ReadAll(response.Body)

// 	//String
// 	stringBody := string(bodyBytes)
// 	fmt.Println(stringBody)

// 	//Json
// 	var todoResponse ToDo
// 	json.Unmarshal(bodyBytes, &todoResponse)
// 	fmt.Println(todoResponse)
// }