package bootstrap

func H1(content string) string {
	return `<h1>` + content + `</h1>`
}

func H2(content string) string {
	return `<h2>` + content + `</h2>`
}

func H3(content string) string {
	return `<h3>` + content + `</h3>`
}

func H4(content string) string {
	return `<h4>` + content + `</h4>`
}

func H5(content string) string {
	return `<h5>` + content + `</h5>`
}

func H6(content string) string {
	return `<h6>` + content + `</h6>`
}

func Small(content string) string {
	return `<small>` + content + `</small>`
}

func P(content string) string {
	return `<p>` + content + `</p>`
}

func Lead(content string) string {
	return `<p class="lead">` + content + `</p>`
}

func Mark(content string) string {
	return `<mark>` + content + `</mark>`
}

func Deleted(content string) string {
	return `<del>` + content + `</del>`
}

func Strikethrough(content string) string {
	return `<s>` + content + `</s>`
}

func Inserted(content string) string {
	return `<ins>` + content + `</ins>`
}

func Underlined(content string) string {
	return `<u>` + content + `</u>`
}

func Bold(content string) string {
	return `<b>` + content + `</b>`
}

func Italics(content string) string {
	return `<em>` + content + `</em>`
}

func Right(content string) string {
	return `<span class="text-right">` + content + `</span>`
}

func Left(content string) string {
	return `<span class="text-left">` + content + `</span>`
}

func Center(content string) string {
	return `<span class="text-center">` + content + `</span>`
}

func Justify(content string) string {
	return `<span class="text-justify">` + content + `</span>`
}

func Lowercase(content string) string {
	return `<span class="text-lowercase">` + content + `</span>`
}

func Uppercase(content string) string {
	return `<span class="text-uppercase">` + content + `</span>`
}

func Capitalize(content string) string {
	return `<span class="text-capitalize">` + content + `</span>`
}

func Abbreviation(full, abbr string) string {
	return `<abbr title="` + full + `">` + abbr + `</abbr>`
}

func Initialism(full, abbr string) string {
	return `<abbr title="` + full + `" class="initialism">` + abbr + `</abbr>`
}

func Address(content string) string {
	return `<address>` + content + `</address>`
}

func BlockQuote(content string, classes ...string) string {
	return `<blockquote>` + content + `</blockquote>`
}
