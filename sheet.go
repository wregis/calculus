package calculus

// SheetState represents a visibility state for a sheet
type SheetState string

const (
	// SheetStateVisible represents a sheet visible
	SheetStateVisible SheetState = "visible"
	// SheetStateHidden represents a sheet hidden
	SheetStateHidden SheetState = "hidden"
	// SheetStateVeryHidden represents a sheet very hidden
	SheetStateVeryHidden SheetState = "veryHidden"
)

// NewSheet creates a new empty sheet.
func NewSheet(name string) Sheet {
	return &sheet{
		name:  name,
		state: SheetStateVisible,
		rows:  NewRows(),
	}
}

// Sheet is a collection of organized data.
type Sheet interface {
	// Name returns the sheet name label.
	Name() string
	// SetName updates the sheet name label.
	SetName(string)
	// Rows returns the row manager instance.
	Rows() Rows
	// Value returns the stored value itself.
	Value(int, int) interface{}
	// ValueByRef returns the stored value itself using a string coordinate.
	ValueByRef(string) interface{}
	// SetValue stores a value on a specific position.
	SetValue(int, int, interface{})
	// SetValue stores a value on a specific position using a string coordinate.
	SetValueByRef(string, interface{})
	State() SheetState
	SetState(SheetState)
}

type sheet struct {
	name  string
	state SheetState
	rows  Rows
}

func (s *sheet) Name() string {
	return s.name
}

func (s *sheet) SetName(name string) {
	s.name = name
}

func (s *sheet) Rows() Rows {
	return s.rows
}

func (s *sheet) Value(row, column int) interface{} {
	if cell := s.rows.Cell(row, column); cell != nil {
		return cell.Value()
	}
	return nil
}

func (s *sheet) ValueByRef(coordinate string) interface{} {
	row, column, err := ParseCoordinate(coordinate)
	if err != nil {
		return nil
	}
	return s.Value(row, column)
}

func (s *sheet) SetValue(row, column int, value interface{}) {
	if cell := NewCell(value); cell != nil {
		s.rows.SetCell(row, column, cell)
	}
}

func (s *sheet) SetValueByRef(coordinate string, value interface{}) {
	row, column, err := ParseCoordinate(coordinate)
	if err != nil {
		return
	}
	s.SetValue(row, column, value)
}

func (s *sheet) State() SheetState {
	return s.state
}

func (s *sheet) SetState(state SheetState) {
	s.state = state
}
