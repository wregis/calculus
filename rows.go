package calculus

import "sort"

// Rows is a row data and properties handler.
type Rows interface {
	// Cell returns a cell of the row.
	Cell(int, int) Cell
	// SetCell stores a cell on the row.
	SetCell(int, int, Cell)
	// Iterate calls a callback to each row stored.
	Iterate(func(int, Row))
	// StableIterate calls a callback to each row stored on a stable order.
	StableIterate(func(int, Row))
}

type rows struct {
	rows map[int]Row
}

func (r *rows) Cell(row, column int) Cell {
	if row, ok := r.rows[row]; ok {
		if cell := row.Cell(column); cell != nil {
			return cell
		}
	}
	return nil
}

func (r *rows) SetCell(row, column int, cell Cell) {
	if _, ok := r.rows[row]; !ok {
		r.rows[row] = NewRow()
	}
	r.rows[row].SetCell(column, cell)
}

func (r *rows) Iterate(cb func(int, Row)) {
	if r.rows == nil {
		return
	}
	for key, row := range r.rows {
		cb(key, row)
	}
}

func (r *rows) StableIterate(cb func(int, Row)) {
	if r.rows == nil {
		return
	}
	keys := make([]int, 0, len(r.rows))
	for key := range r.rows {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		cb(key, r.rows[key])
	}
}

func NewRow() Row {
	return &row{
		cells: make(map[int]Cell),
	}
}

// A row is a collection of data cells.
type Row interface {
	// Cell returns a cell of the row.
	Cell(int) Cell
	// SetCell stores a cell on the row.
	SetCell(int, Cell)
	// Iterate calls a callback to each cell stored.
	Iterate(cb func(int, Cell))
	// StableIterate calls a callback to each row stored on a stable order.
	StableIterate(func(int, Cell))
	// Hidden tells if the row is hidden.
	Hidden() bool
	// SetHidden forces a row to be hidden or not.
	SetHidden(bool)
	// Height returns the row height.
	Height() float32
	// SetHeight updates the row heught.
	SetHeight(float32)
}

type row struct {
	cells  map[int]Cell
	hidden bool
	height float32
}

func (r *row) Cell(column int) Cell {
	if cell, ok := r.cells[column]; ok {
		return cell
	}
	return nil
}

func (r *row) SetCell(column int, cell Cell) {
	r.cells[column] = cell
}

func (r *row) Iterate(cb func(int, Cell)) {
	if r.cells == nil {
		return
	}
	for key, cell := range r.cells {
		cb(key, cell)
	}
}

func (r *row) StableIterate(cb func(int, Cell)) {
	if r.cells == nil {
		return
	}
	keys := make([]int, 0)
	for key := range r.cells {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		cb(key, r.cells[key])
	}
}

func (r *row) Hidden() bool {
	return r.hidden
}

func (r *row) SetHidden(hidden bool) {
	r.hidden = hidden
}

func (r *row) Height() float32 {
	return r.height
}

func (r *row) SetHeight(height float32) {
	r.height = height
}
