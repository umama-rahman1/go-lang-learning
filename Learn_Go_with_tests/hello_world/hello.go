package main

import "fmt"

const englishHelloPreix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPreix + name
}

func main() {
	fmt.Println(Hello("world"))
}
