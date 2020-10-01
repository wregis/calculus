package ods

import "encoding/xml"

type documentSettings struct {
	XMLName  xml.Name `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 document-settings"`
	Version  string   `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 version,attr"`
	Settings settings `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 settings"`
}

type settings struct {
	ConfigItemSet []configItemSet `xml:"urn:oasis:names:tc:opendocument:xmlns:config:1.0 config-item-set"`
}

type configItemSet struct {
	Name        string       `xml:"urn:oasis:names:tc:opendocument:xmlns:config:1.0 name,attr"`
	ConfigItems []configItem `xml:"urn:oasis:names:tc:opendocument:xmlns:config:1.0 config-item"`
}

type configItem struct {
	Name  string `xml:"urn:oasis:names:tc:opendocument:xmlns:config:1.0 name,attr"`
	Type  string `xml:"urn:oasis:names:tc:opendocument:xmlns:config:1.0 type,attr"`
	Value string `xml:",chardata"`
}
