package ods

import "encoding/xml"

type documentStyles struct {
	XMLName         xml.Name        `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 document-styles"`
	Version         string          `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 version,attr"`
	FontFaceDecls   fontFaceDecls   `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 font-face-decls"`
	Styles          styles          `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 styles"`
	AutomaticStyles automaticStyles `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 automatic-styles"`
}

type fontFaceDecls struct {
	FontFaces []fontFace `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 font-face"`
}

type fontFace struct {
	Name              string `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 name,attr"`
	FontFamily        string `xml:"urn:oasis:names:tc:opendocument:xmlns:svg-compatible:1.0 font-family,attr"`
	FontFamilyGeneric string `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 font-family-generic,attr"`
	FontPitch         string `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 font-pitch,attr"`
}

type styles struct{}

type automaticStyles struct {
	Styles []style `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 style"`
}

type style struct {
	Name             string                 `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 name,attr"`
	Family           string                 `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 family,attr"`
	Properties       *tableProperties       `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 table-properties,omitempty"`
	ColumnProperties *tableColumnProperties `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 table-column-properties,omitempty"`
	RowProperties    *tableRowProperties    `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 table-row-properties,omitempty"`
	CellProperties   *cellProperties        `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 table-cell-properties"`
	TextProperties   *textProperties        `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 text-properties"`
}

type tableProperties struct {
	Display     bool   `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 display,attr"`
	WritingMode string `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 writing-mode,attr"`
}

type tableColumnProperties struct {
	BreakBefore string `xml:"urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0 break-before,attr"`
	ColumnWidth string `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 column-width,attr"`
}

type tableRowProperties struct {
	RowHeight           string `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 row-height,attr"`
	BreakBefore         string `xml:"urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0 break-before,attr"`
	UseOptimalRowHeight bool   `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 use-optimal-row-height,attr"`
}

type cellProperties struct {
	RotationAlign string `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 rotation-align"`
}

type textProperties struct {
	Color                string `xml:"urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0 color"`
	TextLutline          bool   `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 text-outline"`
	TextLineThroughStyle string `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 text-line-through-style"`
	TextLineThroughType  string `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 text-line-through-type"`
	FontName             string `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 font-name"`
	FontSize             string `xml:"urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0 font-size"`
	FontStyle            string `xml:"urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0 font-style"`
	TextShadow           string `xml:"urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0 text-shadow"`
	TextUnderlineStyle   string `xml:"urn:oasis:names:tc:opendocument:xmlns:style:1.0 text-underline-style"`
	FontWeight           string `xml:"urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0 font-weight"`
}
