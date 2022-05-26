package demo

import (
	"fmt"

	demoC "github.com/dshechn/go-test-c/demo"
)

func MyPrint() {
	fmt.Println("i am go-test-b, my version is 1")
	demoC.MyPrint()
}
