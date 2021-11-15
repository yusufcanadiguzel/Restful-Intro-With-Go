package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ToDo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetValue() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Bir sorun oluştu: ",err.Error())
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	
	//String değer okumak için
	stringBody := string(bodyBytes)
	fmt.Println(stringBody)

	//Modele uygun şekilde okumak için
	var todo ToDo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo)
}

func SetValue()  {

	var todo ToDo = ToDo{UserId: 1, Id: 1, Title: "Kod yazılıcak.", Completed: true}
	jsonTodo, _ := json.Marshal(todo)

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset = utf-8", bytes.NewBuffer(jsonTodo))

	if err != nil {
		fmt.Println("Bir hata oluştu :", err.Error())
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	//String
	stringBody := string(bodyBytes)
	fmt.Println(stringBody)

	//Json
	var todoResponse ToDo
	json.Unmarshal(bodyBytes, &todoResponse)
	fmt.Println(todoResponse)
}