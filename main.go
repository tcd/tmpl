package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey"
)

var createFlag = flag.Bool("create", false, "Create a template from an existing file")
var deleteFlag = flag.Bool("remove", false, "Delete a template")
var editFlag = flag.Bool("update", false, "Edit an existing template")
var listFlag = flag.Bool("list", false, "List templates")
var newFlag = flag.Bool("new", false, "Add a new template")

func init() {
	flag.BoolVar(createFlag, "c", false, "")
	flag.BoolVar(deleteFlag, "d", false, "")
	flag.BoolVar(editFlag, "e", false, "")
	flag.BoolVar(listFlag, "l", false, "")
	flag.BoolVar(newFlag, "n", false, "")
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if len(readDir(templatesDir())) == 0 {
		makeFirstTemplate()
		os.Exit(0)
	}

	if *createFlag {
		copyToTemplate()
		os.Exit(0)
	}

	if *deleteFlag {
		deleteTemplate()
		os.Exit(0)
	}

	if *editFlag {
		editTemplate()
		os.Exit(0)
	}

	if *listFlag {
		for _, file := range readDir(templatesDir()) {
			fmt.Println(file)
		}
		os.Exit(0)
	}

	if *newFlag {
		newTemplate()
		os.Exit(0)
	}

	useTemplate()
	os.Exit(0)
}

// Copy a template to the current directory
func useTemplate() {
	cwd, _ := os.Getwd()
	fileName := pickFile(cwd, "Choose a template to copy to the current directory")
	sourceFile := filepath.Join(templatesDir(), fileName)
	destFile := filepath.Join(cwd, fileName)

	if _, err := os.Stat(destFile); !os.IsNotExist(err) {
		shouldReplace := false
		err := survey.AskOne(
			&survey.Confirm{Message: "The file already exists, do you want to overwrite it?"},
			&shouldReplace,
			survey.WithValidator(survey.Required),
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if shouldReplace {
			err = overwriteFile(sourceFile, destFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("Template %q copied to current directory", fileName)
		}
	} else if os.IsNotExist(err) {
		err = copyFile(sourceFile, destFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Template %q copied to current directory", fileName)
	}
}

// Create a template from a file in the current directory
func copyToTemplate() {
	cwd, _ := os.Getwd()
	fileName := pickFile(cwd, "Choose a file to make a template from")
	sourceFile := filepath.Join(cwd, fileName)
	destFile := filepath.Join(templatesDir(), fileName)

	if _, err := os.Stat(destFile); !os.IsNotExist(err) {
		shouldReplace := false
		err := survey.AskOne(
			&survey.Confirm{Message: "A template with this name already exists, do you want to overwrite it?"},
			&shouldReplace,
			survey.WithValidator(survey.Required),
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if shouldReplace {
			err = overwriteFile(sourceFile, destFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("New template, %q, was created", fileName)
			return
		}
	} else if os.IsNotExist(err) {
		err = copyFile(sourceFile, destFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("New template, %q, was created", fileName)
		return
	} else {
		fmt.Println("Not sure how we got here...\nMaybe the file exists?")
		os.Exit(1)
	}
}

// Delete an existing template
func deleteTemplate() {
	fileName := pickFile(templatesDir(), "Choose a template to delete")

	shouldDelete := false
	err := survey.AskOne(
		&survey.Confirm{Message: "Are you sure you want to delete this template?"},
		&shouldDelete,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if shouldDelete {
		err := os.Remove(filepath.Join(templatesDir(), fileName))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Template, %q, deleted", fileName)
	}

}

// Edit an existing template
func editTemplate() {
	fileName := pickFile(templatesDir(), "Choose a template to edit")
	path := filepath.Join(templatesDir(), fileName)
	editFile(path)
}

// Make a new template and open it in `$EDITOR`
func newTemplate() {
	fileName := ""
	survey.AskOne(&survey.Input{Message: "Enter a name for this template"}, &fileName)
	filePath := filepath.Join(templatesDir(), fileName)

	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		shouldReplace := false
		err := survey.AskOne(
			&survey.Confirm{Message: "A template with this name already exists, do you want to overwrite it?"},
			&shouldReplace,
			survey.WithValidator(survey.Required),
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if shouldReplace {
			err = ioutil.WriteFile(filePath, []byte{}, os.FileMode(0644))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	} else if os.IsNotExist(err) {
		err = ioutil.WriteFile(filePath, []byte{}, os.FileMode(0644))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	editFile(filePath)
}

// Run if no templates exist
func makeFirstTemplate() {
	firstFile := false
	err := survey.AskOne(
		&survey.Confirm{Message: "Create a new template?"},
		&firstFile,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if firstFile {
		cwd, _ := os.Getwd()
		fileName := pickFile(cwd, "Choose a file to create a template from")

		sourceFile, _ := filepath.Abs(fileName)
		destFile := filepath.Join(templatesDir(), fileName)
		err = copyFile(sourceFile, destFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("New template, %q, was created", fileName)
	}
}

// Choose a file from a given directory
func pickFile(dir, message string) string {
	fileName := ""
	err := survey.AskOne(
		&survey.Select{
			Message: message,
			Options: readDir(dir),
		},
		&fileName,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return fileName
}

// Output of `tmpl --help/-h`
func usage() {
	fmt.Println(titleString())
	fmt.Printf("Usage: %s [OPTIONS] argument\n", os.Args[0])
	flag.PrintDefaults()
}
