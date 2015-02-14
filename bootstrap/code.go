package bootstrap

func Code(content string) string {
	return `<code>` + content + `</code>`
}

func Keyboard(content string) string {
	return `<kbd>` + content + `</kbd>`
}

func Pre(content string) string {
	return `<pre>` + content + `</pre>`
}

func Var(content string) string {
	return `<var>` + content + `</var>`
}
