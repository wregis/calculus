package csv

import (
	"bufio"
	"io"
	"strings"

	"github.com/wregis/calculus"
)

// Read receives a reader and create a workbook from it with default CSV configuration.
func Read(in io.Reader) (calculus.Workbook, error) {
	file := New()
	return file.Read(in)
}

func (f File) Read(in io.Reader) (calculus.Workbook, error) {
	workbook := calculus.New()
	sheet := workbook.ActiveSheet()

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
			sheet.SetValue(lineNo, index, value)
		}

		lineNo++
	}

	if err != io.EOF {
		return nil, calculus.NewError(err, "Failed to read CSV")
	}

	return workbook, nil
}
