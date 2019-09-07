package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/tcd/extedit"
	"github.com/tcd/tmpl/tmpl"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing template",
	PreRun: func(cmd *cobra.Command, args []string) {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			log.Fatal(err)
		}
		if debug {
			log.SetFlags(log.Lshortfile)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		IfBoolFlag(cmd, "content", editContent)
		IfBoolFlag(cmd, "name", editName)
		IfBoolFlag(cmd, "file", editFileName)
		IfBoolFlag(cmd, "data", editData)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().Bool("content", false, "Edit a template's contents")
	editCmd.Flags().Bool("file", false, "Edit a template's filename")
	editCmd.Flags().Bool("name", false, "Edit a template's name")
	editCmd.Flags().Bool("data", false, "Edit a template's data")
}

func editContent() {
	ts, err := tmpl.GetTemplates()
	if err != nil {
		log.Fatal(err)
	}

	name := tmpl.PickTemplate("Please choose a template to edit:")
	t, _ := ts.GetByName(name)

	input := strings.NewReader(t.Content)
	diff, err := extedit.Invoke(input)
	if err != nil {
		log.Fatal(err)
	}

	t.Content = diff.Content()
	err = ts.Update(t)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Template %s updated\n", color.FgBlue.Render(t.Name))
}

func editName() {
	ts, err := tmpl.GetTemplates()
	if err != nil {
		log.Fatal(err)
	}

	name := tmpl.PickTemplate("Please choose a template to edit:")
	t, _ := ts.GetByName(name)

	msg := "Please choose a new name:"
	for {
		var newName string
		prompt := &survey.Input{Message: msg}
		survey.AskOne(prompt, &newName)
		if !ts.NameExists(newName) {
			t.Name = name
			break
		}
		msg = "A template that name already exists. Please choose another:"
	}

	err = ts.Update(t)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Template %s updated\n", color.FgBlue.Render(t.Name))
}

func editFileName() {
	ts, err := tmpl.GetTemplates()
	if err != nil {
		log.Fatal(err)
	}

	name := tmpl.PickTemplate("Please choose a template to edit:")
	t, _ := ts.GetByName(name)

	var newName string
	prompt := &survey.Input{Message: "Please choose a new filename:"}
	survey.AskOne(prompt, &newName)

	t.FileName = newName
	err = ts.Update(t)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Template %s updated\n", color.FgBlue.Render(t.Name))
}

func editData() {
	ts, err := tmpl.GetTemplates()
	if err != nil {
		log.Fatal(err)
	}

	name := tmpl.PickTemplate("Please choose a template to edit:")
	t, _ := ts.GetByName(name)

	var action string
	prompt := &survey.Select{
		Message: "What would you like to do?:",
		Options: []string{"add data", "remove data"},
	}
	survey.AskOne(prompt, &action)

	var newData []string
	switch action {
	case "add data":
		newData = addData(t.Data)
	case "remove data":
		newData = removeData(t.Data)
	default:
		os.Exit(1)
	}
	t.Data = newData
	err = ts.Update(t)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Template %s updated\n", color.FgBlue.Render(t.Name))
}

func addData(data []string) []string {
	for {
		var name string
		prompt1 := &survey.Input{Message: "Enter a name for a new variable"}
		survey.AskOne(prompt1, &name)
		if contains(data, name) {
			log.Println("A variable with this name already exists")
		} else {
			data = append(data, name)
		}

		var more bool
		prompt2 := &survey.Confirm{Message: "Add another?"}
		survey.AskOne(prompt2, &more)
		if !more {
			break
		}
	}
	return data
}

func removeData(data []string) []string {
	toRemove := []string{}
	prompt := &survey.MultiSelect{
		Message: "What days do you prefer:",
		Options: data,
	}
	survey.AskOne(prompt, &toRemove)
	return remove(data, toRemove...)
}
