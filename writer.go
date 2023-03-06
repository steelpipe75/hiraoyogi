package hiraoyogi

import (
	"bufio"
	"os"
)

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

func (hWriter HiraoyogiWriter) Close() {
	check2(hWriter.buffered.WriteString("@enduml\n"))
	check(hWriter.buffered.Flush())
	check(hWriter.loadedFile.Close())
}
