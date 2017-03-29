package main

import "fmt"

func main() {
	for i := 122; i < 300; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
