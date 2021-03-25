package main

import (
	"fmt"
	"time"
)

func main() {
	{
		startedAt := time.Now()
		defer fmt.Println("Defer run ...")
		defer func() {fmt.Println(time.Since(startedAt))}()
		defer func() {
			fmt.Println("Func in Defer func")
		}()
		fmt.Println("Defer run after ...")
	}

	fmt.Println("Main End ...")
}