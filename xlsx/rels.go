package xlsx

import (
	"archive/zip"
	"encoding/xml"
	"fmt"

	"github.com/wregis/calculus"
)

func writeRelationships(writer *zip.Writer, _ calculus.Workbook) error {
	relationships := relationships{
		Relationships: []relationship{
			{
				Target: "xl/workbook.xml",
				Type:   "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument",
				ID:     "rId1",
			},
		},
	}
	return writeXMLToFile(writer, "_rels/.rels", relationships)
}

func writeWorkbookRelationships(writer *zip.Writer, workbook calculus.Workbook) error {
	relationships := relationships{
		Relationships: []relationship{
			{
				Target: "styles.xml",
				Type:   "http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles",
				ID:     "rId1",
			},
			{
				Target: "sharedStrings.xml",
				Type:   "http://schemas.openxmlformats.org/officeDocument/2006/relationships/sharedStrings",
				ID:     "rId2",
			},
		},
	}
	if sheets := workbook.Sheets(); sheets != nil {
		for index := range sheets {
			relationships.Relationships = append(relationships.Relationships, relationship{
				Target: fmt.Sprintf("worksheets/sheet%d.xml", index+1),
				Type:   "http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet",
				ID:     fmt.Sprintf("rId%d", index+1+3),
			})
		}
	}
	return writeXMLToFile(writer, "xl/_rels/workbook.xml.rels", relationships)
}

type relationships struct {
	XMLName       xml.Name       `xml:"http://schemas.openxmlformats.org/package/2006/relationships Relationships"`
	Relationships []relationship `xml:"Relationship"`
}

type relationship struct {
	Target string `xml:"Target,attr"`
	Type   string `xml:"Type,attr"`
	ID     string `xml:"Id,attr"`
}
