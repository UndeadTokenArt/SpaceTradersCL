package handlers

func isValidLength(s string) bool {
	length := len(s)
	return length >= 4 && length <= 14
}
