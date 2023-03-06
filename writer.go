package hiraoyogi

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var supportedTypes = []string{"analog", "binary", "concise", "robust"}

type HiraoyogiError struct {
	msg string
}

func (error *HiraoyogiError) Error() string {
	return error.msg
}

type ElementType struct {
	Name       string
	Type       string
	identifier string
}

type HiraoyogiWriter struct {
	loadedFile          *os.File
	buffered            *bufio.Writer
	stringIdentifierMap map[string]ElementType
	previousTime        uint64
}

func New(filename string) (HiraoyogiWriter, error) {
	f, err := os.Create(filename)
	writer := HiraoyogiWriter{
		loadedFile:          f,
		buffered:            bufio.NewWriter(f),
		stringIdentifierMap: make(map[string]ElementType),
		previousTime:        0,
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

func (hWriter *HiraoyogiWriter) RegisterElements(elements ...ElementType) (map[string]ElementType, error) {
	for _, element := range elements {
		check2(hWriter.buffered.WriteString(element.Type + " \"" + element.Name + "\" as " + element.identifier + "\n"))
		hWriter.stringIdentifierMap[element.Name] = element
	}
	check2(hWriter.buffered.WriteString("\n@0\n"))
	return hWriter.stringIdentifierMap, nil
}

func (hWriter *HiraoyogiWriter) SetValue(time uint64, value string, elementName string) error {
	if time < hWriter.previousTime {
		return fmt.Errorf("changing value from an earlier time: %d < %d", time, hWriter.previousTime)
	}
	if time != hWriter.previousTime {
		check2(hWriter.buffered.WriteString("\n@" + strconv.FormatUint(time, 10) + "\n"))
		hWriter.previousTime = time
	}
	check2(hWriter.buffered.WriteString(hWriter.stringIdentifierMap[elementName].identifier + " is " + value + "\n"))
	return nil
}

func (hWriter *HiraoyogiWriter) RegisterElementList(elements []ElementType) (map[string]ElementType, error) {
	return hWriter.RegisterElements(elements...)
}

func (hWriter HiraoyogiWriter) Close() {
	check2(hWriter.buffered.WriteString("@enduml\n"))
	check(hWriter.buffered.Flush())
	check(hWriter.loadedFile.Close())
}
