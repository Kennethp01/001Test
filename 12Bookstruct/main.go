package main

import (
	"fmt"
)

type Book struct {
	Title   string
	Author  string
	Date    int
	Barcode int
	Genre   string
}

func main() {
	Book1 := Book{
		Title:   "Loom Death",
		Author:  "Kenneth Peterson",
		Date:    1857,
		Barcode: 5435763,
		Genre:   "Action",
	}
	fmt.Println("Title:", Book1.Title)
	fmt.Println("Author:", Book1.Author)
	fmt.Println("Date:", Book1.Date)
	fmt.Println("Barcode:", Book1.Barcode)
	fmt.Println("Genre:", Book1.Genre)

}
