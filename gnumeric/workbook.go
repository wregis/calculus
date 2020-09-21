package gnumeric

import "encoding/xml"

// Boolean
type Boolean uint8

// Orientation is a printing orientation
type Orientation string

const (
	OrientationLandscape Orientation = "landscape"
	OrientationPortrait  Orientation = "portrait"
)

type Visibility string

const (
	VisibilityVisible Visibility = "GNM_SHEET_VISIBILITY_VISIBLE"
)

type ValueType int

const (
	ValueTypeBoolean ValueType = 20
	ValueTypeNumber  ValueType = 40
	ValueTypeString  ValueType = 60
)

// Workbook is a document root element
type Workbook struct {
	XMLName        xml.Name       `xml:"http://www.gnumeric.org/v10.dtd Workbook"`
	Version        *Version       `xml:"Version,omitempty"`
	Attributes     Attributes     `xml:"Attributes"`
	Summary        Summary        `xml:"Summary"`
	SheetNameIndex SheetNameIndex `xml:"SheetNameIndex"`
	Names          string         `xml:"Names"`
	Geometry       Geometry       `xml:"Geometry"`
	Sheets         Sheets         `xml:"Sheets"`
	UIData         UIData         `xml:"UIData"`
}

type Version struct {
	Epoch int    `xml:",attr"`
	Major int    `xml:",attr"`
	Minor int    `xml:",attr"`
	Full  string `xml:",attr"`
}

// Attributes is a document configuration root element, direct child of Workbook.
type Attributes struct {
	Attributes []Attribute `xml:"Attribute"`
}

// Attribute is a single document attribute.
type Attribute struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

// Summary is a document metadata root element, direct child of Workbook.
type Summary struct {
	Items []SummaryItems `xml:"Item"`
}

// SummaryItems is a single document metadata.
type SummaryItems struct {
	Name  string `xml:"name"`
	Value string `xml:"val-string"`
}

// SheetNameIndex
type SheetNameIndex struct {
	Name []string `xml:"SheetName"`
}

// Geometry is the document viewport information, direct child of Workbook.
type Geometry struct {
	Width  int `xml:",attr"`
	Height int `xml:",attr"`
}

// Sheets is the document sheets list root element, direct child of Workbook.
type Sheets struct {
	Sheets []Sheet `xml:"Sheet"`
}

// Sheet is a single worksheet with its own configuration.
type Sheet struct {
	DisplayFormulas     Boolean    `xml:",attr"`
	HideZero            Boolean    `xml:",attr"`
	HideGrid            Boolean    `xml:",attr"`
	HideColHeader       Boolean    `xml:",attr"`
	HideRowHeader       Boolean    `xml:",attr"`
	DisplayOutlines     Boolean    `xml:",attr"` // default true
	OutlineSymbolsBelow Boolean    `xml:",attr"` // default true
	OutlineSymbolsRight Boolean    `xml:",attr"` // default true
	Visibility          Visibility `xml:",attr"`
	GridColor           string     `xml:",attr"`
	Name                string
	MaxColumn           int `xml:"MaxCol"`
	MaxRow              int
	Zoom                float32
	Names               Names            `xml:"Names"`
	PrintInformation    PrintInformation `xml:"PrintInformation"`
	Styles              Styles           `xml:"Styles"`
	Columns             Columns          `xml:"Cols"`
	Rows                Rows
	Selections          Selections `xml:"Selections"`
	Cells               Cells
	SheetLayout         SheetLayout `xml:"SheetLayout"`
	// Solver
}

type Names struct {
	Names []Name `xml:"Name"`
}

type Name struct {
	Name     string `xml:"name"`
	Value    string `xml:"value"`
	Position string `xml:"position"`
}

// PrintInformation
type PrintInformation struct {
	Margins     Margins     `xml:"Margins"`
	Orientation Orientation `xml:"orientation"`
	Paper       string      `xml:"paper"`
}

// Margins
type Margins struct {
	Top    Margin `xml:"top"`
	Bottom Margin `xml:"bottom"`
	Left   Margin `xml:"left"`
	Right  Margin `xml:"right"`
	Header Margin `xml:"header"`
	Footer Margin `xml:"footer"`
}

// Margin
type Margin struct {
	Points        float32 `xml:",attr"`
	PreferredUnit string  `xml:"PrefUnit,attr"`
}

