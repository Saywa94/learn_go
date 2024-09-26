package hello_world

import "fmt"

const (
	englishGreeting = "Hello"
	spanishGreeting = "Hola"
	frenchGreeting  = "Bonjour"
)

func Hello(name string, language string) string {

	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("%s, %s!", greetingPrefix(language), name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = spanishGreeting
	case "french":
		prefix = frenchGreeting
	default:
		prefix = englishGreeting
	}
	return
}

func main() {
	// fmt.Println(Hello(""))
}
