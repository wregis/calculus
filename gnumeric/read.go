package gnumeric

import (
	"compress/gzip"
	"encoding/xml"
	"io"
	"strconv"

	"github.com/wregis/calculus"
	"github.com/wregis/calculus/internal/errors"
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
		return nil, errors.New(err, "Failed to read or decompress data")
	}
	decoder := xml.NewDecoder(r)
	workbook := &Workbook{}
	if err := decoder.Decode(workbook); err != nil {
		return nil, errors.New(err, "Failed to decode XML data")
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
				case ValueTypeEmpty:
					worksheet.SetValue(cell.Row, cell.Column, nil)
				case ValueTypeBoolean:
					worksheet.SetValue(cell.Row, cell.Column, cell.Value == "TRUE")
				case ValueTypeInteger:
					i, err := strconv.ParseInt(cell.Value, 10, 64)
					if err != nil {
						return nil, errors.New(err, "Failed to parse integer value")
					}
					worksheet.SetValue(cell.Row, cell.Column, i)
				case ValueTypeFloat:
					f, err := strconv.ParseFloat(cell.Value, 64)
					if err != nil {
						return nil, errors.New(err, "Failed to parse floating point value")
					}
					worksheet.SetValue(cell.Row, cell.Column, f)
				case ValueTypeError: // TODO?
				case ValueTypeString:
					worksheet.SetValue(cell.Row, cell.Column, cell.Value)
				case ValueTypeCellRange: // TODO?
				case ValueTypeArray: // TODO?
				}
			}
		}
	}
	return workbook, nil
}
