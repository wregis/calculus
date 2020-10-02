package xlsx

import (
	"archive/zip"
	"encoding/xml"

	"github.com/wregis/calculus"
)

func writeStyles(writer *zip.Writer, workbook calculus.Workbook) error {
	numFmts := []numberingFormat{}
	for _, sheet := range workbook.Sheets() {
		sheet.Rows().Walk(func(_ int, row calculus.Row) {
			row.Walk(func(_ int, cell calculus.Cell) {
				if style := cell.Style(); style != nil {
					if style.NumberFormat != "" {
						for _, numFmt := range numFmts {
							if numFmt.FormatCode == style.NumberFormat {
								return
							}
						}
					}
					numFmts = append(numFmts, numberingFormat{
						ID:         164 + uint16(len(numFmts)),
						FormatCode: style.NumberFormat,
					})
				}
			})
		})
	}
	return writeXMLToFile(writer, "xl/styles.xml", stylesheet{
		NumberingFormats: numberingFormats{Count: uint16(len(numFmts)), Formats: numFmts},
		// TODO
		Fonts: fonts{Count: 1, Fonts: []font{
			// {Color: color{RGB: "FF000000"}, Size: fontSize{Value: 10}, Name: fontName{Value: "Arial"}},
		}},
		// TODO
		CellXfs: cellXfs{Count: 1, Xfs: []xf{
			// 	{NumFmtID: 0, FontID: 0, FillID: 0, BorderID: 0},
		}},
	})
}

type stylesheet struct {
	XMLName          xml.Name         `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main styleSheet"`
	NumberingFormats numberingFormats `xml:"numFmts"`
	Fonts            fonts            `xml:"fonts"`
	Fills            fills            `xml:"fills"`
	Borders          borders          `xml:"borders"`
	CellStyleXfs     cellStyleXfs     `xml:"cellStyleXfs"`
	CellXfs          cellXfs          `xml:"cellXfs"`
	CellStyles       cellStyles       `xml:"cellStyles"`
}

type numberingFormats struct {
	Count   uint16            `xml:"count,attr"`
	Formats []numberingFormat `xml:"numFmt"`
}

type numberingFormat struct {
	ID         uint16 `xml:"numFmtId,attr"`
	FormatCode string `xml:"formatCode,attr"`
}

type fonts struct {
	Count uint16 `xml:"count,attr"`
	Fonts []font `xml:"font"`
}

type font struct {
	Color         color       `xml:"color"`
	Size          fontSize    `xml:"sz"`
	Name          fontName    `xml:"name"`
	Bold          fontBoolean `xml:"b"`
	Italic        fontBoolean `xml:"i"`
	Strikethrough fontBoolean `xml:"s"`
	Underline     fontBoolean `xml:"u"`
}

type color struct {
	RGB string `xml:"rgb,attr,omitempty"`
}

type fontSize struct {
	Value float32 `xml:"val,attr"`
}

type fontName struct {
	Value string `xml:"val,attr"`
}

type fontBoolean bool

func (f fontBoolean) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if f {
		e.EncodeElement("", start)
	}
	return nil
}

func (f *fontBoolean) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	_, err := d.Token()
	if err != nil {
		return err
	}
	*f = true
	return nil
}

type fills struct {
	Count uint16 `xml:"count,attr"`
	Fills []fill `xml:"fill"`
}

type fill struct {
	Pattern patternFill `xml:"patternFill"`
}

type patternFill struct {
	PatternType string `xml:"patternType,attr"`
}

type borders struct {
	Border string `xml:"border"`
	Count  uint16 `xml:"count,attr"`
}

type cellStyleXfs struct { // TODO improve name
	Count uint16 `xml:"count,attr"`
	Xfs   []xf   `xml:"xf"`
}

type cellXfs struct { // TODO improve name
	Count uint16 `xml:"count,attr"`
	Xfs   []xf   `xml:"xf"`
}

type xf struct { // TODO improve name
	NumFmtID uint16 `xml:"numFmtId,attr"`
	FontID   uint16 `xml:"fontId,attr"`
	FillID   uint16 `xml:"fillId,attr"`
	BorderID uint16 `xml:"borderId,attr"`
}

type cellStyles struct {
	Count      uint16      `xml:"count,attr"`
	CellStyles []cellStyle `xml:"cellStyle"`
}

type cellStyle struct {
	Name      string `xml:"name,attr"`
	XfID      uint16 `xml:"xfId,attr"`
	BuiltinID uint16 `xml:"builtinId,attr"`
}
