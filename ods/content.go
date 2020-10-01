package ods

import "encoding/xml"

type documentContent struct {
	XMLName         xml.Name        `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 document-content"`
	Version         string          `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 version,attr"`
	Scripts         scripts         `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 scripts,attr"`
	FontFaceDecls   fontFaceDecls   `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 font-face-decls"`
	AutomaticStyles automaticStyles `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 automatic-styles"`
	Body            body            `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 body"`
}

type scripts struct{}

type body struct {
	Spreadsheet spreadsheet `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 spreadsheet"`
}

type spreadsheet struct {
	Tables              []table             `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 table"`
	CalculationSettings calculationSettings `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 calculation-settings"`
}

type calculationSettings struct {
	CaseSensitive         bool            `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 case-sensitive,attr"`
	AutomaticFindLabels   bool            `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 automatic-find-labels,attr"`
	UseRegularExpressions bool            `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 use-regular-expressions,attr"`
	UseWildcards          bool            `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 use-wildcards,attr"`
	Iterationstruct       iterationstruct `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 iteration"`
}

type iterationstruct struct {
	MaximumDifference float64 `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 maximum-difference,attr"`
}

type table struct {
	Name      string        `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 name,attr"`
	StyleName string        `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 style-name,attr"`
	Columns   []tableColumn `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 table-column"`
	Rows      []tableRow    `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 table-row"`
}

type tableColumn struct {
	StyleName string `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 style-name,attr"`
}

type tableRow struct {
	StyleName string      `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 style-name,attr"`
	Cells     []tableCell `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 table-cell"`
}

type tableCell struct {
	ColumnsRepeated uint   `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 number-columns-repeated,attr,omitempty"`
	StyleName       string `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 style-name,attr,omitempty"`
	Formula         string `xml:"urn:oasis:names:tc:opendocument:xmlns:table:1.0 formula,attr,omitempty"`
	ValueType       string `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 value-type,attr,omitempty"`
	CalcValueType   string `xml:"urn:org:documentfoundation:names:experimental:calc:xmlns:calcext:1.0 value-type,attr,omitempty"`
	Value           string `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 value,attr,omitempty"`
	TimeValue       string `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 time-value,attr,omitempty"`
	DateValue       string `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 date-value,attr,omitempty"`
	Text            string `xml:"urn:oasis:names:tc:opendocument:xmlns:text:1.0 p,omitempty"`
}
