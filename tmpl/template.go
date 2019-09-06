package tmpl

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gookit/color"
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
func (t Template) Use() {
	cwd, _ := os.Getwd()
	destFile := filepath.Join(cwd, t.FileName)
	content := t.GetContent()
	blue := color.FgBlue.Render

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
			log.Printf("Template %s copied to current directory\n", blue(t.FileName))
		}
	} else if os.IsNotExist(err) {
		err = StringToFile(content, destFile)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Template %s copied to current directory\n", blue(t.FileName))
	}
}

// GetContent for a template. Prompts for variable values if the template has any.
func (t Template) GetContent() string {
	if len(t.Variables) > 0 {
		return t.Content
	}

	answers := make(map[string]string, len(t.Variables))
	for _, v := range t.Variables {
		response := ""
		prompt := &survey.Input{Message: v + "?"}
		survey.AskOne(prompt, &response)
		answers[v] = response
	}

	var data Data = answers
	tpl, err := template.New(t.FileName).Parse(t.Content)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

// AddVariables to a template.
func (t *Template) AddVariables(vars ...string) {
	for _, v := range vars {
		t.Variables = append(t.Variables, v)
	}
}
