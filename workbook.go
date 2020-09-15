package calculus

// New creates a workbook with default options and a single sheet.
func New() Workbook {
	return &workbook{
		activeSheet: 0,
		sheets: []Sheet{
			NewSheet(),
		},
		showHorizontalScroll: true,
		showVerticalScroll:   true,
		showSheetTabs:        true,
	}
}

// Workbook is a collection of indexed data collections and its metadata.
type Workbook interface {
	// Sheet return the sheet with given name or nil if not found.
	Sheet(string) Sheet
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
	activeSheet          int
	sheets               []Sheet
	showHorizontalScroll bool
	showVerticalScroll   bool
	showSheetTabs        bool
}

func (w *workbook) Sheet(name string) Sheet {
	for index := range w.sheets {
		if w.sheets[index].Name() == name {
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
		return nil, NewErrorf(nil, "Workbook already contains a sheet named %s", name)
	}
	sheet := NewSheet()
	w.sheets = append(w.sheets, sheet)
	w.activeSheet = len(w.sheets) - 1
	return sheet, nil
}

func (w *workbook) AddSheetFirst(name string) (Sheet, error) {
	if w.Sheet(name) != nil {
		return nil, NewErrorf(nil, "Workbook already contains a sheet named %s", name)
	}
	sheet := NewSheet()
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
		}
	}
	return NewErrorf(nil, "There is no sheet with name %s", name)
}

func (w *workbook) Sheets() []Sheet {
	return w.sheets
}
