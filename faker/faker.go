package faker

import (
	"strings"
	"time"
)

func getOption(def int, max []int) int {
	if len(max) == 0 {
		return def
	}

	return max[0]
}

func Words(max ...int) string {
	return randomWords(1, getOption(8, max))
}

func Sentences(max ...int) string {
	return randomSentences(1, getOption(10, max))
}

func Date() time.Time {
	return randomDate()
}

func Number(max ...int) int {
	return randomInRange(1, getOption(100, max))
}

func Email() string {
	a := randomWords(1, 1)
	return strings.ToLower(a) + "@example.com"
}

func Name() string {
	return randomName()
}
