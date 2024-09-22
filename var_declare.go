package main

import "fmt"

func main() {
	action := func() {
		fmt.Print("In function")
	}

	action()
}
