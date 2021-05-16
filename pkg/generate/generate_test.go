package generate_test

import (
	"testing"

	"github.com/apiobuild/post-it-pad/pkg/generate"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAll(t *testing.T) {
	layoutDir := "../../layouts"
	g := generate.NewGenerator(
		&layoutDir,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
	err := g.Generate()
	assert.Nil(t, err)
}

func TestGenerateReceipt(t *testing.T) {
	layoutDir := "../../layouts"
	layoutName := "receipt"
	g := generate.NewGenerator(
		&layoutDir,
		&layoutName,
		nil,
		nil,
		nil,
		nil,
	)
	err := g.Generate()
	assert.Nil(t, err)
}

func TestGenerateLayoutError(t *testing.T) {
	layoutDir := "../../layouts"
	layoutName := "does not exist"
	g := generate.NewGenerator(
		&layoutDir,
		&layoutName,
		nil,
		nil,
		nil,
		nil,
	)
	err := g.Generate()
	assert.Error(t, err)
}

func TestGenerateLayoutDirError(t *testing.T) {
	layoutDir := "does not exist"
	g := generate.NewGenerator(
		&layoutDir,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
	err := g.Generate()
	assert.Error(t, err)
}

func TestGenerateJSONArgsError(t *testing.T) {
	layoutDir := "../../layouts"
	argsPath := "does not exist"
	g := generate.NewGenerator(
		&layoutDir,
		nil,
		nil,
		nil,
		&argsPath,
		nil,
	)
	err := g.Generate()
	assert.Error(t, err)
}
