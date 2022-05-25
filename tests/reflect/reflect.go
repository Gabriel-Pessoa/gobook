package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	Id      int      `param:"id"`
	Title   string   `param:"title"`
	Price   float32  `param:"price"`
	Authors []string `param:"authors"`
}

func main() {
	var book Book
	ListFields(&book)
}

func ListFields(a interface{}) {
	v := reflect.ValueOf(a).Elem()
	for j := 0; j < v.NumField(); j++ {
		f := v.Field(j)
		n := v.Type().Field(j).Name
		t := f.Type().String()
		p := v.Type().Field(j).Tag.Get("param")
		fmt.Printf("Name: %s  Kind: %s  Type: %s Param: %s\n", n, f.Kind(), t, p)
	}
}
