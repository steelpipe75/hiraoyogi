package main

import (
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

	writer.SetValue(0, "1.0", "test analog")
	writer.SetValue(0, "low", "test binary")
	writer.SetValue(0, "Idle", "test concise")
	writer.SetValue(0, "アイドル", "test robust")

	writer.SetValue(10, "2.0", "test analog")
	writer.SetValue(10, "high", "test binary")
	writer.SetValue(10, "Runing", "test concise")
	writer.SetValue(10, "実行中", "test robust")

	writer.SetValue(20, "0.0", "test analog")
	writer.SetValue(20, "low", "test binary")
	writer.SetValue(20, "{-}", "test concise")
	writer.SetValue(20, "返信待ち", "test robust")
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

	writer.SetValue(0, "1.0", "test analog")
	writer.SetValue(0, "low", "test binary")
	writer.SetValue(0, "idle", "test concise")
	writer.SetValue(0, "アイドル", "test robust")

	writer.SetValue(100, "2.0", "test analog")
	writer.SetValue(100, "high", "test binary")
	writer.SetValue(100, "Runing", "test concise")
	writer.SetValue(100, "実行中", "test robust")

	writer.SetValue(200, "0.0", "test analog")
	writer.SetValue(200, "low", "test binary")
	writer.SetValue(200, "{-}", "test concise")
	writer.SetValue(200, "返信待ち", "test robust")
}

func main() {
	listExample()
	vParamExample()
}
