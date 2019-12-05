package main

import (
	"fmt"

	"github.com/frzifus/winsize"
)

func main() {
	s, err := winsize.Get()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", s)
}
