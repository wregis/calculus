package calculus

import (
	"github.com/wregis/calculus/internal/errors"
)

// New creates a workbook with default options and a single sheet.
func New() Workbook {
	return &workbook{
		properties:           &workbookProperties{},
		activeSheet:          -1,
		sheets:               []Sheet{},
		showHorizontalScroll: true,
		showVerticalScroll:   true,
		showSheetTabs:        true,
	}
}

// Workbook is a collection of indexed data collections and its metadata.
type Workbook interface {
	Properties() WorkbookProperties
	ShowHorizontalScroll() bool
	SetShowHorizontalScroll(bool)
	ShowVerticalScroll() bool
	SetShowVerticalScroll(bool)
	ShowSheetTabs() bool
	SetShowSheetTabs(bool)
	// Sheet return the sheet with given name or nil if not found.
	Sheet(string) Sheet
	// SetActive updates the active sheet index and returns it.
	SetActive(name string) Sheet
	// ActiveSheet retrieves the the sheet currently marked as active.
	ActiveSheet() Sheet
	// AddSheet creates a new sheet with given name and appends it to the workbook. If another sheet with the same name
	// already exists on the work book the operation will fail.
	AddSheet(string) (Sheet, error)
	// AddSheetFirst creates a new sheet with given name and put it to the start of the workbook. If another sheet with the
	// same name already exists on the work book the operation will fail.
	AddSheetFirst(string) (Sheet, error)
	// RemoveSheet finds a sheet with given name and removes it from the workbook.
	RemoveSheet(string) error
	// Sheets returns the sheet collection of the workbook.
	Sheets() []Sheet
}

type workbook struct {
	properties           WorkbookProperties
	activeSheet          int
	sheets               []Sheet
	showHorizontalScroll bool
	showVerticalScroll   bool
	showSheetTabs        bool
}

func (w *workbook) Properties() WorkbookProperties {
	return w.properties
}

func (w *workbook) ShowHorizontalScroll() bool {
	return w.showHorizontalScroll
}

func (w *workbook) SetShowHorizontalScroll(showHorizontalScroll bool) {
	w.showHorizontalScroll = showHorizontalScroll
}

func (w *workbook) ShowVerticalScroll() bool {
	return w.showVerticalScroll
}

func (w *workbook) SetShowVerticalScroll(showVerticalScroll bool) {
	w.showVerticalScroll = showVerticalScroll
}

func (w *workbook) ShowSheetTabs() bool {
	return w.showSheetTabs
}

func (w *workbook) SetShowSheetTabs(showSheetTabs bool) {
	w.showSheetTabs = showSheetTabs
}

func (w *workbook) Sheet(name string) Sheet {
	for index := range w.sheets {
		if w.sheets[index].Name() == name {
			return w.sheets[index]
		}
	}
	return nil
}

func (w *workbook) SetActive(name string) Sheet {
	for index := range w.sheets {
		if w.sheets[index].Name() == name {
			w.activeSheet = index
			return w.sheets[index]
		}
	}
	return nil
}

func (w *workbook) ActiveSheet() Sheet {
	if len(w.sheets) == 0 {
		return nil
	}
	return w.sheets[w.activeSheet]
}

func (w *workbook) AddSheet(name string) (Sheet, error) {
	if w.Sheet(name) != nil {
		return nil, errors.Newf(nil, "Workbook already contains a sheet named %s", name)
	}
	sheet := NewSheet(name)
	w.sheets = append(w.sheets, sheet)
	w.activeSheet = len(w.sheets) - 1
	return sheet, nil
}

func (w *workbook) AddSheetFirst(name string) (Sheet, error) {
	if w.Sheet(name) != nil {
		return nil, errors.Newf(nil, "Workbook already contains a sheet named %s", name)
	}
	sheet := NewSheet(name)
	w.sheets = append([]Sheet{sheet}, w.sheets...)
	w.activeSheet = 0
	return sheet, nil
}

func (w *workbook) RemoveSheet(name string) error {
	for index := range w.sheets {
		if w.sheets[index].Name() == name {
			w.sheets = append(w.sheets[:index], w.sheets[index+1:]...)
			if w.activeSheet >= len(w.sheets) {
				w.activeSheet--
			}
			return nil
		}
	}
	return errors.Newf(nil, "There is no sheet with name %s", name)
}

func (w *workbook) Sheets() []Sheet {
	return w.sheets
}
