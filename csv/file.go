package csv

// File is a CSV file configuration.
type File struct {
	Delimiter  string
	Enclosure  string
	EscapeChar rune
	Comment    string
}

// New creates a new CSV file default configuration.
func New() *File {
	return &File{
		Delimiter:  `,`,
		Enclosure:  `"`,
		EscapeChar: '\\',
		Comment:    "",
	}
}

// Hint is a metadata to allow CSV to work with types and formatting
type Hint struct {
	// Type must be one of the following calculus.CellValueType constants:
	// * CellValueTypeBoolean
	// * CellValueTypeInteger
	// * CellValueTypeFloat
	// * CellValueTypeDate
	// * CellValueTypeDuration
	// * Any other value will be read as string
	//
	// All types are parsed according to the default format or simply using strconv methods.
	Type uint8
}
