package hiraoyogi

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var supportedTypes = []string{"analog", "binary", "concise", "robust"}

type HiraoyogiError struct {
	msg string
}

func (error *HiraoyogiError) Error() string {
	return error.msg
}

type HiraoyogiWriter struct {
	loadedFile *os.File
	buffered   *bufio.Writer
}

type ElementType struct {
	Name       string
	Type       string
	identifier string
}

func New(filename string) (HiraoyogiWriter, error) {
	f, err := os.Create(filename)
	writer := HiraoyogiWriter{
		loadedFile: f,
		buffered:   bufio.NewWriter(f),
	}
	if err == nil {
		check2(writer.buffered.WriteString("@startuml\n"))
		check(writer.buffered.Flush())
	}
	return writer, err
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func check2(data interface{}, e error) interface{} {
	if e != nil {
		panic(e)
	}
	return data
}

func (hWriter *HiraoyogiWriter) NewElement(elementName string, typeName string) ElementType {
	if !stringInSlice(typeName, supportedTypes) {
		errorStr := fmt.Sprintf("unsupported data type: \"%s\" supported types: %v", typeName, supportedTypes)
		panic(errorStr)
	}
	return ElementType{Name: elementName, Type: typeName, identifier: strings.Replace(elementName, " ", "_", -1)}
}

func (hWriter *HiraoyogiWriter) RegisterElements(elements ...ElementType) error {
	for _, element := range elements {
		check2(hWriter.buffered.WriteString(element.Type + " \"" + element.Name + "\" as " + element.identifier + "\n"))
	}
	check2(hWriter.buffered.WriteString("\n"))
	return nil
}

func (hWriter *HiraoyogiWriter) RegisterElementList(elements []ElementType) error {
	return hWriter.RegisterElements(elements...)
}

func (hWriter HiraoyogiWriter) Close() {
	check2(hWriter.buffered.WriteString("@enduml\n"))
	check(hWriter.buffered.Flush())
	check(hWriter.loadedFile.Close())
}
