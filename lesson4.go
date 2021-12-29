package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func makeWorkers(i *int) {
	context := context.Background()
	for *i = 0; *i < 1000; *i++ {
		go inc(context, i)
	}
	context.Done()
}

func inc(ctx context.Context, num *int) {
	for range ctx.Done() {
		*num++
		return
	}
}

func systemCalled(channel chan os.Signal) {
	for range channel {
		time.Sleep(time.Second)
		fmt.Println("slept")
		return
	}
}
