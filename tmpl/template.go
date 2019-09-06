package tmpl

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
)

// Data contains variables used in templates.
type Data map[string]string

// Template is a single file.
type Template struct {
	Name      string   // Name used to identify the template. These are unique.
	MultiFile bool     // True if the template is for more than a single file.
	FileName  string   // Name to use when creating a copy of this template.
	Content   string   // The content of the template itself.
	Variables []string // Data used for dynamic template content.
}

// Use copies a single template to a user's current directory,
// applying any data the template has.
func (tmpl Template) Use() {
	cwd, _ := os.Getwd()
	destFile := filepath.Join(cwd, tmpl.FileName)
	content := tmpl.GetContent()

	if _, err := os.Stat(destFile); !os.IsNotExist(err) {
		shouldReplace := false
		err := survey.AskOne(
			&survey.Confirm{Message: "The file already exists, do you want to overwrite it?"},
			&shouldReplace,
			survey.WithValidator(survey.Required),
		)
		if err != nil {
			log.Fatal(err)
		}
		if shouldReplace {
			err = StringToFile(content, destFile)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Template %q copied to current directory\n", tmpl.FileName)
		}
	} else if os.IsNotExist(err) {
		err = StringToFile(content, destFile)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Template %q copied to current directory\n", tmpl.FileName)
	}
}

// GetContent for a template. Prompts for variable values if the template has any.
func (tmpl Template) GetContent() string {
	if len(tmpl.Variables) > 0 {
		return tmpl.Content
	}

	answers := make(map[string]string, len(tmpl.Variables))
	for _, v := range tmpl.Variables {
		response := ""
		prompt := &survey.Input{
			Message: fmt.Sprintf("%s?", v),
		}
		survey.AskOne(prompt, &response)
		answers[v] = response
	}

	var data Data = answers
	t, err := template.New(tmpl.FileName).Parse(tmpl.Content)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

// AddVariables to a template.
func (tmpl *Template) AddVariables(vars ...string) {
	for _, v := range vars {
		tmpl.Variables = append(tmpl.Variables, v)
	}
}
