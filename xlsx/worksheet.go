package xlsx

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"github.com/wregis/calculus"
	"github.com/wregis/calculus/internal/coordinate"
	"github.com/wregis/calculus/internal/errors"
	"github.com/wregis/calculus/internal/format"
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

				reference, _ := coordinate.Format(rIndex, cIndex)
				cell := cell{CellReference: reference}
				switch c.Type() {
				case calculus.CellValueTypeEmpty:
				case calculus.CellValueTypeBoolean:
					cell.DataType = dataTypeBoolean
					if c.Value().(bool) {
						cell.CellValue = "1"
					} else {
						cell.CellValue = "0"
					}
				case calculus.CellValueTypeInteger:
					str, err := stringifyInteger(c.Value())
					if err != nil {
						cell.DataType = dataTypeError
						break
					}
					cell.CellValue = str
				case calculus.CellValueTypeFloat:
					str, err := stringifyFloat(c.Value())
					if err != nil {
						cell.DataType = dataTypeError
						break
					}
					cell.CellValue = str
				case calculus.CellValueTypeError:
					cell.DataType = dataTypeError
				case calculus.CellValueTypeString:
					cell.DataType = dataTypeSharedString
					cell.CellValue = strconv.Itoa(sharedStrings.Add(c.Value().(string)))
				case calculus.CellValueTypeDate:
					cell.DataType = dataTypeDate
					cell.CellValue = strconv.FormatFloat(format.ToTime1900(
						c.Value().(time.Time),
						workbook.Properties().Date1904(),
					), 'f', -1, 64)
					// TODO cell.Format = "d/m/yyyy"
				}
				row.Records = append(row.Records, cell)
			})
			worksheet.Data.Rows = append(worksheet.Data.Rows, row)
		})
		reference, _ := coordinate.Format(maxRow, maxColumn)
		worksheet.Dimension = &dimension{Reference: "A1:" + reference}

		if err := writeXMLToFile(writer, fmt.Sprintf("xl/worksheets/sheet%d.xml", index+1), worksheet); err != nil {
			return err
		}
	}
	return writeXMLToFile(writer, "xl/sharedStrings.xml", sharedStrings)
}

func stringifyInteger(value interface{}) (string, error) {
	switch v := value.(type) {
	case int:
		return strconv.FormatInt(int64(v), 10), nil
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint64:
		return strconv.FormatUint(v, 10), nil
	}
	return "", errors.New(nil, "Invalid integer value")
}

func stringifyFloat(value interface{}) (string, error) {
	switch v := value.(type) {
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	}
	return "", errors.New(nil, "Invalid floating point value")
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
	DataType      dataType     `xml:"t,attr,omitempty"`
	CellValue     string       `xml:"v,omitempty"`
	CellFormula   *cellFormula `xml:"f,omitempty"`
	StyleIndex    uint16       `xml:"s,attr,omitempty"`
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
