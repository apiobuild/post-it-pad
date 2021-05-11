package args

import "html/template"

// FeatureUpdatesArgs is the default args for feature-updates layout
type FeatureUpdatesArgs struct {
	SharedArgs
	Announcements []Announcement `json:"announcements"`
}

// Announcement ...
type Announcement struct {
	ImageURL    string        `json:"image_url"`
	Title       string        `json:"title"`
	Description template.HTML `json:"description"`
	Link        string        `json:"link"`
}
