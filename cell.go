package calculus

import (
	"fmt"
	"time"
)

// Default known cell value types.
const (
	CellValueTypeEmpty uint8 = iota
	CellValueTypeBoolean
	CellValueTypeInteger
	CellValueTypeFloat
	CellValueTypeError
	CellValueTypeString
	CellValueTypeDate
	CellValueTypeDuration
)

// Cell value error texts
const (
	ErrorInvalidValue     string = "#VALUE!"
	ErrorInvalidFormula   string = "#NAME?"
	ErrorDivisionByZero   string = "#DIV/0!"
	ErrorInvalidReference string = "#REF!"
	ErrorNoIntersection   string = "#NULL!"
	ErrorNotFound         string = "#N/A"
	ErrorInvalidNumber    string = "#NUM!"
)

// NewCell creates a cell with given value, doing auto detection of its type.
func NewCell(value interface{}) Cell {
	cell := cell{}
	cell.SetValue(value)
	return &cell
}

// Cell stores a value and some metadata.
type Cell interface {
	// Type returns the detected or defined type for the stored value.
	Type() uint8
	// Value returns the stored value itself.
	Value() interface{}
	SetValue(interface{})
	// Comment
	Comment() string
	// SetComment
	SetComment(string)
	// Style
	Style() *Style
	// Style
	SetStyle(*Style)
}

type cell struct {
	valueType uint8
	value     interface{}
	comment   string
	style     *Style
}

func (c cell) Type() uint8 {
	return c.valueType
}

func (c cell) Value() interface{} {
	return c.value
}

func (c *cell) SetValue(value interface{}) {
	switch value.(type) {
	case nil:
		c.valueType = CellValueTypeEmpty
		c.value = value
	case bool:
		c.valueType = CellValueTypeBoolean
		c.value = value
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		c.valueType = CellValueTypeInteger
		c.value = value
	case float32, float64:
		c.valueType = CellValueTypeFloat
		c.value = value
	case string, []rune:
		c.valueType = CellValueTypeString
		c.value = value
	case time.Time:
		c.valueType = CellValueTypeDate
		c.value = value
	case time.Duration:
		c.valueType = CellValueTypeDuration
		c.value = value
	default:
		c.valueType = CellValueTypeString
		c.value = fmt.Sprint(value)
	}
}

func (c cell) Comment() string {
	return c.comment
}

func (c *cell) SetComment(comment string) {
	c.comment = comment
}

func (c cell) Style() *Style {
	return c.style
}

func (c *cell) SetStyle(style *Style) {
	c.style = style
}
