package tmpl

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
)

// ListTemplates lists a user's templates.
func ListTemplates() {
	tmplDir := viper.GetString("templatesdir")
	templates := ReadDir(tmplDir)
	if len(templates) == 0 {
		log.Println("No templates")
	} else {
		for _, file := range templates {
			log.Println(file)
		}
	}
}

// UseTemplate copies a template to the current directory.
func UseTemplate() {
	tmplDir := viper.GetString("templatesdir")
	cwd, _ := os.Getwd()
	fileName := PickFile(tmplDir, "Choose a template to copy to the current directory")
	sourceFile := filepath.Join(tmplDir, fileName)
	destFile := filepath.Join(cwd, fileName)

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
			err = overwriteFile(sourceFile, destFile)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Template %q copied to current directory\n", fileName)
		}
	} else if os.IsNotExist(err) {
		err = copyFile(sourceFile, destFile)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Template %q copied to current directory\n", fileName)
	}
}

// CopyToTemplate creates a template from a file in the current directory.
func CopyToTemplate() {
	tmplDir := viper.GetString("templatesdir")
	cwd, _ := os.Getwd()
	fileName := PickFile(cwd, "Choose a file to make a template from")
	sourceFile := filepath.Join(cwd, fileName)
	destFile := filepath.Join(tmplDir, fileName)

	if _, err := os.Stat(destFile); !os.IsNotExist(err) {
		shouldReplace := false
		err := survey.AskOne(
			&survey.Confirm{Message: "A template with this name already exists, do you want to overwrite it?"},
			&shouldReplace,
			survey.WithValidator(survey.Required),
		)
		if err != nil {
			log.Fatal(err)
		}
		if shouldReplace {
			err = overwriteFile(sourceFile, destFile)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("New template %q was created\n", fileName)
		}
	} else {
		err = copyFile(sourceFile, destFile)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("New template %q was created\n", fileName)
	}
}

// DeleteTemplate deletes an existing template.
func DeleteTemplate() {
	tmplDir := viper.GetString("templatesdir")
	fileName := PickFile(tmplDir, "Choose a template to delete")

	shouldDelete := false
	err := survey.AskOne(
		&survey.Confirm{Message: "Are you sure you want to delete this template?"},
		&shouldDelete,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		log.Fatal(err)
	}

	if shouldDelete {
		err := os.Remove(filepath.Join(tmplDir, fileName))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Template %q deleted\n", fileName)
	}
}

// EditTemplate edits an existing template.
func EditTemplate() {
	tmplDir := viper.GetString("templatesdir")
	fileName := PickFile(tmplDir, "Choose a template to edit")
	path := filepath.Join(tmplDir, fileName)
	editFile(path)
}

// NewTemplate makes a new template and opens it in `$EDITOR`.
func NewTemplate() {
	tmplDir := viper.GetString("templatesdir")
	fileName := ""
	survey.AskOne(&survey.Input{Message: "Enter a name for this template"}, &fileName)
	filePath := filepath.Join(tmplDir, fileName)

	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		shouldReplace := false
		err := survey.AskOne(
			&survey.Confirm{Message: "A template with this name already exists, do you want to overwrite it?"},
			&shouldReplace,
			survey.WithValidator(survey.Required),
		)
		if err != nil {
			log.Fatal(err)
		}
		if shouldReplace {
			err = ioutil.WriteFile(filePath, []byte{}, os.FileMode(0644))
			if err != nil {
				log.Fatal(err)
			}
		}
	} else if os.IsNotExist(err) {
		err = ioutil.WriteFile(filePath, []byte{}, os.FileMode(0644))
		if err != nil {
			log.Fatal(err)
		}
	}
	editFile(filePath)
}

// MakeFirstTemplate is run if no templates exist.
func MakeFirstTemplate() {

	firstFile := false
	err := survey.AskOne(
		&survey.Confirm{Message: "Create a new template?"},
		&firstFile,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		log.Fatal(err)
	}

	if firstFile {
		tmplDir := viper.GetString("templatesdir")
		cwd, _ := os.Getwd()
		fileName := PickFile(cwd, "Choose a file to create a template from")

		sourceFile, _ := filepath.Abs(fileName)
		destFile := filepath.Join(tmplDir, fileName)
		err = copyFile(sourceFile, destFile)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("New template %q was created\n", fileName)
	}
}

// PickFile prompts a user to choose a file from a given directory.
func PickFile(dir, message string) string {
	fileName := ""
	err := survey.AskOne(
		&survey.Select{
			Message: message,
			Options: ReadDir(dir),
		},
		&fileName,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		log.Fatal(err)
	}
	return fileName
}
