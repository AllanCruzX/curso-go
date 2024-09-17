package main

import (
	"bufio"
	"fmt"
	"os"
)

const txt = ".txt"

type File struct {
	Name   string
	Format string
}

func NewFile(name string) *File {
	return &File{Name: name, Format: txt}
}

func (f File) getFile() string {
	return f.Name + f.Format
}

func main() {

	fileObj := NewFile("documento")
	document := fileObj.getFile()

	//createFile(document)
	loadFile(document)
	//loadFileLazy(document)
	//removeFile(document)

}

func createFile(fileParam string) {

	file, err := os.Create(fileParam)
	if err != nil {
		panic(err)
	}

	tamanho, err := file.Write([]byte("Escrevendo alguma coisa no arquivo"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes\n", tamanho)
	file.Close()

}

func loadFile(fileParam string) {

	file, err := os.ReadFile(fileParam)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))
}

func loadFileLazy(fileParam string) {

	file, err := os.Open(fileParam)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			println(".")
			break
		}

		fmt.Print(string(buffer[:n]))
	}

}

func removeFile(fileParam string) {

	err := os.Remove(fileParam)
	if err != nil {
		panic(err)
	}

}
