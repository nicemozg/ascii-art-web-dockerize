package funcs

import (
	"fmt"
	"os"
	"strings"
)

func ConvertText(word, banner string) string {
	fmt.Println(banner)
	filePath := "./ascii_art/" + banner + ".txt"
	text, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		bannerError := "Please choose banner to convert the text"
		return bannerError
	}
	if word == "" {
		return ""
	}
	word = strings.ReplaceAll(word, "\\n", "\n")
	newlineCount := strings.Count(word, "\n")
	if len(word) == newlineCount {
		for i := 0; i < newlineCount; i++ {
			fmt.Println()
		}
		return word
	}

	text = []byte(strings.ReplaceAll(string(text), "\r", ""))
	lines := strings.Split(string(text), "\n\n")
	var symbols [][]string
	for _, line := range lines {
		symbols = append(symbols, strings.Split(line, "\n"))
	}
	textForResponse := ""
	words := strings.Split(word, "\n")
	for _, currWord := range words {
		if currWord == "" && words != nil {
			textForResponse += "\n"
		}
		text := PrintChangedWords(symbols, currWord)
		textForResponse += text
	}
	return textForResponse
}
