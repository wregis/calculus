package gnumeric

import (
	"compress/gzip"
	"encoding/xml"
	"io"
	"strconv"

	"github.com/wregis/calculus"
)

// Read creates a Workbook from an input data on GNumeric format.
func Read(in io.Reader) (calculus.Workbook, error) {
	workbook, err := ReadFrom(in)
	if err != nil {
		return nil, err
	}
	return ParseWorkbook(workbook)
}

// ReadFrom decodes data on GNumeric format.
func ReadFrom(in io.Reader) (*Workbook, error) {
	r, err := gzip.NewReader(in)
	if err != nil {
		return nil, calculus.NewError(err, "Failed to read or decompress data")
	}
	decoder := xml.NewDecoder(r)
	workbook := &Workbook{}
	if err := decoder.Decode(workbook); err != nil {
		return nil, calculus.NewError(err, "Failed to decode XML data")
	}
	return workbook, nil
}

// ParseWorkbook converts a GNumeric workbook object to a Workbook object.
func ParseWorkbook(source *Workbook) (calculus.Workbook, error) {
	var workbook = calculus.New()
	// TODO Attributes
	// TODO read Summary
	// TODO SheetNameIndex
	// TODO Names?
	// TODO Geometry?
	for _, sheet := range source.Sheets.Sheets {
		worksheet, err := workbook.AddSheet(sheet.Name)
		if err != nil {
			return nil, err
		}
		if sheet.Cells.Cells != nil {
			for _, cell := range sheet.Cells.Cells {
				switch cell.Type {
				case ValueTypeString:
					worksheet.SetValue(cell.Row, cell.Column, cell.Value)
				case ValueTypeBoolean:
					worksheet.SetValue(cell.Row, cell.Column, cell.Value == "TRUE")
				case ValueTypeNumber:
					f, err := strconv.ParseFloat(cell.Value, 64)
					if err != nil {
						return nil, calculus.NewError(err, "Failed to parse numeric value")
					}
					worksheet.SetValue(cell.Row, cell.Column, f)
				}
			}
		}
	}
	return workbook, nil
}
