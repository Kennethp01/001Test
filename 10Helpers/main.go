package main

import (
	"bufio"
	"fmt"
	"os"
)

func greet(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your full name")
	name, _ := reader.ReadString('\n')
	message := greet("") + (name) + ("!")
	fmt.Println(message)
}
