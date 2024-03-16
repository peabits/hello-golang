package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	if content, err := os.ReadFile("resources/gopher.txt"); err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", content)
	}
}
