package main

import (
	"fmt"

	"github.com/steelpipe75/hiraoyogi"
)

type TestElement struct {
	Name string
	Type string
}

func listExample() {
	filename := "listExample.puml"
	writer, e := hiraoyogi.New(filename)
	if e != nil {
		panic(e)
	}
	defer writer.Close()
	elementSlice := []hiraoyogi.ElementType{}

	testSlice := []TestElement{}
	testSlice = append(testSlice, TestElement{Name: "test analog", Type: "analog"})
	testSlice = append(testSlice, TestElement{Name: "test binary", Type: "binary"})
	testSlice = append(testSlice, TestElement{Name: "test concise", Type: "concise"})
	testSlice = append(testSlice, TestElement{Name: "test robust", Type: "robust"})

	for _, t := range testSlice {
		element := writer.NewElement(t.Name, t.Type)
		elementSlice = append(elementSlice, element)
	}

	writer.RegisterElementList(elementSlice)
}

func vParamExample() {
	filename := "vParamExample.puml"
	writer, e := hiraoyogi.New(filename)
	if e != nil {
		panic(e)
	}
	defer writer.Close()

	writer.RegisterElements(
		writer.NewElement("test analog", "analog"),
		writer.NewElement("test binary", "binary"),
		writer.NewElement("test concise", "concise"),
		writer.NewElement("test robust", "robust"),
	)
}

func main() {
	listExample()
	vParamExample()
	fmt.Println("Hello World!")
}
