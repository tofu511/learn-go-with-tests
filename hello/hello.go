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

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}