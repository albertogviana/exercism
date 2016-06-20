package raindrops

import "fmt"

const testVersion = 2

var rain = make([]string, 10)

// Convert a number to a string
func Convert(number int) string {
	rain[3] = "Pling"
	rain[5] = "Plang"
	rain[7] = "Plong"

	sounds := ""

	for i, sound := range rain {
		if 0 < i && number%i == 0 {
			sounds += sound
		}
	}

	if sounds == "" {
		sounds = fmt.Sprintf("%d", number)
	}

	return sounds
}
