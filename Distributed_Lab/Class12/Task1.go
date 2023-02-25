package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	err := os.Mkdir(`Test`, 0777)
	if err != nil {
		if errors.Is(err, fs.ErrExist) {
			fmt.Println("Каталог існує")
		} else {
			fmt.Println(err)
		}
		return
	}
	mode := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(`e:\Reposit\Test\test.txt`, mode, 0o777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	if _, err = file.WriteString("Я люблю програмування\n"); err != nil {
		fmt.Println(err)
		return
	}
	mod := os.O_WRONLY | os.O_APPEND
	f, err := os.OpenFile(`e:\Reposit\Test\test.txt`, mod, 0o777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	if _, err = file.WriteString("Мова програмування Golang\n"); err != nil {
		fmt.Println(err)
		return
	}

}
