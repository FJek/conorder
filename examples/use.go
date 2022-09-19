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

	for i := 0; i < 20; i++ {
		pool.Put(&conorder.Task{
			Handler: func(v ...interface{}) {
				fmt.Println(v)
			},
			Params: []interface{}{i},
		})
	}

	time.Sleep(1e9)
}
