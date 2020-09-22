package xlsx

import (
	"archive/zip"
	"encoding/xml"
	"fmt"

	"github.com/wregis/calculus"
)

func writeContentTypes(writer *zip.Writer, workbook calculus.Workbook) error {
	types := contentTypes{
		Defaults: []contentType{
			{Extension: "xml", ContentType: "application/xml"},
			{Extension: "rels", ContentType: "application/vnd.openxmlformats-package.relationships+xml"},
			{Extension: "png", ContentType: "image/png"},
			{Extension: "jpeg", ContentType: "image/jpeg"},
		},
		Overrides: []contentType{
			{PartName: "/xl/workbook.xml", ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"},
			{PartName: "/xl/sharedStrings.xml", ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sharedStrings+xml"},
			{PartName: "/docProps/app.xml", ContentType: "application/vnd.openxmlformats-officedocument.extended-properties+xml"},
			{PartName: "/docProps/core.xml", ContentType: "application/vnd.openxmlformats-package.core-properties+xml"},
			{PartName: "/xl/styles.xml", ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml"},
		},
	}
	if sheets := workbook.Sheets(); sheets != nil {
		for index := range sheets {
			types.Overrides = append(types.Overrides, contentType{
				ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml",
				PartName:    fmt.Sprintf("/xl/worksheets/sheet%d.xml", index+1),
			})
		}
	}
	return writeXMLToFile(writer, "[Content_Types].xml", types)
}

type contentTypes struct {
	XMLName   xml.Name      `xml:"http://schemas.openxmlformats.org/package/2006/content-types Types"`
	Defaults  []contentType `xml:"Default"`
	Overrides []contentType `xml:"Override"`
}

type contentType struct {
	Extension   string `xml:"Extension,attr,omitempty"`
	PartName    string `xml:"PartName,attr,omitempty"`
	ContentType string `xml:"ContentType,attr"`
}
