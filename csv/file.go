package csv

// File is a CSV file configuration.
type File struct {
	// Delimiter tells how to break a line into fields.
	Delimiter string
	// Enclosure indicates how strings with spaces, Delimiter, Enclosure or EscapeChar should be enveloped.
	Enclosure string
	// EscapeChar is used for certain cases where Delimiter, Enclosure or EspaceChar itself are present on the string.
	EscapeChar rune
	// Comment indicates how lines to ignore are prefixed.
	Comment string
	// Hints makes it possible for the reader to convert read data from string more specific types.
	Hints []Hint
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
