package reverse_string

import "strings"

func ReverseString(input string) (output string) {
	letters := strings.Split(input, "")
	for i := len(letters) - 1; i >= 0; i-- {
		output += letters[i]
	}
	return output
}
