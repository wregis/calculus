package xlsx

import (
	"archive/zip"
	"encoding/xml"
	"time"

	"github.com/wregis/calculus"
)

func writePropertiesApp(writer *zip.Writer, workbook calculus.Workbook) error {
	app := docPropsApp{
		Application: workbook.Properties().Application(),
	}
	if err := writeXMLToFile(writer, "docProps/app.xml", app); err != nil {
		return err
	}

	props := docPropsCore{
		Creator:        workbook.Properties().Creator(),
		LastModifiedBy: workbook.Properties().LastModifiedBy(),
		Title:          workbook.Properties().Title(),
		Subject:        workbook.Properties().Subject(),
		Description:    workbook.Properties().Description(),
		Keywords:       workbook.Properties().Keywords(),
		Category:       workbook.Properties().Category(),
	}
	if created := workbook.Properties().Created(); !created.IsZero() {
		props.Created = &xmlDCTermsDate{
			Type:  "dcterms:W3CDTF",
			Value: created.Format(time.RFC3339),
		}
	}
	if modified := workbook.Properties().Modified(); !modified.IsZero() {
		props.Modified = &xmlDCTermsDate{
			Type:  "dcterms:W3CDTF",
			Value: modified.Format(time.RFC3339),
		}
	}
	return writeXMLToFile(writer, "docProps/core.xml", props)
}

type docPropsApp struct {
	XMLName            xml.Name `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties Properties"`
	Application        string   `xml:",omitempty"`
	ApplicationVersion string   `xml:"AppVersion,omitempty"`
	DocumentSecurity   int      `xml:"DocSecurity,omitempty"`
	Company            string   `xml:",omitempty"`
	ScaleCrop          bool     `xml:",omitempty"`
	LinksUpToDate      bool     `xml:",omitempty"`
	SharedDocument     bool     `xml:"SharedDoc,omitempty"`
	HyperlinksChanged  bool     `xml:",omitempty"`
}

type docPropsCore struct {
	XMLName        xml.Name        `xml:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties coreProperties"`
	Creator        string          `xml:"http://purl.org/dc/elements/1.1/ creator,omitempty"`
	LastModifiedBy string          `xml:"lastModifiedBy,omitempty"`
	Created        *xmlDCTermsDate `xml:"http://purl.org/dc/terms/ created,omitempty"`
	Modified       *xmlDCTermsDate `xml:"http://purl.org/dc/terms/ modified,omitempty"`
	Title          string          `xml:"http://purl.org/dc/elements/1.1/ title,omitempty"`
	Description    string          `xml:"http://purl.org/dc/elements/1.1/ description,omitempty"`
	Subject        string          `xml:"http://purl.org/dc/elements/1.1/ subject,omitempty"`
	Keywords       string          `xml:"http://purl.org/dc/elements/1.1/ keywords,omitempty"`
	Category       string          `xml:"http://purl.org/dc/elements/1.1/ category,omitempty"`
}

type xmlDCTermsDate struct {
	Type  string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	Value string `xml:",chardata"`
}
