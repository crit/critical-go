package bootstrap

import (
	"strings"
)

func Container(content string) string {
	return `<div class="container">` + content + `</div>`
}

func ContainerFluid(content string) string {
	return `<div class="container-fluid">` + content + `</div>`
}

func Row(content string) string {
	return `<div class="row">` + content + `</div>`
}

func Col(content string, sizes ...string) string {
	var a string

	for _, v := range sizes {
		a = a + " col-" + v
	}

	return `<div class="` + strings.Trim(a, " ") + `">` + content + `</div>`
}
