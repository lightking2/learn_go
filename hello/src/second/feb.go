package main

import "fmt"

func main() {
	a := 1
	b := 1
	for i := 0; i < 5; i++ {
		fmt.Print(b)
		tmp := b
		b = a + b
		a = tmp
	}
}
