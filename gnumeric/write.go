package gnumeric

import (
	"compress/gzip"
	"encoding/xml"
	"io"
	"strconv"
	"time"

	"github.com/wregis/calculus"
	"github.com/wregis/calculus/internal/errors"
	"github.com/wregis/calculus/internal/format"
)

func Write(workbook calculus.Workbook, out io.Writer) error {
	dest, err := ComposeWorkbook(workbook)
	if err != nil {
		return err
	}
	return WriteTo(dest, out)
}

func ComposeWorkbook(wb calculus.Workbook) (*Workbook, error) {
	workbook := &Workbook{
		Attributes: Attributes{
			Attributes: []Attribute{
				{Name: "WorkbookView::show_horizontal_scrollbar", Value: "TRUE"},
				{Name: "WorkbookView::show_vertical_scrollbar", Value: "TRUE"},
				{Name: "WorkbookView::show_notebook_tabs", Value: "TRUE"},
				{Name: "WorkbookView::is_protected", Value: "FALSE"},
			},
		},
		Summary: Summary{
			Items: []SummaryItems{
				{Name: "application", Value: wb.Properties().Application()},
				{Name: "author", Value: wb.Properties().Creator()},
				{Name: "last_author", Value: wb.Properties().LastModifiedBy()},
				{Name: "title", Value: wb.Properties().Title()},
				{Name: "category", Value: wb.Properties().Category()},
				{Name: "keywords", Value: wb.Properties().Keywords()},
			},
		},
		Geometry: Geometry{
			Width:  734,
			Height: 422,
		},
	}
	for _, s := range wb.Sheets() {
		sheet := Sheet{
			DisplayOutlines:     BooleanTrue,
			OutlineSymbolsBelow: BooleanTrue,
			OutlineSymbolsRight: BooleanTrue,
			Visibility:          VisibilityVisible,
			Name:                s.Name(),
			Zoom:                1.0,
			GridColor:           "0:0:0",
			PrintInformation: PrintInformation{
				Margins: Margins{
					Top:    Margin{Points: 120, PreferredUnit: "mm"},
					Bottom: Margin{Points: 120, PreferredUnit: "mm"},
					Left:   Margin{Points: 72, PreferredUnit: "mm"},
					Right:  Margin{Points: 72, PreferredUnit: "mm"},
					Header: Margin{Points: 72, PreferredUnit: "mm"},
					Footer: Margin{Points: 72, PreferredUnit: "mm"},
				},
				Orientation: OrientationPortrait,
				Paper:       "iso_a4",
			},
			Columns:     Columns{DefaultSizePoints: 48},
			Rows:        Rows{DefaultSizePts: 12.8},
			SheetLayout: SheetLayout{TopLeft: "A1"},
		}
		s.Rows().Walk(func(rIndex int, r calculus.Row) {
			row := RowInfo{
				Number: rIndex,
				Height: r.Height(),
				Hidden: BooleanFalse,
			}
			if r.Hidden() {
				row.Hidden = 1
			}
			sheet.Rows.RowInfo = append(sheet.Rows.RowInfo, row)

			r.Walk(func(cIndex int, c calculus.Cell) {
				cell := Cell{
					Column: cIndex,
					Row:    rIndex,
				}
				switch c.Type() {
				case calculus.CellValueTypeEmpty:
					cell.Type = ValueTypeEmpty
				case calculus.CellValueTypeBoolean:
					cell.Type = ValueTypeBoolean
					if c.Value().(bool) {
						cell.Value = "TRUE"
					} else {
						cell.Value = "FALSE"
					}
				case calculus.CellValueTypeInteger:
					str, err := stringifyInteger(c.Value())
					if err != nil {
						cell.Type = ValueTypeError
						break
					}
					cell.Type = ValueTypeInteger
					cell.Value = str
				case calculus.CellValueTypeFloat:
					str, err := stringifyFloat(c.Value())
					if err != nil {
						cell.Type = ValueTypeError
						break
					}
					cell.Type = ValueTypeFloat
					cell.Value = str
				case calculus.CellValueTypeError:
					cell.Type = ValueTypeError
				case calculus.CellValueTypeString:
					cell.Type = ValueTypeString
					cell.Value = c.Value().(string)
				case calculus.CellValueTypeDate:
					cell.Type = ValueTypeFloat
					cell.Value = strconv.FormatFloat(format.ToTime1900(
						c.Value().(time.Time),
						wb.Properties().Date1904(),
					), 'f', -1, 64)
					cell.Format = "d/m/yyyy" // TODO
				}
				sheet.Cells.Cells = append(sheet.Cells.Cells, cell)
			})
		})
		workbook.Sheets.Sheets = append(workbook.Sheets.Sheets, sheet)
	}
	return workbook, nil
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

func WriteTo(workbook *Workbook, out io.Writer) error {
	writer := gzip.NewWriter(out)
	if _, err := writer.Write([]byte(xml.Header)); err != nil {
		return errors.New(err, "Failed write header")
	}
	xml, err := xml.Marshal(workbook)
	if err != nil {
		return errors.New(err, "Failed to encode XML data")
	}
	if _, err := writer.Write(xml); err != nil {
		return errors.New(err, "Failed to write or compress data")
	}
	err = writer.Flush()
	if err != nil {
		return errors.New(err, "Failed to write file")
	}
	err = writer.Close()
	if err != nil {
		return errors.New(err, "Failed to write file")
	}
	return nil
}
