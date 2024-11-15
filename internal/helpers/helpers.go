package helpers

// RevertString - функция, которая переворачивает входную строку
func RevertString(str string) string {
	data := []rune(str)

	for i := 0; i < len(data)/2; i++ {
		data[i], data[len(data)-1-i] = data[len(data)-1-i], data[i]
	}

	return string(data)
}
