package main

import "fmt"

func main() {
	var number = 20000.98
	fmt.Printf("%T", number)

	fmt.Println()

	number2 := 2 // After this, we can't assingn/ overwrite any type to number 2, except int
	fmt.Printf("%T", number2)

	fmt.Println()

	var number3 uint64 // if we haven't assign a value to this variable the value will be 0(default)
	var bl bool        // the default value is false
	fmt.Println(number3, bl)

}

// go run [filename.go] to automatically compile and run the program
// go build [filename.go] to only compile the program
