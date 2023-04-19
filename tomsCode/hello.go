package main

import "fmt"

const greetingPrefix = "Hello, "

func hello(name string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix + name + "!"
}

func main() {
	fmt.Println(hello("world"))
}
