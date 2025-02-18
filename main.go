package main

import (
	"app/table"
)

type Item struct {
	name  string
	price int
}

type Person struct {
	name string
	age  int
	item Item
}

func main() {
	// str1 := "amy"
	// data := []any{str1, 200, "abc", []int{1, 2, 3, 4}, Person{"Bob", 23}}

	table.Print([]Person{
		{"Bob", 23, Item{}},
		{"Chris", 24, Item{"Some item", 213}},
	})
}
