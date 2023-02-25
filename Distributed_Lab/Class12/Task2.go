package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileStat, err := os.Stat(`e:\Reposit\Test\test.txt`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Is Directory: ", fileStat.IsDir())
	fmt.Println("File Name:", fileStat.Name())
	fmt.Println("Permissions:", fileStat.Mode())
	fmt.Println("Last Modified:", fileStat.ModTime())
	fmt.Println("Size:", fileStat.Size())
}
