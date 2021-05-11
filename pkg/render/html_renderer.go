package render

import (
	"bytes"
	"fmt"
	"html/template"
)

type HTMLRenderer struct{}

func NewHTMLRenderer() *HTMLRenderer {
	return &HTMLRenderer{}
}

func (r *HTMLRenderer) Render(input interface{}) ([]byte, error) {
	htmlTemplate := template.New("template")
	htmlTemplate, err := htmlTemplate.Parse(`{{define "T"}}{{.}}{{end}}`)
	if err != nil {
		return nil, err
	}

	buf := bytes.Buffer{}
	if err = htmlTemplate.ExecuteTemplate(&buf, "T", input); err != nil {
		return nil, fmt.Errorf("html render invalid model")
	}

	return buf.Bytes(), nil
}
