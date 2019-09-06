package tmpl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/viper"
)

// Templates models a user's existing templates.
type Templates struct {
	T []Template `json:"templates"`
}

// Add a new template.
func (tmpls *Templates) Add(newTmpl Template) error {
	if tmpls.NameExists(newTmpl.Name) {
		return fmt.Errorf("A template named %q already exists", newTmpl.Name)
	}
	tmpls.T = append(tmpls.T, newTmpl)
	err := tmpls.Save()
	if err != nil {
		return err
	}
	return nil
}

// Remove an existing template, itendified by its name.
func (tmpls *Templates) Remove(name string) error {
	templates := make([]Template, len(tmpls.T)-1)
	if !tmpls.NameExists(name) {
		return fmt.Errorf("No template named %q found", name)
	}
	i := 0
	for _, tmpl := range tmpls.T {
		if tmpl.Name != name {
			templates[i] = tmpl
			i++
		}
	}
	tmpls.T = templates
	err := tmpls.Save()
	if err != nil {
		return err
	}
	return nil
}

// GetByName returns a template with a given name and an error if no template is found.
func (tmpls Templates) GetByName(name string) (Template, error) {
	for _, tmpl := range tmpls.T {
		if tmpl.Name == name {
			return tmpl, nil
		}
	}
	return Template{}, fmt.Errorf("GetByName: no template with that name found")

}

// NameExists checks to see if a given name is already being used for a template.
func (tmpls *Templates) NameExists(name string) bool {
	for _, tmpl := range tmpls.T {
		if tmpl.Name == name {
			return true
		}
	}

	return false
}

// Save templates to a user's templates.json file.
func (tmpls Templates) Save() error {
	dataDir := viper.GetString("templatesdir")
	dataFile := filepath.Join(dataDir, "templates.json")

	bytes, err := json.MarshalIndent(tmpls, "", "  ")
	if err != nil {
		return err
	}

	err = CreateOrOverwrite(dataFile, bytes)
	if err != nil {
		return err
	}

	return nil
}

// GetTemplates returns the user's templates.
func GetTemplates() (Templates, error) {
	var tmpls Templates
	tmplDir := viper.GetString("templatesdir")
	tmplFile := filepath.Join(tmplDir, "templates.json")

	bytes, err := ioutil.ReadFile(tmplFile)
	if err != nil {
		// Return nil if tmplFile doesn't exist yet.
		// We'll make on Templates.Save().
		return tmpls, nil
	}

	err = json.Unmarshal(bytes, &tmpls)
	if err != nil {
		return tmpls, err
	}

	return tmpls, nil
}

// Names returns the names of all templates.
func (tmpls Templates) Names() []string {
	names := make([]string, len(tmpls.T))
	for i, tmpl := range tmpls.T {
		names[i] = tmpl.Name
	}
	return names
}
