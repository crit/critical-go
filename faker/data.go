package faker

import (
	"strconv"
	"strings"
)

func randomElements(provider []string, count int) []string {
	length := len(provider)

	if count >= length {
		return provider
	}

	out := []string{}

	for i := 0; i < count; i++ {
		j := randomInRange(0, length-1)
		out = append(out, provider[j])
	}

	return out
}

func randomElement(provider []string) string {
	found := randomElements(provider, 1)
	return found[0]
}

func numerify(provider string) string {
	set := strings.Fields(provider)

	for i, p := range set {
		if p == "#" {
			set[i] = strconv.Itoa(randomInRange(0, 9))
		}
	}

	return strings.Join(set, "")
}
