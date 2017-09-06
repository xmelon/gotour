package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	fmt.Println("Hello, World!")

	errors.New("err")
}
