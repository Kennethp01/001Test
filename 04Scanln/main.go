package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your full name")
	name, _ := reader.ReadString('\n') //includes newline character
	fmt.Println("Your name is ", name)
	// fmt.Println(reflect.Type0f(name))

	fmt.Println("What is your age?")
	age, _ := reader.ReadString('\n')
	fmt.Println("You are this ", age)

	fmt.Println("what is your lucky Number")
	luckyNumber, _ := reader.ReadString('\n')
	fmt.Println("Your lucky Number is ", luckyNumber)

	fmt.Println("what is your funfact")
	funfact, _ := reader.ReadString('\n')
	fmt.Println("Your fun fact is ", funfact)
}
