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

// Update an existing template.
func (ts *Templates) Update(t Template) error {
	for i, tmpl := range ts.T {
		if tmpl.Name == t.Name {
			ts.T[i] = t
		}
	}
	err := ts.Save()
	if err != nil {
		return err
	}
	return nil
}

// Add a new template.
func (ts *Templates) Add(newTmpl Template) error {
	if ts.NameExists(newTmpl.Name) {
		return fmt.Errorf("A template named %q already exists", newTmpl.Name)
	}
	ts.T = append(ts.T, newTmpl)
	err := ts.Save()
	if err != nil {
		return err
	}
	return nil
}

// Remove an existing template, itendified by its name.
func (ts *Templates) Remove(name string) error {
	templates := make([]Template, len(ts.T)-1)
	if !ts.NameExists(name) {
		return fmt.Errorf("No template named %q found", name)
	}
	i := 0
	for _, tmpl := range ts.T {
		if tmpl.Name != name {
			templates[i] = tmpl
			i++
		}
	}
	ts.T = templates
	err := ts.Save()
	if err != nil {
		return err
	}
	return nil
}

// GetByName returns a template with a given name and an error if no template is found.
func (ts Templates) GetByName(name string) (Template, error) {
	for _, tmpl := range ts.T {
		if tmpl.Name == name {
			return tmpl, nil
		}
	}
	return Template{}, fmt.Errorf("GetByName: no template with that name found")
}

// NameExists checks to see if a given name is already being used for a template.
func (ts *Templates) NameExists(name string) bool {
	for _, tmpl := range ts.T {
		if tmpl.Name == name {
			return true
		}
	}

	return false
}

// Save templates to a user's templates.json file.
func (ts Templates) Save() error {
	dataDir := viper.GetString("templatesdir")
	dataFile := filepath.Join(dataDir, "templates.json")

	bytes, err := json.MarshalIndent(ts, "", "  ")
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
	var ts Templates
	tmplDir := viper.GetString("templatesdir")
	tmplFile := filepath.Join(tmplDir, "templates.json")

	bytes, err := ioutil.ReadFile(tmplFile)
	if err != nil {
		// Return nil if tmplFile doesn't exist yet.
		// We'll make on Templates.Save().
		return ts, nil
	}

	err = json.Unmarshal(bytes, &ts)
	if err != nil {
		return ts, err
	}

	return ts, nil
}

// Names returns the names of all templates.
func (ts Templates) Names() []string {
	names := make([]string, len(ts.T))
	for i, tmpl := range ts.T {
		names[i] = tmpl.Name
	}
	return names
}
