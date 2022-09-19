package main

import (
	"fmt"
	"time"

	"conorder"
)

func main() {
	pool, err := conorder.NewPool(10)
	if err != nil {
		panic(err)
	}

	pool.PanicHandler = func(r interface{}) {
		fmt.Printf("Warning!!! %s", r)
	}

	pool.Put(&conorder.Task{
		Handler: func(v ...interface{}) {
			panic("somthing wrong!")
		},
	})

	time.Sleep(1e9)
}
