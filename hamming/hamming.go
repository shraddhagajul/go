package hamming

import (
	"errors"
)

/*Distance is the hamming distance between 2 strands of DNA
strands must have same length*/
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("Strings must be of equal length")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
