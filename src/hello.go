package mainapp

const prefixEnglishHello string = "Hello"
const prefixSpanishHello string = "Hola"
const prefixFrenchHello string = "Bonjour"

func CustomHello(name string, language string) string {
	if name == "" {
		return prefixEnglishHello + "!"
	}

	switch language {
	case "spanish":
		return parseWithPrefix(prefixSpanishHello, name)
	case "french":
		return parseWithPrefix(prefixFrenchHello, name)
	}

	return parseWithPrefix(prefixEnglishHello, name)
}

func parseWithPrefix(prefix string, name string) (value string) {
	value = prefix + ", " + name + "!"
	return value
}
