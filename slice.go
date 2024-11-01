package main

import "fmt"

func main() {
	days := [...]string{"mon", "tue", "wed", "thu", "fri", "sat", "sun"}
	fmt.Println("OLD days >> ", len(days))

	slice1 := days[3:5]
	fmt.Println(slice1)

	slice1[0] = "new_thu"
	fmt.Println(slice1)
	fmt.Println(days)

	slice2 := append(slice1, "week")
	fmt.Println("Slice 2: ", slice2)
	fmt.Println(days)
	fmt.Println("Capacity Slice 2 >> ", cap(slice2))
}
