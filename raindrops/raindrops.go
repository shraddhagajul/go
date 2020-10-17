package raindrops
import "fmt"
// rain drop sounds
var rain = make([]string, 10)
func Convert(input int) string {
	sounds := ""
	rain[3] = "Pling"
	rain[5] = "Plang"
	rain[7] = "Plong"
	for i, sound := range rain {
		if 0 < i && input%i == 0 {
			sounds += sound
		}
	}
	if sounds == "" {
		sounds = fmt.Sprintf("%d", input)
	}
	return sounds
}