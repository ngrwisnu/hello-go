package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) string {
	filteredName := filter(name)

	return "Hello " + filteredName
}

func filterName(name string) string {
	if name == "damn" {
		return "..."
	} else {
		return name
	}
}

func main() {
	res1 := sayHelloWithFilter("damn", filterName)
	res2 := sayHelloWithFilter("Bob", filterName)

	fmt.Println(res1)
	fmt.Println(res2)
}
