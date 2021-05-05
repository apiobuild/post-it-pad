package generate

import (
	"fmt"
	"html/template"
	"os"
	"path"
)

// Default constants
const (
	defaultHTML          = "default.html"
	sharedDir            = "shared"
	LayoutDirNotFound    = "Layout Directory Not Found"
	LayoutNotFoundError  = "Layout Not Found"
	ParseLayoutError     = "Error Parsing Layout"
	ExecuteTemplateError = "Error Execute Template"
)

type templateError struct {
	layoutName *string
	message    string
}

func (e *templateError) Error() string {
	if e.layoutName != nil {
		return fmt.Sprintf("Template Error: %s. Layout Name: %s.", e.message, *e.layoutName)
	}
	return fmt.Sprintf("Template Error: %s.", e.message)
}

func templateErrorGenerator(layoutName string, reasonCode string) error {
	return &templateError{layoutName: &layoutName, message: reasonCode}
}

func checkLayoutDir(layoutDir string) (err error) {
	if _, e := os.Stat(layoutDir); os.IsNotExist(e) {
		err = &templateError{layoutName: nil, message: LayoutDirNotFound}
	}
	return
}

func (g Generator) getLayoutTemplateDefault(layoutName string) (templateName string) {
	templateName = path.Join(g.LayoutDir, sharedDir, defaultHTML)
	if _, err := os.Stat(path.Join(g.LayoutDir, layoutName, defaultHTML)); os.IsExist(err) {
		g.getLogFields(err).Info("Custom layout default.html found")
		templateName = path.Join(g.LayoutDir, layoutName, defaultHTML)
		err = nil
	}
	return
}

func (g *Generator) getSharedTemplate(layoutName string) (t *template.Template, err error) {
	t, err = template.New(defaultHTML).Funcs(FuncMap).ParseGlob(path.Join(g.LayoutDir, sharedDir, "*.html"))

	if err != nil {
		g.getLogFields(err).Errorf("Error parsing shared layout template")
		err = templateErrorGenerator(layoutName, ParseLayoutError)
		return
	}
	return
}

// GetTemplateByLayout reads layout and shared directory to create base template
func (g Generator) GetTemplateByLayout(layoutName string) (err error) {
	if err = checkLayoutDir(g.LayoutDir); err != nil {
		g.getLogFields(err).Error("Base layout directory not found")
		err = templateErrorGenerator(layoutName, LayoutNotFoundError)
		return
	}

	if _, err = os.Stat(path.Join(g.LayoutDir, layoutName)); os.IsNotExist(err) {
		g.getLogFields(err).Warn("Layout directory not found")
		err = templateErrorGenerator(layoutName, LayoutNotFoundError)
		return
	}

	baseTemplate, err := g.getSharedTemplate(layoutName)
	if err != nil {
		return
	}

	t, err := template.Must(baseTemplate.Clone()).ParseGlob(path.Join(g.LayoutDir, layoutName, "*.html"))
	if err != nil {
		g.getLogFields(err).Error("Error creating template")
		err = templateErrorGenerator(layoutName, ParseLayoutError)
		return
	}

	if err = t.ExecuteTemplate(g.HTML, defaultHTML, g.Args); err != nil {
		g.getLogFields(err).Error("Error creating template")
		err = templateErrorGenerator(layoutName, ParseLayoutError)
		return
	}
	return
}
