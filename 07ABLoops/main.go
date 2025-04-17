package main

import "fmt"

func main() {
	for abUpp := 'A'; abUpp <= 'Z'; abUpp++ {
		fmt.Printf("%c and type: %T\n ", abUpp, abUpp)
	}
	for abLow := 'a'; abLow <= 'z'; abLow++ {
		fmt.Printf("%c and type: %T\n ", abLow, abLow)
	}
}
