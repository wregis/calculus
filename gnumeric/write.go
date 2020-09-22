package gnumeric

import (
	"compress/gzip"
	"encoding/xml"
	"io"
	"strconv"

	"github.com/wregis/calculus"
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
			DisplayOutlines:     Boolean(1),
			OutlineSymbolsBelow: Boolean(1),
			OutlineSymbolsRight: Boolean(1),
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
				Hidden: Boolean(0),
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
				case calculus.CellValueTypeString:
					cell.Type = ValueTypeString
					cell.Value = c.Value().(string)
				case calculus.CellValueTypeNumber:
					n, err := stringifyNumber(c.Value())
					if err != nil {
						return // TODO handle errors better
					}
					cell.Type = ValueTypeNumber
					cell.Value = n
				case calculus.CellValueTypeBoolean:
					cell.Type = ValueTypeBoolean
					if c.Value().(bool) {
						cell.Value = "TRUE"
					} else {
						cell.Value = "FALSE"
					}
				default:
					return
					// TODO Handle unknown?
				}
				sheet.Cells.Cells = append(sheet.Cells.Cells, cell)
			})
		})
		workbook.Sheets.Sheets = append(workbook.Sheets.Sheets, sheet)
	}
	return workbook, nil
}

func stringifyNumber(value interface{}) (string, error) {
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
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	}
	return "", calculus.NewError(nil, "Invalid numeric value")
}

func WriteTo(workbook *Workbook, out io.Writer) error {
	writer := gzip.NewWriter(out)
	if _, err := writer.Write([]byte(xml.Header)); err != nil {
		return calculus.NewError(err, "Failed write header")
	}
	xml, err := xml.Marshal(workbook)
	if err != nil {
		return calculus.NewError(err, "Failed to encode XML data")
	}
	if _, err := writer.Write(xml); err != nil {
		return calculus.NewError(err, "Failed to write or compress data")
	}
	err = writer.Flush()
	if err != nil {
		return calculus.NewError(err, "Failed to write file")
	}
	err = writer.Close()
	if err != nil {
		return calculus.NewError(err, "Failed to write file")
	}
	return nil
}
