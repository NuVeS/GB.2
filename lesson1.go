package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type TimeError struct {
	curTime time.Time
}

func (err *TimeError) Error() string {
	return err.curTime.String()
}

func ex1() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Error.ex1.recovered")
		}
	}()

	fmt.Println("Started flow")
	panic("paniced")
}

func sayHello() {
	fmt.Println("Hello")
}

func ex2() {
	timeErr := &TimeError{time.Now()}
	panic(timeErr)
}

func ex3() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Error.ex3.recovered: %v", v)
		}
	}()

	fileName := "text"
	fileType := ".txt"

	if err := os.Mkdir("test_data", 0777); err != nil {
		panic(err)
	}

	for i := 1; i < 3; i++ {

		file, err := os.Create(fileName + strconv.Itoa(i) + fileType)
		if err != nil {
			panic(err)
		}

		defer file.Close()
	}
}
