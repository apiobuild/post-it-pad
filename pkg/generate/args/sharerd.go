package args

import (
	"html/template"
)

type ArgsI interface {
	Process() (err error)
}

// SharedArgs ...
type SharedArgs struct {
	FromEmail                     EmailAddress  `json:"from"`
	ToEmail                       EmailAddress  `json:"to"`
	Subject                       string        `json:"subject"`
	HeaderImageURL                string        `json:"header_image_url"`
	TopContentMarkdown            Markdown      `json:"top_content"`
	TopContentHTML                template.HTML `json:"-"`
	EndContentMarkdown            Markdown      `json:"end_content"`
	EndContentHTML                template.HTML `json:"-"`
	ContactInfo                   *ContactInfo  `json:"contact_info"`
	UnsubscribeDisclaimerMarkdown Markdown      `json:"unsubscribe_disclaimer"`
	UnsubscribeDisclaimerHTML     template.HTML `json:"-"`
}

func (s *SharedArgs) Process() (err error) {
	s.TopContentHTML, err = s.TopContentMarkdown.process()
	if err != nil {
		return
	}
	s.EndContentHTML, err = s.EndContentMarkdown.process()
	if err != nil {
		return
	}
	s.UnsubscribeDisclaimerHTML, err = s.UnsubscribeDisclaimerMarkdown.process()
	if err != nil {
		return
	}
	return
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
	Youtube     string `json:"youtube,omitempty"`
	Twitter     string `json:"twitter,omitempty"`
	Instagram   string `json:"instagram,omitempty"`
	PhoneNumber string `json:"phone,omitempty"`
	Name        string `json:"name"`
	Website     string `json:"website,omitempty"`
}
