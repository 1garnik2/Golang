package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

func main() {
	makedir()
	copy()
	rename()
	create()
	getDirContents()
}
func makedir() {
	err := os.Mkdir(`Test2`, 0777)
	if err != nil {
		if errors.Is(err, fs.ErrExist) {
			fmt.Println("Каталог існує")
		} else {
			fmt.Println(err)
		}
		return
	}
}
func copy() {
	sourceFile, err := os.Open("e:/Reposit/Test/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()
	newFile1, err := os.Create("E:/Reposit/Test2/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	bytesCopied, err := io.Copy(newFile1, sourceFile)
	if err != nil {
		fmt.Printf("Error %d", bytesCopied)
	}
	newFile1.Close()
}
func rename() {
	err := os.Rename(`E:/Reposit/Test2/test.txt`, `E:/Reposit/Test2/test2.txt`)
	if err != nil {
		log.Fatal(err)
	}
}
func create() {
	file, err := os.Create(`E:/Reposit/Test2/test3.txt`)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()
	file2, err := os.Create(`E:/Reposit/Test2/test4.txt`)
	if err != nil {
		fmt.Println(err)
		return
	}
	file2.Close()
}
func getDirContents() {
	files, err := os.ReadDir("E:/Reposit/Test2/")
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range files {
		fmt.Println(name.Name())
	}
}
