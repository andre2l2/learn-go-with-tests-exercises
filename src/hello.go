package mainapp

const prefixEnglishHello string = "Hello"
const prefixSpanishHello string = "Hola"
const prefixFrenchHello string = "Bonjour"

func CustomHello(name string, language string) string {
	if name == "" {
		return prefixEnglishHello + "!"
	}

	if language == "spanish" {
		return prefixSpanishHello + ", " + name + "!"
	}

	if language == "french" {
		return prefixFrenchHello + ", " + name + "!"
	}

	return prefixEnglishHello + ", " + name + "!"
}
