package funcs

func IsPrintable(word string) bool {
	for _, char := range word {
		if (char < 32 || char > 126) && char != 10 {
			return true
		}
	}
	return false
}
