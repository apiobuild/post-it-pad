package args

import "html/template"

// SharedArgs ...
type SharedArgs struct {
	FromEmail      EmailAddress  `json:"from"`
	ToEmail        EmailAddress  `json:"to"`
	Subject        string        `json:"subject"`
	HeaderImageURL string        `json:"header_image_url"`
	ContentHTML    template.HTML `json:"content_html"`
	ContactInfo    *ContactInfo  `json:"contact_info"`
}

// EmailAddress ...
type EmailAddress struct {
	Name    string `json:"name"`
	Address string `json:"email"`
}

// Option ...
type Option struct {
	Color       string `json:"color,omitempty"`
	ReplyText   string `json:"option_text,omitempty"`
	URL         string `json:"option_url,omitempty"`
	Description string `json:"display_text,omitempty"`
}

// ContactInfo ...
type ContactInfo struct {
	Logo        string `json:"logo_url,omitempty"`
	Line        string `json:"line,omitempty"`
	Email       string `json:"email,omitempty"`
	Facebook    string `json:"facebook,omitempty"`
	Twitter     string `json:"twitter,omitempty"`
	Instagram   string `json:"instagram,omitempty"`
	PhoneNumber string `json:"phone,omitempty"`
	Name        string `json:"name"`
	Website     string `json:"website,omitempty"`
}
