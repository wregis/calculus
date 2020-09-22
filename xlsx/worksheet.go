package xlsx

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/wregis/calculus"
)

type dataType string

const (
	dataTypeBoolean      dataType = "b"
	dataTypeDate         dataType = "d"
	dataTypeError        dataType = "e"
	dataTypeInlineString dataType = "inlineStr"
	dataTypeNumber       dataType = "n"
	dataTypeSharedString dataType = "s"
	dataTypeFormula      dataType = "str"
)

type cellFormulaType string

const (
	cellFormulaTypeNormal cellFormulaType = "normal"
	cellFormulaTypeShared cellFormulaType = "shared"
)

func writeWorksheets(writer *zip.Writer, workbook calculus.Workbook) error {
	sharedStrings := sharedStringTable{}
	for index, sheet := range workbook.Sheets() {
		worksheet := worksheet{}
		maxRow, maxColumn := 0, 0
		sheet.Rows().Walk(func(rIndex int, r calculus.Row) {
			if rIndex > maxRow {
				maxRow = rIndex
			}

			row := row{Reference: rIndex + 1}
			r.Walk(func(cIndex int, c calculus.Cell) {
				if cIndex > maxColumn {
					maxColumn = cIndex
				}

				reference, _ := calculus.Coordinate(rIndex, cIndex)
				cell := cell{CellReference: reference}
				switch c.Type() {
				case calculus.CellValueTypeString:
					cell.DataType = dataTypeSharedString
					cell.CellValue = strconv.Itoa(sharedStrings.Add(c.Value().(string)))
				case calculus.CellValueTypeNumber:
					str, _ := stringifyNumber(c.Value()) // TODO error
					cell.CellValue = str
				case calculus.CellValueTypeBoolean:
					cell.DataType = dataTypeBoolean
					if c.Value().(bool) {
						cell.CellValue = "1"
					} else {
						cell.CellValue = "0"
					}
				case calculus.CellValueTypeFormula:
					cell.DataType = dataTypeFormula
					cell.CellValue = c.Value().(string)
				case calculus.CellValueTypeError:
					cell.DataType = dataTypeError
				}
				row.Records = append(row.Records, cell)
			})
			worksheet.Data.Rows = append(worksheet.Data.Rows, row)
		})
		reference, _ := calculus.Coordinate(maxRow, maxColumn)
		worksheet.Dimension = &dimension{Reference: "A1:" + reference}

		if err := writeXMLToFile(writer, fmt.Sprintf("xl/worksheets/sheet%d.xml", index+1), worksheet); err != nil {
			return err
		}
	}
	return writeXMLToFile(writer, "xl/sharedStrings.xml", sharedStrings)
}

func stringifyNumber(number interface{}) (string, error) {
	switch n := number.(type) {
	case int:
		return strconv.FormatInt(int64(n), 10), nil
	case int8:
		return strconv.FormatInt(int64(n), 10), nil
	case int16:
		return strconv.FormatInt(int64(n), 10), nil
	case int32:
		return strconv.FormatInt(int64(n), 10), nil
	case int64:
		return strconv.FormatInt(n, 10), nil
	case uint:
		return strconv.FormatUint(uint64(n), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(n), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(n), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(n), 10), nil
	case uint64:
		return strconv.FormatUint(n, 10), nil
	case float32:
		return strconv.FormatFloat(float64(n), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(n, 'f', -1, 64), nil
	default:
		return "", calculus.NewErrorf(nil, "Invalid number type: %T", number)
	}
}

type worksheet struct {
	XMLName          xml.Name               `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main worksheet"`
	SheetViews       *sheetViews            `xml:"sheetViews,omitempty"`
	Dimension        *dimension             `xml:"dimension,omitempty"`
	Data             sheetData              `xml:"sheetData"`
	FormatProperties *sheetFormatProperties `xml:"sheetFormatPr,omitempty"`
	MergeCells       *mergeCells            `xml:"mergeCells,omitempty"`
}

type sheetViews struct {
	SheetView sheetView
}

type sheetView struct {
	WorkbookViewID int `xml:"workbookViewId,attr"`
}

type sheetFormatProperties struct {
	CustomHeight     float32 `xml:"customHeight,attr"`
	DefaultColWidth  float32 `xml:"defaultColWidth,attr"`
	DefaultRowHeight float32 `xml:"defaultRowHeight,attr"`
}

type dimension struct {
	Reference string `xml:"ref,attr"`
}

type sheetData struct {
	Rows []row `xml:"row"`
}

type row struct {
	Reference  int    `xml:"r,attr"`
	Spans      string `xml:"spans,attr,omitempty"`
	StyleIndex string `xml:"s,attr,omitempty"`
	Hidden     bool   `xml:"hidden,attr,omitempty"`
	Collapsed  bool   `xml:"collapsed,attr,omitempty"`
	Records    []cell `xml:"c"`
}

type cell struct {
	CellReference string       `xml:"r,attr"`
	CellMetaIndex int          `xml:"cm,attr,omitempty"`
	DataType      dataType     `xml:"t,attr,omitempty"`
	CellValue     string       `xml:"v,omitempty"`
	CellFormula   *cellFormula `xml:"f,omitempty"`
	StyleIndex    string       `xml:"s,attr,omitempty"`
}

type cellFormula struct {
	Value            string          `xml:",chardata"`
	Type             cellFormulaType `xml:"t,attr,omitempty"`
	RangeOfCells     string          `xml:"ref,attr,omitempty"`
	SharedGroupIndex int             `xml:"si,attr,omitempty"`
}

type mergeCells struct {
	Count      int         `xml:"count,attr"`
	MergeCells []mergeCell `xml:"mergeCell"`
}

type mergeCell struct {
	Reference string `xml:"ref,attr"`
}
