package xlsx

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"

	"github.com/wregis/calculus"
)

// ReadFile takes a filename of a valid XLSX file as input and generates a workbook object from it.
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
	if err := readWorksheets(reader, wb); err != nil {
		return nil, err
	}
	// TODO themes, styles...
	if err := readWorkbook(reader, wb); err != nil {
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

var sheetFilenameExp = regexp.MustCompile(`^xl/worksheets/sheet[0-9]+.xml$`)

func readWorksheets(reader *zip.Reader, target calculus.Workbook) error {
	sttData, err := getZipFile(reader, "xl/sharedStrings.xml")
	if err != nil {
		return err
	}
	var stt sharedStringTable
	if sttData != nil {
		if err := xml.Unmarshal(sttData, &stt); err != nil {
			return calculus.NewError(err, "Failed to parse shared strings table file")
		}
	}

	for _, file := range reader.File {
		if sheetFilenameExp.MatchString(file.Name) {
			data, err := readZipFile(file)
			if err != nil {
				return err
			}
			var worksheet worksheet
			xml.Unmarshal(data, &worksheet)

			sheet, err := target.AddSheet(file.Name)
			if err != nil {
				return calculus.NewError(err, "Failed to create worksheet")
			}
			for _, row := range worksheet.Data.Rows {
				for _, record := range row.Records {
					row, column, err := calculus.ParseCoordinate(record.CellReference)
					if err != nil {
						return calculus.NewError(err, "Failed to parse worksheet")
					}
					if record.CellValue == "" {
						continue // TODO empty cells might appear for styles
					}
					switch record.DataType {
					case dataTypeBoolean:
						sheet.SetValue(row, column, record.CellValue == "1")
					case dataTypeInlineString, dataTypeFormula:
						sheet.SetValue(row, column, record.CellValue)
					case dataTypeNumber, "":
						f, err := strconv.ParseFloat(record.CellValue, 64)
						if err != nil {
							return calculus.NewErrorf(err, "Failed to parse value as number: \"%+v\"", record.CellValue)
						}
						sheet.SetValue(row, column, f)
					case dataTypeSharedString:
						index, err := strconv.Atoi(record.CellValue)
						if err != nil {
							return calculus.NewErrorf(err, "Failed to parse value as integer: %+v", record.CellValue)
						}
						sheet.SetValue(row, column, stt.Get(index))
					}
				}
			}
		}
	}

	return nil
}

func readWorkbook(reader *zip.Reader, target calculus.Workbook) error {
	wbData, err := getZipFile(reader, "xl/workbook.xml")
	if err != nil {
		return err
	} else if wbData == nil {
		return calculus.NewError(nil, "Missing workbook file")
	}
	var workbook workbook
	if err := xml.Unmarshal(wbData, &workbook); err != nil {
		return calculus.NewError(err, "Failed to parse workbook file")
	}

	relsData, err := getZipFile(reader, "xl/_rels/workbook.xml.rels")
	if err != nil {
		return err
	} else if relsData == nil {
		return calculus.NewError(nil, "Missing workbook relationships file")
	}
	var workbookRels relationships
	if err := xml.Unmarshal(relsData, &workbookRels); err != nil {
		return calculus.NewError(err, "Failed to parse workbook file")
	}

	target.Properties().SetApplication(workbook.FileVersion.AppName)
	target.Properties().SetDate1904(workbook.WorkbookProperties.Date1904)
	target.SetShowHorizontalScroll(workbook.BookViews.WorkbookView.ShowHorizontalScroll)
	target.SetShowVerticalScroll(workbook.BookViews.WorkbookView.ShowVerticalScroll)
	target.SetShowSheetTabs(workbook.BookViews.WorkbookView.ShowSheetTabs)
	for _, sheet := range workbook.Sheets.Sheets {
		for _, relationship := range workbookRels.Relationships {
			if sheet.ReferenceID == relationship.ID {
				worksheet := target.Sheet("xl/" + relationship.Target)
				if worksheet == nil {
					return calculus.NewError(nil, "Invalid workbook sheet reference")
				}
				worksheet.SetName(sheet.Name)
				worksheet.SetState(calculus.SheetState(sheet.Visibility))
				break
			}
		}
	}

	return nil
}
