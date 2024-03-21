package main

import (
	"fmt"
)

func main() {
	char := make(chan rune)
	exit := make(chan struct{})
	go func() {
		str := "Hello, world!"
		for _, ch := range str {
			char <- ch
		}
		close(char)
		exit <- struct{}{}
	}()
forLoop:
	for {
		select {
		case ch := <-char:
			fmt.Printf("%c", ch)
		case <-exit:
			fmt.Printf("\n")
			close(exit)
			break forLoop
		}
	}
}