// Styles
type Styles struct {
	StyleRegions []StyleRegion `xml:"StyleRegion"`
}

// StyleRegion
type StyleRegion struct {
	StartColumn int   `xml:"startCol,attr"`
	StartRow    int   `xml:"startRow,attr"`
	EndColumn   int   `xml:"endCol,attr"`
	EndRow      int   `xml:"endRow,attr"`
	Style       Style `xml:"Style"`
}

// Style
type Style struct {
	HAlign       string  `xml:",attr"`
	VAlign       string  `xml:",attr"`
	WrapText     Boolean `xml:",attr"`
	ShrinkToFit  Boolean `xml:",attr"`
	Rotation     Boolean `xml:",attr"`
	Orient       Boolean `xml:",attr"`
	Shade        Boolean `xml:",attr"`
	Indent       Boolean `xml:",attr"`
	Locked       Boolean `xml:",attr"`
	Hidden       Boolean `xml:",attr"`
	Fore         string  `xml:",attr"`
	Back         string  `xml:",attr"`
	PatternColor string  `xml:",attr"`
	Format       string  `xml:",attr"`
	Font         *Font   `xml:"Font,omitempty"`
	// StyleBorder  *StyleBorder `xml:"StyleBorder,omitempty"`
}

// Font
type Font struct {
	Unit          float32 `xml:",attr"`
	Bold          Boolean `xml:",attr"`
	Italic        Boolean `xml:",attr"`
	Underline     Boolean `xml:",attr"`
	StrikeThrough Boolean `xml:",attr"`
	Script        string  `xml:",attr"` // ???
	Name          string  `xml:",chardata"`
}

// StyleBorder
type StyleBorder struct {
	Top         Border
	Botttom     Border
	Left        Border
	Right       Border
	Diagonal    Border
	RevDiagonal Border `xml:"Rev-Diagonal"`
}

// Border
type Border struct {
	Style Boolean `xml:",attr"`
}

// Columns is the metadata root element related to a sheet column, direct child of Workbook.
type Columns struct {
	DefaultSizePoints float32      `xml:"DefaultSizePts,attr"`
	ColInfo           []ColumnInfo `xml:"ColInfo"`
}

// ColumnInfo
type ColumnInfo struct {
	Number      int     `xml:"No,attr"`
	Width       float32 `xml:"Unit,attr"`
	LeftMargin  int     `xml:"MarginA,attr,omitempty"`
	RightMargin int     `xml:"MarginB,attr,omitempty"`
	Hidden      Boolean `xml:",attr,omitempty"`
	Count       int     `xml:",attr,omitempty"`
}

// Rows is the metadata root element related to a sheet row, direct child of Workbook.
type Rows struct {
	DefaultSizePts float32   `xml:",attr"`
	RowInfo        []RowInfo `xml:"RowInfo"`
}

// RowInfo
type RowInfo struct {
	Number       int     `xml:"No,attr"`
	Height       float32 `xml:"Unit,attr"`
	TopMargin    int     `xml:"MarginA,attr,omitempty"`
	BottomMargin int     `xml:"MarginB,attr,omitempty"`
	Hidden       Boolean `xml:",attr,omitempty"`
	Count        int     `xml:",attr,omitempty"`
}

// Selections
type Selections struct {
	CursorColumn int         `xml:"CursorCol,attr"`
	CursorRow    int         `xml:",attr"`
	Selections   []Selection `xml:"Selection"`
}

// Selection
type Selection struct {
	StartColumn int `xml:"startCol,attr"`
	StartRow    int `xml:"startRow,attr"`
	EndColumn   int `xml:"endCol,attr"`
	EndRow      int `xml:"endRow,attr"`
}

// Cells is the sheet group of values.
type Cells struct {
	Cells []Cell `xml:"Cell"`
}

// Cell is a sheet individual value.
type Cell struct {
	Column int       `xml:"Col,attr"`
	Row    int       `xml:",attr"`
	Style  int       `xml:",attr,omitempty"`
	Type   ValueType `xml:"ValueType,attr,omitempty"`
	Value  string    `xml:",chardata"`
}

// SheetLayout
type SheetLayout struct {
	TopLeft string `xml:",attr"`
}

// UIData
type UIData struct {
	SelectedTab int `xml:",attr"`
}
