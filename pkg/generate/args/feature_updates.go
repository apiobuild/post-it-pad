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
	ImageURL    string        `json:"image_url"`
	Title       string        `json:"title"`
	Description template.HTML `json:"description"`
	Link        string        `json:"link"`
}
