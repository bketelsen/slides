package cmd

import (
	"io"
	"io/ioutil"
	"strings"
)

// Slide represents a slide deck and its metadata
type Slide struct {
	Date         string
	EventName    string
	EventURL     string
	VideoURL     string
	MarkdownFile string
	HTMLFile     string
	Title        string
	Image        string
	ImageAlt     string
	Source       string
	Twitter      string
}

// FromReader creates a slide from an io.Reader
func FromReader(r io.Reader) (*Slide, error) {
	s := &Slide{}
	bb, err := ioutil.ReadAll(r)
	if err != nil {
		return s, err
	}
	s.Source = string(bb)
	s.Process()
	return s, nil
}

// Process runs macros and sets struct fields
// for metadata
func (s *Slide) Process() {
	s.getMetadata()
	s.Source = faReplace(string(s.Source))
}

func (s *Slide) getMetadata() {
	lines := strings.Split(s.Source, "\n")
	s.Title = strings.Replace(lines[0], "#", "", -1)
	for _, line := range lines {
		// [twitter]: # (@bketelsen)
		if strings.HasPrefix(line, "[twitter]") {
			s.Twitter = getValue(line)
		}
		if strings.HasPrefix(line, "[event]") {
			s.EventName = getValue(line)
		}
		if strings.HasPrefix(line, "[eventurl]") {
			s.EventURL = getValue(line)
		}
		if strings.HasPrefix(line, "[title]") {
			s.Title = getValue(line)
		}
		if strings.HasPrefix(line, "[image]") {
			s.Image = getValue(line)
		}
		if strings.HasPrefix(line, "[imagealt]") {
			s.ImageAlt = getValue(line)
		}
		if strings.HasPrefix(line, "[date]") {
			s.Date = getValue(line)
		}
		if strings.HasPrefix(line, "[videourl]") {
			s.VideoURL = getValue(line)
		}
	}
}

func getValue(s string) string {
	i := strings.Index(s, "#")
	s = s[i+1:]
	s = strings.TrimSpace(s)
	s = strings.TrimLeft(s, "(")
	s = strings.TrimRight(s, ")")
	return s
}
