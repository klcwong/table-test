package main

import "app/table"

func main() {
	table.Print([][][]int{
		{{1, 2, 3}, {4, 5, 6}},
	})
}
