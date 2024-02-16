package caesar

func Cipher(s string, shift int) string {
	var result []rune
	for _, char := range s {
		var newChar = char
		result = append(result, newChar)
	}
	return string(result)
}
