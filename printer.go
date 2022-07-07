package printer

func FullLine(text string) error {
	fmt.Print(text)
	w, _, err := terminal.GetSize(0)
	if err != nil {
		return err
	}
	more := w - len(text) - 2
	if more <= 1 {
		return nil
	}
	fmt.Print(" " + strings.Repeat("*", more))
	fmt.Println()
	return nil
}
