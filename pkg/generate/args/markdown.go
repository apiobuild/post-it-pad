package args

import (
	"bytes"
	"html/template"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

type Markdown string

func (m Markdown) process() (output template.HTML, err error) {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
		),
	)
	var buf bytes.Buffer
	if err = md.Convert([]byte(m), &buf); err != nil {
		return
	}
	output = template.HTML(buf.String())
	return
}
