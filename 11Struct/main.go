package main

import "fmt"

type Student struct {
	Name   string
	School string
	Year   int
}

func main() {
	student1 := Student{
		Name:   "Ken",
		School: "Founders College",
		Year:   2025,
	}

	fmt.Println("Student inforamtion:")
	fmt.Println("Name:", student1.Name)
	fmt.Println("School:", student1.School)
	fmt.Println("Year:", student1.Year)
}
