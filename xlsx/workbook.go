package xlsx

import (
	"archive/zip"
	"encoding/xml"
	"fmt"

	"github.com/wregis/calculus"
)

func writeWorkbook(writer *zip.Writer, wb calculus.Workbook) error {
	workbook := workbook{
		FileVersion: fileVersion{
			AppName: wb.Properties().Application(),
		},
		WorkbookProperties: workbookProperties{
			Date1904: wb.Properties().Date1904(),
		},
		BookViews: bookViews{
			WorkbookView: workbookView{
				ShowHorizontalScroll: wb.ShowHorizontalScroll(),
				ShowVerticalScroll:   wb.ShowVerticalScroll(),
				ShowSheetTabs:        wb.ShowSheetTabs(),
				WindowWidth:          16384,
				WindowHeight:         8192,
				TabRatio:             500,
			},
		},
		Sheets: sheets{
			Sheets: []sheet{},
		},
	}
	for index, worksheet := range wb.Sheets() {
		workbook.Sheets.Sheets = append(workbook.Sheets.Sheets, sheet{
			Visibility:  "visible",
			Name:        worksheet.Name(),
			ID:          index + 1,
			ReferenceID: fmt.Sprintf("rId%d", index+1+3),
		})
	}
	if err := writeXMLToFile(writer, "xl/workbook.xml", workbook); err != nil {
		return err
	}
	return writeWorkbookRelationships(writer, wb)
}

type workbook struct {
	XMLName            xml.Name           `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main workbook"`
	FileVersion        fileVersion        `xml:"fileVersion"`
	WorkbookProperties workbookProperties `xml:"workbookPr"`
	BookViews          bookViews          `xml:"bookViews"`
	Sheets             sheets             `xml:"sheets"`
}

type fileVersion struct {
	AppName string `xml:"appName,attr"`
}

type workbookProperties struct {
	Date1904 bool `xml:"date1904,attr,omitempty"`
}

type bookViews struct {
	WorkbookView workbookView `xml:"workbookView"`
}

type workbookView struct {
	ShowHorizontalScroll bool `xml:"showHorizontalScroll,attr,omitempty"`
	ShowVerticalScroll   bool `xml:"showVerticalScroll,attr,omitempty"`
	ShowSheetTabs        bool `xml:"showSheetTabs,attr,omitempty"`
	XWindow              int  `xml:"xWindow,attr,omitempty"`
	YWindow              int  `xml:"yWindow,attr,omitempty"`
	WindowWidth          int  `xml:"windowWidth,attr,omitempty"`
	WindowHeight         int  `xml:"windowHeight,attr,omitempty"`
	TabRatio             int  `xml:"tabRatio,attr,omitempty"`
	FirstSheet           int  `xml:"firstSheet,attr,omitempty"`
	ActiveTab            int  `xml:"activeTab,attr,omitempty"`
}

type sheets struct {
	Sheets []sheet `xml:"sheet"`
}

type sheet struct {
	Visibility  string `xml:"state,attr,omitempty"`
	Name        string `xml:"name,attr"`
	ID          int    `xml:"sheetId,attr"`
	ReferenceID string `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
}
