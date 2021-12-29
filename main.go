package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Главная функция, для проверки заданий
func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Error.main: %v", v)
		}
	}()

	// ex1()
	//sayHello()

	// ex2()

	// ex3()

	// i := 0
	// makeWorkers(&i)
	// fmt.Println(i)

	sysCallChan := make(chan os.Signal, 1)
	signal.Notify(sysCallChan, syscall.SIGTERM)
	systemCalled(sysCallChan)
}
