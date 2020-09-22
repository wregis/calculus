package xlsx

import (
	"encoding/xml"
)

// sharedStringTable holds a string dictionary to be refered by worksheets.
type sharedStringTable struct {
	XMLName     xml.Name     `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main sst"`
	Count       int          `xml:"count,attr"`
	UniqueCount int          `xml:"uniqueCount,attr"`
	Strings     []stringItem `xml:"si"`
}

func (s *sharedStringTable) Get(index int) string {
	if len(s.Strings) > index {
		return s.Strings[index].Text
	}
	return ""
}

// Add includes a new string to the dictionary, if not already there, and returns the index for it to be refered.
func (s *sharedStringTable) Add(text string) int {
	if s.Strings != nil {
		for index, value := range s.Strings {
			if value.Text == text {
				return index
			}
		}
	}
	s.Strings = append(s.Strings, stringItem{Text: text})
	s.Count, s.UniqueCount = len(s.Strings), len(s.Strings)
	return s.Count - 1
}

// stringItem is a string dictionary value holder.
type stringItem struct {
	Text string `xml:"t"`
}
