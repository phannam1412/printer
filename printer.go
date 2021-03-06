package printer

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func FullLine(text string) {
	fmt.Print(text)
	w, _, err := terminal.GetSize(0)
	if err != nil {
		fmt.Println()
		return
	}
	more := w - len(text) - 1
	if more <= 2 {
		fmt.Println()
		return
	}
	fmt.Print(" " + strings.Repeat("*", more) + "\n")
}
