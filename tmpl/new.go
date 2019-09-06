package tmpl

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

// SingleFromFile creates a new single-file template from an existing file.
func SingleFromFile() error {
	var t Template

	templates, err := GetTemplates()
	if err != nil {
		return err
	}

	cwd, _ := os.Getwd()
	fileName := PickFile(cwd, "Choose a file to make a template from")
	t.FileName = fileName

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	t.Content = string(bytes)

	for {
		var name string
		prompt := &survey.Input{
			Message: "Please choose a name for this template",
		}
		survey.AskOne(prompt, &name)
		if !templates.NameExists(name) {
			t.Name = name
			break
		}
		log.Println("A template with this name already exists. Please choose another.")
	}
	err = templates.Add(t)

	return err
}
