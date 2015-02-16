package faker

import "math/rand"

func randomInRange(min, max int) int {
	// max value inclusive
	i := max - min

	if i < 1 {
		i = 1
	}

	return rand.Intn(i) + min
}
