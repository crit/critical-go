package bootstrap

import (
	"fmt"
)

func classer(base string, classes ...string) {
	return fmt.Sprintf(base, classes)
}
