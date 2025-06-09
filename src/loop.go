package mainapp

func Loop(value string) string {
	var acc string
	for i := 0; i < 5; i++ {
		acc += value
	}

	return acc
}
