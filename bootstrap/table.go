package bootstrap

func Table(content string) string {
	return `<table class="table">` + content + `</table>`
}

func TableStripped(content string) string {
	return `<table class="table table-stripped">` + content + `</table>`
}

func TableBordered(content string) string {
	return `<table class="table table-bordered">` + content + `</table>`
}

func TableHover(content string) string {
	return `<table class="table table-hover">` + content + `</table>`
}

func TableCondensed(content string) string {
	return `<table class="table table-condensed">` + content + `</table>`
}

func TableHeader(content string) string {
	return `<th>` + content + `</th>`
}

func TableBody(content string) string {
	return `<tbody>` + content + `</tbody>`
}

func TableFooter(content string) string {
	return `<tfooter></tfooter>`
}

func TableRow(content string) string {
	return `<tr>` + content + `</tr>`
}

func TableRowActive(content string) string {
	return `<tr class="active">` + content + `</tr>`
}

func TableRowSuccess(content string) string {
	return `<tr class="success">` + content + `</tr>`
}

func TableRowInfo(content string) string {
	return `<tr class="info">` + content + `</tr>`
}

func TableRowWarning(content string) string {
	return `<tr class="warning">` + content + `</tr>`
}

func TableRowDanger(content string) string {
	return `<tr class="danger">` + content + `</tr>`
}
