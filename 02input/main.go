package main

import "fmt"

func main() {
	var age int
	var studentName string
	var luckyNumber int

	fmt.Println(luckyNumber)
	fmt.Print("Your lucky Number is: ")
	fmt.Scanln(&luckyNumber)
	fmt.Println(studentName)
	fmt.Print("Enter your Name: ")
	fmt.Scanln(&studentName)
	fmt.Println(age)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Println("your age is ", age)
}
