package main

import "fmt"

type Reader interface {
	Read()
}
type Writer interface {
	Write()
}

type ReadWrite interface {
	Reader
	Writer
}

type File struct{}

func (File) Read() {
	fmt.Println("Reading the File")
}

func (File) Write() {
	fmt.Println("Writing the File")
}

func main() {
	var rw ReadWrite = File{}
	rw.Read()
	rw.Write()
}
