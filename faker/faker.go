package faker

import (
	"fmt"
	"strings"
	"time"
)

func firstOrDefault(max []int, def int) int {
	if len(max) == 0 {
		return def
	}

	return max[0]
}

func Words(max ...int) string {
	return randomWords(1, firstOrDefault(max, 8))
}

func Sentences(max ...int) string {
	return randomSentences(1, firstOrDefault(max, 10))
}

func Date() time.Time {
	return randomDate()
}

func Number(max ...int) int {
	return randomInRange(1, firstOrDefault(max, 100))
}

func Email() string {
	parts := strings.Split(randomName(), " ")

	return fmt.Sprintf(
		"%s@%s.com",
		strings.ToLower(parts[0]),
		strings.ToLower(parts[1]),
	)
}

func Name() string {
	return randomName()
}
