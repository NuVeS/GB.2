package main

import (
	"fmt"
)

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Error.main: %v", v)
		}
	}()

	// ex1()
	//sayHello()

	// ex2()

	ex3()
}
