package main

import (
	"fmt"

	"example.com/stringutil"
)

func main() {
	fmt.Println("Hello, World!")

	message := stringutil.Reverse("!oG ,olleH")
	fmt.Println(message)

	message1 := stringutil.Hello("Lester")
	fmt.Printf(message1)
}
