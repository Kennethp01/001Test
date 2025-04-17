package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}

	var startingASCIINumber int = 65

	fmt.Println("Printing the alphabets from A to Z using a loop")

	for AZ := 0; AZ < 26; AZ++ {
		fmt.Print(string(rune(startingASCIINumber+AZ)), " ")

	}
}
