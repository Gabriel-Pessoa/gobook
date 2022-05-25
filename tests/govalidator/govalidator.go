package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

type Post struct {
	Title    string `valid:"alphanum,required"`
	Message  string `valid:"duck,ascii"`
	Message2 string `valid:"animal(dog)"`
	AuthorIP string `valid:"ipv4"`
	Date     string `valid:"-"`
}

type Book struct {
	Id      string   `valid:"required,uuid"`
	Title   string   `json:"title" valid:"required,stringlength(3|10)"`
	Price   float32  `valid:"float"`
	Authors []string `valid:"-"`
}

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func main() {
	post := &Post{
		Title:    "My Example Post",
		Message:  "duck",
		Message2: "dog",
		AuthorIP: "123.234.54.3",
	}
	if valid, err := govalidator.ValidateStruct(post); !valid {
		getValidationErrors(err)
	}

	book := &Book{
		Id:    "be7a20c4-db72-11ec-9d64-0242ac12000q",
		Title: "A volta dos que não foram",
		Price: 22.80,
	}
	if valid, err := govalidator.ValidateStruct(book); !valid {
		getValidationErrors(err)
	}
}

func getValidationErrors(err error) {
	gvErrs := err.(govalidator.Errors).Errors()
	for _, e := range gvErrs {
		gvErr := e.(govalidator.Error)

		fmt.Printf("Name: %v, Error: %v, Path: %v\n", gvErr.Name, gvErr.Err, gvErr.Path)
	}
}
