package main

import (
	"fmt"
	"github.com/angrytit/test-dep-go/localpackage"
)

func main() {
	fmt.Printf("Hello, %s!", localpackage.PublicConst)
}
