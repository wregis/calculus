package csv

type File struct {
	Delimiter  string
	Enclosure  string
	EscapeChar rune
	Comment    string
}

func New() *File {
	return &File{
		Delimiter:  `,`,
		Enclosure:  `"`,
		EscapeChar: '\\',
		Comment:    "",
	}
}
