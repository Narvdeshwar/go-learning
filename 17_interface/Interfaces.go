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

func (f File) Read() {
	fmt.Println("Reading the File")
}

func (f File) Write() {
	fmt.Println("Writing the File")
}

func main() {
	newFile := File{}
	newFile.Read()
	newFile.Write()
}
