package main

import "fmt"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	var prefix string

	prefix = greetingPrefix(language)

	return fmt.Sprintf("%s, %s", prefix, name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = "Hola"
	case "French":
		prefix = "Bonjour"
	default :
		prefix = "Hello"
	}
	return

}

func main() {

}
