package main

import "fmt"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return fmt.Sprintf("Hola, %s", name)
	} else {
		return fmt.Sprintf("Hello, %s", name)
	}
}

func main() {

}
