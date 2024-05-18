package funcs

import (
	"bytes"
)

func PrintChangedWords(symbols [][]string, currWord string) string {
	var result bytes.Buffer // Создаем буфер для накопления результата

	for i := 0; i < 8; i++ {
		if i == 0 {
			symbols[0][0] = "      "
		}
		for _, char := range currWord {
			if char >= 32 && char <= 126 && char != 10 {
				result.WriteString(symbols[char-32][i]) // Записываем символ в буфер
			}
		}
		if currWord != "" {
			result.WriteString("\n") // Добавляем перевод строки, если слово не пустое
		}
	}

	return result.String() // Возвращаем содержимое буфера как строку
}
