package generate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"

	"github.com/apiobuild/post-it-pad/pkg/generate/args"
)

// LayoutNameEnum is the layout name enum type
type LayoutNameEnum string

// Layout Names
const (
	Receipt        LayoutNameEnum = "receipt"
	FeatureUpdates LayoutNameEnum = "feature-updates"
	Reply          LayoutNameEnum = "reply"
)

var (
	Layouts = []LayoutNameEnum{
		Receipt,
		FeatureUpdates,
		// Reply,
	}
)

// Default constants
const (
	defaultHTML          = "default.html"
	defaultArgsJSONFile  = "args.json"
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

func (g *Generator) getSharedTemplate(layoutName string) (t *template.Template, err error) {
	t, err = template.New(defaultHTML).Funcs(FuncMap).ParseGlob(path.Join(g.LayoutDir, sharedDir, "*.html"))

	if err != nil {
		g.getLogFields(err).Errorf("Error parsing shared layout template")
		err = templateErrorGenerator(layoutName, ParseLayoutError)
		return
	}
	return
}

func (g Generator) getOrDefaultArgsPath(layoutName string) (argsPath string) {
	if g.ArgsPath == nil {
		argsPath = path.Join(g.LayoutDir, layoutName, defaultArgsJSONFile)
		g.ArgsPath = &argsPath
		return
	}
	argsPath = *g.ArgsPath
	return
}

func (g Generator) getArgsBytesFromFile(layoutName string) (b []byte, err error) {
	var jsonFile *os.File
	argsPath := g.getOrDefaultArgsPath(layoutName)
	g.getLogFields(nil).Infof("Reading args json file from %s", argsPath)
	jsonFile, err = os.Open(argsPath)
	if err != nil {
		g.getLogFields(err).Info("Error reading args json file")
		return
	}
	defer jsonFile.Close()
	b, _ = ioutil.ReadAll(jsonFile)
	return
}

func (g *Generator) loadArgs(layoutName string, args args.ArgsI) (err error) {
	var b []byte
	if g.ArgsJSON != nil {
		g.getLogFields(nil).Info("Reading args json from input")
		b = []byte(*g.ArgsJSON)
	} else {
		if b, err = g.getArgsBytesFromFile(layoutName); err != nil {
			return
		}
	}

	if err = json.Unmarshal(b, &args); err != nil {
		return
	}
	if err = args.Process(); err != nil {
		return
	}
	g.Args = args
	return
}

func layoutNameArgsStructLookup(layoutName string) (argsVal args.ArgsI, err error) {
	switch layoutName {
	case string(Receipt):
		argsVal = &args.ReceiptArgs{}
	case string(FeatureUpdates):
		argsVal = &args.FeatureUpdatesArgs{}
	}
	return
}

// GetTemplateByLayout reads layout and shared directory to create base template
func (g *Generator) GetTemplateByLayout(layoutName string) (err error) {
	g.getLogFields(nil).Info("Generate layout ", layoutName)

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

	args, err := layoutNameArgsStructLookup(layoutName)
	if err != nil {
		g.getLogFields(err).Error("Error processing args")
		return
	}
	if err = g.loadArgs(layoutName, args); err != nil {
		g.getLogFields(err).Error("Error loading args")
		return
	}

	g.getLogFields(nil).Infof("Generate with args: %v", g.Args)

	g.HTML = new(bytes.Buffer)
	if err = t.ExecuteTemplate(g.HTML, defaultHTML, g.Args); err != nil {
		g.getLogFields(err).Error("Error creating template")
		err = templateErrorGenerator(layoutName, ParseLayoutError)
		return
	}
	return
}
