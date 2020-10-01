package ods

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/wregis/calculus"
	"github.com/wregis/calculus/internal/duration"
)

// ReadFile takes a filename of a valid OpenDocument Spreadsheet file as input and generates a workbook object from it.
func ReadFile(path string) (calculus.Workbook, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, calculus.NewError(err, "Failed to open file")
	}
	reader := bytes.NewReader(b)
	return Read(reader, int64(reader.Len()))
}

// Read generates a workbook object from the bytes read from the input.
func Read(in io.ReaderAt, size int64) (calculus.Workbook, error) {
	reader, err := zip.NewReader(in, size)
	if err != nil {
		return nil, calculus.NewError(err, "Failed to read file")
	}
	wb := calculus.New()
	if err := readContent(reader, wb); err != nil {
		return nil, err
	}
	return wb, nil
}

func getZipFile(reader *zip.Reader, filename string) ([]byte, error) {
	var file *zip.File
	for _, f := range reader.File {
		if f.Name == filename {
			file = f
			break
		}
	}
	if file == nil {
		return nil, nil
	}
	return readZipFile(file)
}

func readZipFile(file *zip.File) ([]byte, error) {
	rc, err := file.Open()
	if err != nil {
		return nil, calculus.NewErrorf(err, "Failed to read file: %s", file.Name)
	}
	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, calculus.NewErrorf(err, "Failed to read data: %s", file.Name)
	}
	err = rc.Close()
	if err != nil {
		return nil, calculus.NewErrorf(err, "Failed to close reader: %s", file.Name)
	}
	return data, nil
}

func readContent(reader *zip.Reader, target calculus.Workbook) error {
	contentData, err := getZipFile(reader, "content.xml")
	if err != nil {
		return err
	}
	var content documentContent
	if contentData != nil {
		if err := xml.Unmarshal(contentData, &content); err != nil {
			return calculus.NewError(err, "Failed to parse shared strings table file")
		}
	}
	for _, table := range content.Body.Spreadsheet.Tables {
		sheet, err := target.AddSheet(table.Name)
		if err != nil {
			return calculus.NewError(err, "Failed to create worksheet")
		}
		for rowIndex, row := range table.Rows {
			for cellIndex, cell := range row.Cells {
				log.Printf("%#v", cell)
				switch cell.ValueType {
				case "float", "currency", "percentage":
					f, err := strconv.ParseFloat(cell.Value, 64)
					if err != nil {
						return calculus.NewErrorf(err, "Failed to parse value as number: \"%+v\"", cell.Value)
					}
					sheet.SetValue(rowIndex, cellIndex, f)
				case "date":
					format := "2006-01-02"
					if len(cell.DateValue) > 10 {
						format = "2006-01-02T15:04:05.999999"
					}
					t, err := time.Parse(format, cell.DateValue)
					if err != nil {
						return calculus.NewErrorf(err, "Failed to parse value as date: \"%+v\"", cell.DateValue)
					}
					sheet.SetValue(rowIndex, cellIndex, t)
				case "time":
					d, err := duration.Parse(cell.TimeValue)
					if err != nil {
						return calculus.NewErrorf(err, "Failed to parse value as time: \"%+v\"", cell.TimeValue)
					}
					sheet.SetValue(rowIndex, cellIndex, d)
				case "string":
					sheet.SetValue(rowIndex, cellIndex, cell.Text)
				default:
					// TODO Style-only cell
				}
			}
		}
	}
	return nil
}
