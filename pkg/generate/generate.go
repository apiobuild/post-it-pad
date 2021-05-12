package generate

import (
	"bytes"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

// Generator is the base struct to generate templated email
type Generator struct {
	LayoutDir  string
	LayoutName *string
	DestPath   string
	HTML       *bytes.Buffer
	ArgsPath   *string
	ArgsJSON   *string
	Args       interface{}
}

// Default values when not provided
const (
	DefaultLayoutDir = "./layouts"
	DefaultDestPath  = "./output/generated.html"
	// DefaultArgsPath  = "./args.json"
	// DefaultLayoutDir = "../../layouts"
	// DefaultDestPath = "../../generated.html"
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
func NewGenerator(layoutDir *string, layoutName *string, destPath *string, argsPath *string, argsJSON *string) (g Generator) {
	g = Generator{
		LayoutDir:  getStringOrDefault(layoutDir, DefaultLayoutDir),
		LayoutName: layoutName,
		DestPath:   getStringOrDefault(destPath, DefaultDestPath),
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
	return
}

func (g Generator) generateByLayout() (err error) {
	return g.GetTemplateByLayout(*g.LayoutName)
}

func (g Generator) writeToFile() (err error) {
	g.getLogFields(nil).Info("Writing generated html to file to", g.DestPath)
	// NOTE: 0755: overwrite
	err = ioutil.WriteFile(g.DestPath, g.HTML.Bytes(), 0755)
	return
}

// Generate generates actual templated email html
func (g Generator) Generate() (err error) {
	if g.isGenerateAll() {
		g.getLogFields(nil).Info("Generate for all layouts")
		err = g.generateAll()
	} else {
		g.getLogFields(nil).Info("Generate by layout name specified")
		err = g.generateByLayout()
	}
	if err != nil {
		g.getLogFields(err).Fatal("Error executing email generator")
		return
	}
	if err = g.writeToFile(); err != nil {
		g.getLogFields(err).Fatal("Error writing generated html to file")
		return
	}
	return
}

func (g Generator) getLogFields(err error) *log.Entry {
	var fields map[string]interface{} = map[string]interface{}{}

	fields["LayoutDir"] = g.LayoutDir
	fields["destPath"] = g.DestPath

	if err != nil {
		fields["Err"] = err.Error()
	}
	if g.LayoutName != nil {
		fields["layoutName"] = *g.LayoutName
	}

	return log.WithFields(fields)
}
