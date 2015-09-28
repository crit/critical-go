package faker

import (
	"github.com/crit/critical-go/faker/data"
	"strings"
)

func randomSentences(min, max int) string {
	count := randomInRange(min, max)

	lines := []string{}

	for i := 0; i < count; i++ {
		words := randomElements(data.Words, randomInRange(3, 10))
		words[0] = strings.Title(words[0])

		sentence := strings.Join(words, " ") + randomElement([]string{".", "?", "!"})

		lines = append(lines, sentence)
	}

	return strings.Join(lines, " ")
}
