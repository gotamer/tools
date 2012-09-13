package tools

// Reverses a string
func Reverse(s string) string {
	b := []byte(s)
	n := ""
	for i := len(b); i > 0; i-- {
		n += string(b[i-1])
	}
	return string(n)
}
