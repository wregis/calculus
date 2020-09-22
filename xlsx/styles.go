package xlsx

import (
	"archive/zip"
	"encoding/xml"

	"github.com/wregis/calculus"
)

func writeStyles(writer *zip.Writer, workbook calculus.Workbook) error {
	return writeXMLToFile(writer, "xl/styles.xml", stylesheet{
		XMLSpace: "preserve",
		Fonts: fonts{Count: 1, Fonts: []font{
			{Color: color{RGB: "FF000000"}, Size: fontSize{Value: 10}, Name: fontName{Value: "Arial"}},
		}},
		Fills: fills{Count: 2, Fills: []fill{
			{Pattern: patternFill{PatternType: "none"}},
			{Pattern: patternFill{PatternType: "lightGray"}},
		}},
		Borders: borders{Count: 1},
		CellStyleXfs: cellStyleXfs{Count: 1, Xfs: []xf{
			{NumFmtID: 0, FontID: 0, FillID: 0, BorderID: 0},
		}},
		CellXfs: cellXfs{Count: 1, Xfs: []xf{
			{NumFmtID: 0, FontID: 0, FillID: 0, BorderID: 0},
		}},
		CellStyles: cellStyles{Count: 1, CellStyles: []cellStyle{
			{Name: "Normal", XfID: 0, BuiltinID: 0},
		}},
	})
}

type stylesheet struct {
	XMLName      xml.Name         `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main styleSheet"`
	XMLSpace     string           `xml:"xml:space,attr"`
	NumFonts     numberingFormats `xml:"numFmts"`
	Fonts        fonts            `xml:"fonts"`
	Fills        fills            `xml:"fills"`
	Borders      borders          `xml:"borders"`
	CellStyleXfs cellStyleXfs     `xml:"cellStyleXfs"`
	CellXfs      cellXfs          `xml:"cellXfs"`
	CellStyles   cellStyles       `xml:"cellStyles"`
}

type numberingFormats struct {
	Count int `xml:"count,attr"`
}

type fonts struct {
	Count int    `xml:"count,attr"`
	Fonts []font `xml:"font"`
}

type font struct {
	Color color    `xml:"color"`
	Size  fontSize `xml:"sz"`
	Name  fontName `xml:"name"`
}

type color struct {
	RGB string `xml:"rgb,attr,omitempty"`
}

type fontSize struct {
	Value int `xml:"val,attr"`
}

type fontName struct {
	Value string `xml:"val,attr"`
}

type fills struct {
	Count int    `xml:"count,attr"`
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
	Count  int    `xml:"count,attr"`
}

type cellStyleXfs struct { // TODO improve name
	Count int  `xml:"count,attr"`
	Xfs   []xf `xml:"xf"`
}

type cellXfs struct { // TODO improve name
	Count int  `xml:"count,attr"`
	Xfs   []xf `xml:"xf"`
}

type xf struct { // TODO improve name
	NumFmtID int `xml:"numFmtId,attr"`
	FontID   int `xml:"fontId,attr"`
	FillID   int `xml:"fillId,attr"`
	BorderID int `xml:"borderId,attr"`
}

type cellStyles struct {
	Count      int         `xml:"count,attr"`
	CellStyles []cellStyle `xml:"cellStyle"`
}

type cellStyle struct {
	Name      string `xml:"name,attr"`
	XfID      int    `xml:"xfId,attr"`
	BuiltinID int    `xml:"builtinId,attr"`
}
