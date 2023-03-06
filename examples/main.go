package main

import (
	"fmt"

	"github.com/steelpipe75/hiraoyogi"
)

func main() {
	filename := "example.puml"
	writer, e := hiraoyogi.New(filename)
	if e != nil {
		panic(e)
	}
	defer writer.Close()
	fmt.Println("Hello World!")
}
