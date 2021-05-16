package args_test

import (
	"testing"

	"github.com/apiobuild/post-it-pad/pkg/generate/args"
	"github.com/stretchr/testify/assert"
)

func TestSharedArgs(t *testing.T) {
	s1 := args.SharedArgs{}
	err := s1.Process()
	assert.Nil(t, err)

	s2 := args.SharedArgs{
		TopContentMarkdown: "# hello",
	}
	err = s2.Process()
	assert.Equal(t, string(s2.TopContentHTML), "<h1>hello</h1>\n")
	assert.Nil(t, err)
}

func TestFeatureUpdatesArgs(t *testing.T) {
	s1 := args.FeatureUpdatesArgs{}
	err := s1.Process()
	assert.Nil(t, err)

	s2 := args.FeatureUpdatesArgs{
		Announcements: []args.Announcement{
			{
				DescriptionMarkDown: "# hello",
			},
		},
	}
	err = s2.Process()
	assert.Equal(t, string(s2.Announcements[0].DescriptionHTML), "<h1>hello</h1>\n")
	assert.Nil(t, err)
}

func TestReceiptArgs(t *testing.T) {
	s1 := args.ReceiptArgs{}
	err := s1.Process()
	assert.Nil(t, err)
}
