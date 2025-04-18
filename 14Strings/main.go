package main

import "fmt"

func main() {
	s1 := "Hello" + "World"

	s2 := `A "raw" string literal can include line breaks.`

	// Outputs: 10
	fmt.Println(len(s1))

	// Outputs: Hello
	fmt.Println(string(s1[0:5]))
	// Outputs: all of s2
	fmt.Println(string(s2[0:47]))
	// Outputs: all of s1
	fmt.Println(string(s1[0:10]))
}
