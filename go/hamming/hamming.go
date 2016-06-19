package hamming

import "errors"

const testVersion = 4

// Distance compare two dna strands and return the distance or an error
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Strands have different sizes.")
	}

	distance := 0
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
