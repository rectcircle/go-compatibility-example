package main

import (
	"fmt"
	"runtime"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover: ", r)
		} else {
			fmt.Println("recover is nil")
		}
	}()
	fmt.Println(runtime.Version())
	panic(nil)
}
