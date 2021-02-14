package main

import "fmt"

func main() {
	fmt.Println("Number: %e", 0.000057346)
	var out string = fmt.Sprintf("Number: %07d is cool", 45)
	fmt.Println(out)
}

//https://golang.org/pkg/fmt/
