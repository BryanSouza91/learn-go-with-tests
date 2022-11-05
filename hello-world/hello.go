package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const helloSuffix = "!"

func Hello(name string, lang string) string {
	if name == "" {
		name = "wonderful world"
	}

	if lang == spanish {
		return spanishHelloPrefix + name + helloSuffix
	}

	if lang == french {
		return frenchHelloPrefix + name + helloSuffix
	}
	return englishHelloPrefix + name + helloSuffix
}

func main() {
	fmt.Println(Hello("Bryan", ""))
}
