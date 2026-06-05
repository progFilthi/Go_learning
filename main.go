package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix

	}
	return prefix + name
}

func main() {

	fmt.Println(hello("FBI", ""))
}
