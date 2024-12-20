package templates

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"text/template"
)

//go:embed errors.json
var errors []byte

var templates map[string]*template.Template

func init() {
	var errorsMap map[string]string
	err := json.Unmarshal(errors, &errorsMap)
	if err != nil {
		panic(err)
	}

	templates = make(map[string]*template.Template, len(errorsMap))
	for key, tmpl := range errorsMap {
		templates[key] = template.Must(template.New(key).Parse(tmpl))
	}
}

func CreateMessage(key string, data map[string]any) string {
	var text bytes.Buffer
	err := templates[key].Execute(&text, data)
	if err != nil {
		panic(err)
	}
	return text.String()
}
