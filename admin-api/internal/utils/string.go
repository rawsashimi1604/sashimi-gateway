package utils

func TrimLastChar(s string) string {
	r := []rune(s)
	return string(r[:len(r)-1])
}
