package form

import (
	"bytes"
	"html/template"
)

const textInput = `<input type="text" name="{{.Name}}" value="{{.Value}}"{{if .Required}} required{{end}}{{if .Disabled}} disabled{{end}}>`
const hiddenInput = `<input type="hidden" name="{{.Name}}" value="{{.Value}}">`
const checkboxInput = `<input type="checkbox" name="{{.Name}}"{{if .Active}} checked{{end}}{{if .Required}} required{{end}}{{if .Disabled}} disabled{{end}}>`

func toTextField(i input) string {
	r := new(bytes.Buffer)
	t := template.Must(template.New("textInput").Parse(textInput))

	t.Execute(r, i)

	return r.String()
}

func toHiddenField(i input) string {
	r := new(bytes.Buffer)
	t := template.Must(template.New("hiddenInput").Parse(hiddenInput))

	t.Execute(r, i)

	return r.String()
}

func toCheckboxField(i input) string {
	r := new(bytes.Buffer)
	t := template.Must(template.New("checkboxInput").Parse(checkboxInput))

	t.Execute(r, i)

	return r.String()
}
