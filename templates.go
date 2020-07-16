package utils

import (
	"bytes"
	"html/template"
	"log"
)

// TemplateParseString parses a template string with the passed data
func TemplateParseString(templateString string, data interface{}) string {
	t, err := template.New("template").Funcs(template.FuncMap{
		"htmlSafe": func(html string) template.HTML {
			return template.HTML(html)
		},
	}).Parse(templateString)
	if err != nil {
		log.Panic(err.Error())
	}
	var doc bytes.Buffer
	errE := t.Execute(&doc, data)
	if errE != nil {
		log.Panic(errE.Error())
	}
	s := doc.String()
	return s
}
