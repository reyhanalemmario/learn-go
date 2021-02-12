package main

import "fmt"

func main() {
	var name string = "Hello"

	// var name_of_the_variable type
	var name2 string
	name2 = "Another Hello"
	name2 = "Change Hello" //Overwrite

	var number uint = 260

	fmt.Println("Hello Cruel World!", name, name2, number)
}

// go run [filename.go] to automatically compile and run the program
// go build [filename.go] to only compile the program
