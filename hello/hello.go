package hello

import "fmt"

const englishPrefix string = "Hello, "
const spanishPrefix string = "Hola, "

func Hello(name string, language string) string {
	var prefix string
	if name == "" {
		name = "World"
	}
	if language == "English" || language == "" {
		prefix = englishPrefix
	} else if language == "Spanish" {
		prefix = spanishPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("", "English"))
}
