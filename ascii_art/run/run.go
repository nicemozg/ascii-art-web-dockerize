package run

import (
	"ascii-art-web/ascii_art/funcs"
	"fmt"
	"net/http"
)

func Run(word, banner string) (string, int) {
	fmt.Println(word + " for error")
	fmt.Println(banner + " for error")
	check := funcs.IsPrintable(word)
	if check {
		return "", http.StatusBadRequest
	}
	if banner == "standard" || banner == "shadow" || banner == "thinkertoy" {
		fmt.Println("5")
		text := funcs.ConvertText(word, banner)
		fmt.Println(word)
		fmt.Println(banner)
		return text, http.StatusOK
	}
	return "", http.StatusBadRequest
}
