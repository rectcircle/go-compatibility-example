package main

import (
	"fmt"
	"runtime"

	"github.com/rectcircle/go-compatibility-example/02-gotoolchain/01-goline-new/math"
)

func main() {
	_ = math.Add(1, 2)
	fmt.Println(runtime.Version())
}
