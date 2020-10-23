package luhn
import (
"strings"
"strconv"
)

func Valid(input string) bool{
// no space allowed 
	input = strings.ReplaceAll(input," ","")

// len > 1
	if len(input) <= 1{
		return false
	}

	sum := 0
	for i := len(input); i > 0; i-- {
		// convert string to int
		digit, _ := strconv.Atoi(string(input[i]))

		if i%2 == 0{
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum % 10 == 0
}
