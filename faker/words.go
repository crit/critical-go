package faker

import (
	"github.com/crit/critical-go/faker/data"

	"strings"
)

func randomWords(min, max int) string {
	count := randomInRange(min, max)
	out := randomElements(data.Words, count)

	return strings.Join(out, " ")
}
