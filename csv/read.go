package csv

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/wregis/calculus/internal/duration"

	"github.com/wregis/calculus"
	"github.com/wregis/calculus/internal/errors"
)

// ReadFile opens a file for reading as CSV and create a workbook from it with default configuration.
func ReadFile(path string, hints ...Hint) (calculus.Workbook, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New(err, "Failed to open file")
	}
	return Read(file, hints...)
}

// Read receives a reader and create a workbook from it with default CSV configuration.
func Read(in io.Reader, hints ...Hint) (calculus.Workbook, error) {
	file := New()
	return file.Read(in, hints...)
}

// Read reads a CSV file with given configuration.
func (f File) Read(in io.Reader, hints ...Hint) (calculus.Workbook, error) {
	workbook := calculus.New()
	sheet, _ := workbook.AddSheet("Sheet1")

	reader := bufio.NewReader(in)
	var lineNo int
	var line []byte
	var err error
	for err == nil {
		line, err = reader.ReadSlice('\n')

		if f.Comment != "" && strings.HasPrefix(string(line), f.Comment) {
			continue
		}

		lineStr := strings.TrimRight(string(line), "\r\n")
		if len(lineStr) == 0 {
			lineNo++
			continue
		}

		parts := strings.Split(lineStr, f.Delimiter)
		for index, value := range parts {
			if len(value) >= 2 {
				if strings.Index(value, f.Enclosure) == 0 && strings.LastIndex(value, f.Enclosure) == len(value)-1 {
					value = value[1 : len(value)-1]
				}
			}
			if len(value) == 0 {
				continue
			}
			if index < len(hints) {
				switch hint := hints[index]; hint.Type {
				case calculus.CellValueTypeBoolean:
					b, err := strconv.ParseBool(value)
					if err != nil {
						return nil, errors.Newf(err, `Failed to read CSV boolean value: "%s"`, value)
					}
					sheet.SetValue(lineNo, index, b)
				case calculus.CellValueTypeInteger:
					// TODO format
					i, err := strconv.ParseInt(value, 10, 64)
					if err != nil {
						return nil, errors.Newf(err, `Failed to read CSV integer value: "%s"`, value)
					}
					sheet.SetValue(lineNo, index, i)
				case calculus.CellValueTypeFloat:
					// TODO format
					f, err := strconv.ParseFloat(value, 'f')
					if err != nil {
						return nil, errors.Newf(err, `Failed to read CSV floating point value: "%s"`, value)
					}
					sheet.SetValue(lineNo, index, f)
				case calculus.CellValueTypeDate:
					// TODO format
					// TODO date, time and datetime
					format := "2006-01-02"
					if len(value) > 10 {
						format = "2006-01-02T15:04:05.999999"
					}
					t, err := time.Parse(format, value)
					if err != nil {
						return nil, errors.Newf(err, `Failed to parse value as date: "%s"`, value)
					}
					sheet.SetValue(lineNo, index, t)
				case calculus.CellValueTypeDuration:
					d, err := duration.Parse(value)
					if err != nil {
						return nil, errors.Newf(err, `Failed to read CSV duration value: "%s"`, value)
					}
					sheet.SetValue(lineNo, index, d)
				default:
					sheet.SetValue(lineNo, index, value)
				}
			} else {
				sheet.SetValue(lineNo, index, value)
			}
		}

		lineNo++
	}

	if err != io.EOF {
		return nil, errors.New(err, "Failed to read CSV")
	}

	return workbook, nil
}
