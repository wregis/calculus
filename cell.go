package calculus

// CellValueType is a possible type for a cell value.
type CellValueType string

// Default known cell value types.
const (
	CellValueTypeString  CellValueType = "string"
	CellValueTypeNumber  CellValueType = "number"
	CellValueTypeBoolean CellValueType = "boolean"
	CellValueTypeFormula CellValueType = "formula"
	CellValueTypeError   CellValueType = "error"
)

// NewCell creates a cell with given value, doing auto detection of its type. If the type is not compatible, nil
// is returned.
func NewCell(value interface{}) Cell {
	cell := cell{value: value}
	switch value := value.(type) {
	case bool:
		cell.valueType = CellValueTypeBoolean
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32, float64:
		cell.valueType = CellValueTypeNumber
	case string:
		if len(value) > 0 && value[0] == '=' {
			cell.valueType = CellValueTypeFormula
		} else {
			cell.valueType = CellValueTypeString
		}
	default:
		return nil
	}
	return &cell
}

// Cell stores a value and some metadata.
type Cell interface {
	// Type returns the detected or defined type for the stored value.
	Type() CellValueType
	// Value returns the stored value itself.
	Value() interface{}
}

type cell struct {
	valueType CellValueType
	value     interface{}
}

func (c cell) Type() CellValueType {
	return c.valueType
}

func (c cell) Value() interface{} {
	return c.value
}
