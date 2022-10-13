package main

import "fmt"

func main() {
	for x := 1; true; x += 10 {
		fmt.Println(x)
	}
}
