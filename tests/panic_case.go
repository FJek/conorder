package main

import (
	"fmt"

	"conorder"
)

func main() {
	case1()
}

func case1() {
	fmt.Println("--- case1 start ---")
	defer fmt.Println("--- case1 stoped ---")

	pool, err := conorder.NewPool(1)
	if err != nil {
		panic(err)
	}

	pool.Put(&conorder.Task{
		Handler: func(v ...interface{}) {
			panic("aaaaaa!")
		},
		Params: []interface{}{"hi!"},
	})

	pool.Put(&conorder.Task{
		Handler: func(v ...interface{}) {
			fmt.Println(v)
		},
		Params: []interface{}{"hi!"},
	})

	pool.Close()
	err = pool.Put(&conorder.Task{
		Handler: func(v ...interface{}) {},
	})
	if err != nil {
		fmt.Println(err)
	}
}
