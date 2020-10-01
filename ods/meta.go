package ods

import "encoding/xml"

type documentMeta struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 document-meta"`
	Version string   `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 version,attr"`
	Meta    meta     `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 meta"`
}

type meta struct {
	Generator      string `xml:"urn:oasis:names:tc:opendocument:xmlns:meta:1.0 generator,attr"`
	Title          string `xml:"http://purl.org/dc/elements/1.1/ title,attr"`
	Description    string `xml:"http://purl.org/dc/elements/1.1/ description,attr"`
	Subject        string `xml:"http://purl.org/dc/elements/1.1/ subject,attr"`
	Keyword        string `xml:"urn:oasis:names:tc:opendocument:xmlns:meta:1.0 keyword,attr"`
	InitialCreator string `xml:"urn:oasis:names:tc:opendocument:xmlns:meta:1.0 initial-creator,attr"`
	Creator        string `xml:"http://purl.org/dc/elements/1.1/ creator,attr"`
	CreationDate   string `xml:"urn:oasis:names:tc:opendocument:xmlns:meta:1.0 creation-date,attr"`
	Date           string `xml:"http://purl.org/dc/elements/1.1/ date,attr"`
	Language       string `xml:"http://purl.org/dc/elements/1.1/ language,attr"`
}
