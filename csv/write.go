package csv

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/wregis/calculus"
	"github.com/wregis/calculus/internal/duration"
	"github.com/wregis/calculus/internal/errors"
)

// Write generates a CSV file from the current active worksheet on the workbook with default CSV configutation.
func Write(workbook calculus.Workbook, out io.Writer) error {
	file := New()
	return file.Write(workbook, out)
}

// Write generates a CSV file from the current active worksheet on the workbook.
//
// Most of types will be printed in the most basic forms and special formatation will be ignored.
func (f File) Write(workbook calculus.Workbook, out io.Writer) (err error) {
	sheet := workbook.ActiveSheet()
	if sheet == nil {
		return errors.New(nil, "No sheet to write CSV")
	}

	prevRow := 0
	writer := bufio.NewWriter(out)
	sheet.Rows().OrderedWalk(func(rIndex int, row calculus.Row) {
		if prevRow != 0 {
			writer.WriteRune('\n')
		}
		if rIndex > prevRow {
			for ; prevRow < rIndex; prevRow++ {
				writer.WriteRune('\n')
			}
		}
		prevRow++

		prevCol := 0
		row.OrderedWalk(func(cIndex int, cell calculus.Cell) {
			if prevCol != 0 {
				writer.WriteString(f.Delimiter)
			}
			if cIndex > prevCol {
				for ; prevCol < cIndex; prevCol++ {
					writer.WriteString(f.Delimiter)
				}
			}
			prevCol++

			var value = f.stringValue(cell.Value())
			var enclosure bool
			if strings.Contains(value, f.Enclosure) {
				value = strings.ReplaceAll(value, f.Enclosure, string(f.EscapeChar)+f.Enclosure)
				enclosure = true
			}
			if strings.Contains(value, f.Delimiter) {
				value = strings.ReplaceAll(value, f.Delimiter, string(f.EscapeChar)+f.Delimiter)
				enclosure = true
			}
			if strings.ContainsRune(value, f.EscapeChar) {
				strings.ReplaceAll(value, string(f.EscapeChar), string(f.EscapeChar)+string(f.EscapeChar))
				enclosure = true
			}
			if enclosure || strings.IndexFunc(value, unicode.IsSpace) >= 0 {
				value = f.Enclosure + value + f.Enclosure
			}
			if _, werr := writer.WriteString(value); werr != nil {
				err = errors.New(werr, "Failed to write item")
			}
		})
	})
	writer.Flush()

	return err
}

func (f File) stringValue(value interface{}) string {
	if value != nil {
		switch v := value.(type) {
		case bool:
			return strconv.FormatBool(v)
		case float32:
			return strconv.FormatFloat(float64(v), 'f', -1, 32)
		case float64:
			return strconv.FormatFloat(v, 'f', -1, 64)
		case int:
			return strconv.FormatInt(int64(v), 10)
		case int8:
			return strconv.FormatInt(int64(v), 10)
		case int16:
			return strconv.FormatInt(int64(v), 10)
		case int32:
			return strconv.FormatInt(int64(v), 10)
		case int64:
			return strconv.FormatInt(v, 10)
		case uint:
			return strconv.FormatUint(uint64(v), 10)
		case uint8:
			return strconv.FormatUint(uint64(v), 10)
		case uint16:
			return strconv.FormatUint(uint64(v), 10)
		case uint32:
			return strconv.FormatUint(uint64(v), 10)
		case uint64:
			return strconv.FormatUint(v, 10)
		case string:
			return v
		case time.Time:
			return v.Format("2006-01-02T15:04:05.999999")
		case time.Duration:
			return duration.Format(v)
		default:
			return fmt.Sprint(v)
		}
	}
	return ""
}
