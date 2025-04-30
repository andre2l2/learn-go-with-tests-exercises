package mainapp

const prefixHello string = "Hello"

func CustomHello(name string, language string) string {
	if name == "" {
		return prefixHello + "!"
	}

	if language == "spanish" {
		return "Hola, " + name + "!"
	}

	return prefixHello + ", " + name + "!"
}
