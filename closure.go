package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		counter := 10

		counter++
		fmt.Println("Incr Counter >> ", counter)
	}

	increment()
	increment()

	fmt.Println("Main Counter >> ", counter)
}
