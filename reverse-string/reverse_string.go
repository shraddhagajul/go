package reverse

// Reverse returns the reversed string of the input.
func Reverse(input string) string {
	reverse := ""
	runeSlice := []rune(input)
	if len(input) == 0 {
		return reverse
	}

	for i:= len(runeSlice) - 1 ; i >= 0; i--{
		reverse += string(runeSlice[i])
	}
	return reverse
}