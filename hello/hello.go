package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// named return value!
// evaluates to:
// func getGreetingPrefix(language string) (prefix string) {
// 		string prefix = ""
//		...
//		return prefix
// }
// starts with lowercase letter == private function
func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

// starts with uppercase letter == public function
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getGreetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
