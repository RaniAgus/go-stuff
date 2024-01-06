package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	// se usa `json:"-` para ignorar un campo
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		// todo := unmarshalTodo(response)
		todo := decodeTodo(response)
		fmt.Printf("%+v\n", todo)
		fmt.Println(todo.Marshal())
		fmt.Println(todo.MarshalIndent())
	}

}

func unmarshalTodo(response *http.Response) Todo {
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	todo := Todo{}

	if err = json.Unmarshal(bytes, &todo); err != nil {
		log.Fatal(err)
	}

	return todo
}

func decodeTodo(response *http.Response) Todo {
	decoder := json.NewDecoder(response.Body)
	decoder.DisallowUnknownFields() // error if field is not in struct

	todo := Todo{}

	if err := decoder.Decode(&todo); err != nil {
		log.Fatal(err)
	}

	return todo
}

func (t Todo) Marshal() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}

func (t Todo) MarshalIndent() string {
	bytes, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
