package args

import "html/template"

// StyleEnum is the style enum types
type StyleEnum string

// Styles
const (
	LeftRight  StyleEnum = "left-right"
	LargeImage StyleEnum = "large-image"
)

// FeatureUpdatesArgs is the default args for feature-updates layout
type FeatureUpdatesArgs struct {
	SharedArgs
	Announcements []Announcement `json:"announcements"`
	Style         StyleEnum      `json:"style"`
}

// Announcement ...
type Announcement struct {
	ImageURL            string        `json:"image_url"`
	Title               string        `json:"title"`
	DescriptionMarkDown Markdown      `json:"description"`
	DescriptionHTML     template.HTML `json:"-"`
	Link                string        `json:"link"`
}

func (f *FeatureUpdatesArgs) Process() (err error) {
	f.SharedArgs.Process()
	for i, a := range f.Announcements {
		a.Process()
		f.Announcements[i] = a
	}
	return
}

func (a *Announcement) Process() (err error) {
	a.DescriptionHTML, err = a.DescriptionMarkDown.process()
	if err != nil {
		return
	}
	return
}
