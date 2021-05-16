package generate

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
)

// Generator is the base struct to generate templated email
type Generator struct {
	LayoutDir  string
	LayoutName *string
	DestPath   string
	DestDir    string
	HTML       *bytes.Buffer
	ArgsPath   *string
	ArgsJSON   *string
	Args       interface{}
}

// Default values when not provided
const (
	// NOTE: for local testing
	// DefaultLayoutDir = "../../layouts"
	DefaultLayoutDir = "./layouts"
	DefaultDestDir   = "./output"
	DefaultDestPath  = "./output/generated.html"
)

func getStringOrDefault(val *string, defaultVal string) (useVal string) {
	if val == nil {
		useVal = defaultVal
	} else {
		useVal = *val
	}
	return
}

// NewGenerator creates new templated email generator
func NewGenerator(layoutDir *string, layoutName *string, destPath *string, destDir *string, argsPath *string, argsJSON *string) (g Generator) {
	g = Generator{
		LayoutDir:  getStringOrDefault(layoutDir, DefaultLayoutDir),
		LayoutName: layoutName,
		DestPath:   getStringOrDefault(destPath, DefaultDestPath),
		DestDir:    getStringOrDefault(destDir, DefaultDestDir),
		ArgsPath:   argsPath,
		ArgsJSON:   argsJSON,
		HTML:       new(bytes.Buffer),
	}
	g.getLogFields(nil).Info("New email generator created")
	return
}

func (g Generator) isGenerateAll() bool {
	return g.LayoutName == nil
}

func (g Generator) generateAll() (err error) {
	for _, layout := range Layouts {
		if err = g.GetTemplateByLayout(string(layout)); err != nil {
			return
		}
		layoutStr := string(layout)
		if err = g.writeToFile(&layoutStr); err != nil {
			g.getLogFields(err).Fatal("Error writing generated html to file")
			return
		}
	}
	return
}

func (g Generator) generateByLayout() (err error) {
	return g.GetTemplateByLayout(*g.LayoutName)
}

func (g Generator) writeToFile(filename *string) (err error) {
	destPath := g.DestPath
	if filename != nil {
		os.MkdirAll(g.DestDir, os.ModePerm)
		destPath = path.Join(g.DestDir, fmt.Sprintf("generated-%s.html", *filename))
	}
	g.getLogFields(nil).Infof("Writing generated html to file to %s", destPath)
	// NOTE: 0755: overwrite
	err = ioutil.WriteFile(destPath, g.HTML.Bytes(), 0755)
	return
}

// Generate generates actual templated email html
func (g Generator) Generate() (err error) {
	if g.isGenerateAll() {
		g.getLogFields(nil).Info("Generate for all layouts")
		err = g.generateAll()
		return
	}
	g.getLogFields(nil).Info("Generate by layout name specified")
	err = g.generateByLayout()
	if err != nil {
		g.getLogFields(err).Fatal("Error executing email generator")
		return
	}
	if err = g.writeToFile(nil); err != nil {
		g.getLogFields(err).Fatal("Error writing generated html to file")
		return
	}
	return
}

func (g Generator) getLogFields(err error) *log.Entry {
	var fields map[string]interface{} = map[string]interface{}{}

	fields["LayoutDir"] = g.LayoutDir
	fields["destDir"] = g.DestDir
	fields["destPath"] = g.DestPath

	if err != nil {
		fields["Err"] = err.Error()
	}
	if g.LayoutName != nil {
		fields["layoutName"] = *g.LayoutName
	}

	return log.WithFields(fields)
}
