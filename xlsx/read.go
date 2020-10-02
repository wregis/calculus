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
	"github.com/wregis/calculus/internal/coordinate"
	"github.com/wregis/calculus/internal/errors"
)

// ReadFile takes a filename of a valid XLSX file as input and generates a workbook object from it.
func ReadFile(path string) (calculus.Workbook, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New(err, "Failed to open file")
	}
	reader := bytes.NewReader(b)
	return Read(reader, int64(reader.Len()))
}

var sheetFilenameExp = regexp.MustCompile(`^xl/worksheets/sheet[0-9]+.xml$`)

// Read generates a workbook object from the bytes read from the input.
func Read(in io.ReaderAt, size int64) (calculus.Workbook, error) {
	reader, err := zip.NewReader(in, size)
	if err != nil {
		return nil, errors.New(err, "Failed to read file")
	}

	workbook, wb, err := readWorkbook(reader)
	if err != nil {
		return nil, err
	}
	workbookRels, err := readWorkbookRels(reader)
	if err != nil {
		return nil, err
	}
	styles, err := readStyles(reader)
	if err != nil {
		return nil, err
	}
	stt, err := readSharedStringTable(reader)
	if err != nil {
		return nil, err
	}
	for _, file := range reader.File {
		if sheetFilenameExp.MatchString(file.Name) {
			data, err := readZipFile(file)
			if err != nil {
				return nil, err
			}
			var worksheet worksheet
			if err := xml.Unmarshal(data, &worksheet); err != nil {
				return nil, errors.New(err, "Unable to parse worksheet")
			}

			var name string
			for _, relationship := range workbookRels.Relationships {
				if file.Name == "xl/"+relationship.Target {
					for _, sheet := range workbook.Sheets.Sheets {
						if sheet.ReferenceID == relationship.ID {
							name = sheet.Name
						}
					}
					break
				}
			}
			sheet, err := wb.AddSheet(name)
			if err != nil {
				return nil, errors.New(err, "Failed to create worksheet")
			}
			if err := readWorksheet(&worksheet, stt, sheet, styles); err != nil {
				return nil, err
			}
		}
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
		return nil, errors.Newf(err, "Failed to read file: %s", file.Name)
	}
	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, errors.Newf(err, "Failed to read data: %s", file.Name)
	}
	err = rc.Close()
	if err != nil {
		return nil, errors.Newf(err, "Failed to close reader: %s", file.Name)
	}
	return data, nil
}

func readWorkbook(reader *zip.Reader) (*workbook, calculus.Workbook, error) {
	wbData, err := getZipFile(reader, "xl/workbook.xml")
	if err != nil {
		return nil, nil, err
	} else if wbData == nil {
		return nil, nil, errors.New(nil, "Missing workbook file")
	}

	var workbook workbook
	if err := xml.Unmarshal(wbData, &workbook); err != nil {
		return nil, nil, errors.New(err, "Failed to parse workbook file")
	}
	wb := calculus.New()
	wb.Properties().SetApplication(workbook.FileVersion.AppName)
	wb.Properties().SetDate1904(workbook.WorkbookProperties.Date1904)
	wb.SetShowHorizontalScroll(workbook.BookViews.WorkbookView.ShowHorizontalScroll)
	wb.SetShowVerticalScroll(workbook.BookViews.WorkbookView.ShowVerticalScroll)
	wb.SetShowSheetTabs(workbook.BookViews.WorkbookView.ShowSheetTabs)

	return &workbook, wb, nil
}

func readWorkbookRels(reader *zip.Reader) (*relationships, error) {
	relsData, err := getZipFile(reader, "xl/_rels/workbook.xml.rels")
	if err != nil {
		return nil, err
	} else if relsData == nil {
		return nil, errors.New(nil, "Missing workbook relationships file")
	}
	var workbookRels relationships
	if err := xml.Unmarshal(relsData, &workbookRels); err != nil {
		return nil, errors.New(err, "Failed to parse workbook file")
	}
	return &workbookRels, nil
}

func readStyles(reader *zip.Reader) (map[uint16]*calculus.Style, error) {
	stylesData, err := getZipFile(reader, "xl/styles.xml")
	if err != nil {
		return nil, err
	}
	var styles stylesheet
	if err := xml.Unmarshal(stylesData, &styles); err != nil {
		return nil, errors.New(err, "Failed to parse styles file")
	}

	if styles.CellXfs.Count <= 0 {
		return nil, nil
	}

	numFmts := map[uint16]string{}
	if styles.NumberingFormats.Count > 0 {
		for _, fmt := range styles.NumberingFormats.Formats {
			numFmts[fmt.ID] = fmt.FormatCode
		}
	}

	fonts := map[int]*calculus.Font{}
	if styles.Fonts.Count > 0 {
		for index, font := range styles.Fonts.Fonts {
			f := calculus.Font{
				Size:          font.Size.Value,
				Name:          font.Name.Value,
				Color:         font.Color.RGB, // TODO Indexed
				Bold:          bool(font.Bold),
				Italic:        bool(font.Italic),
				Strikethrough: bool(font.Strikethrough),
				Underline:     bool(font.Underline),
			}
			fonts[index] = &f
		}
	}
	cellStyles := map[uint16]*calculus.Style{}
	for index, stl := range styles.CellXfs.Xfs {
		style := calculus.Style{
			Font:         fonts[int(stl.FontID)],
			NumberFormat: numFmts[stl.NumFmtID],
		}
		cellStyles[uint16(index)] = &style
	}
	return cellStyles, nil
}

func readSharedStringTable(reader *zip.Reader) (*sharedStringTable, error) {
	sttData, err := getZipFile(reader, "xl/sharedStrings.xml")
	if err != nil {
		return nil, err
	}
	var stt sharedStringTable
	if err := xml.Unmarshal(sttData, &stt); err != nil {
		return nil, errors.New(err, "Failed to parse shared strings table file")
	}
	return &stt, nil
}

func readWorksheet(worksheet *worksheet, stt *sharedStringTable, sheet calculus.Sheet, styles map[uint16]*calculus.Style) error {
	for _, row := range worksheet.Data.Rows {
		for _, record := range row.Records {
			row, column, err := coordinate.Parse(record.CellReference)
			if err != nil {
				return errors.New(err, "Failed to parse worksheet")
			}
			var cell calculus.Cell
			if record.CellValue != "" {
				switch record.DataType {
				case dataTypeBoolean:
					cell = sheet.SetValue(row, column, record.CellValue == "1")
				case dataTypeInlineString, dataTypeFormula:
					cell = sheet.SetValue(row, column, record.CellValue)
				case dataTypeNumber, "":
					f, err := strconv.ParseFloat(record.CellValue, 64)
					if err != nil {
						return errors.Newf(err, "Failed to parse value as number: \"%+v\"", record.CellValue)
					}
					cell = sheet.SetValue(row, column, f)
				case dataTypeSharedString:
					index, err := strconv.Atoi(record.CellValue)
					if err != nil {
						return errors.Newf(err, "Failed to parse value as integer: %+v", record.CellValue)
					}
					cell = sheet.SetValue(row, column, stt.Get(index))
				}
			} else {
				cell = sheet.SetValue(row, column, nil)
			}

			if style, ok := styles[record.StyleIndex]; ok {
				cell.SetStyle(style)
			}
		}
	}
	return nil
}
