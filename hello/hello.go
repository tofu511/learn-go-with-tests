package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("world", "English"))
}

const spanish  = "Spanish"
const french = "French"
const helloPrefix = "Hello, "
const spanishHelloPrefix  = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}
	
	if language == french {
		return frenchHelloPrefix + name
	}

	return helloPrefix + name
}