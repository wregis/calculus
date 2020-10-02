package gnumeric

import "encoding/xml"

// Most structs here are based on the gnumeric spec at https://gitlab.gnome.org/GNOME/gnumeric/-/blob/master/gnumeric.xsd

// GNumeric boolean equivalent
const (
	BooleanFalse uint8 = 0
	BooleanTrue  uint8 = 1
)

// Sheet printing orientation
const (
	OrientationLandscape string = "landscape"
	OrientationPortrait  string = "portrait"
)

// Sheet visibility
const (
	VisibilityVisible    string = "GNM_SHEET_VISIBILITY_VISIBLE"
	VisibilityHidden     string = "GNM_SHEET_VISIBILITY_HIDDEN"
	VisibilityVeryHidden string = "GNM_SHEET_VISIBILITY_VERY_HIDDEN"
)

// Cell value type
const (
	ValueTypeEmpty     uint8 = 10
	ValueTypeBoolean   uint8 = 20
	ValueTypeInteger   uint8 = 30
	ValueTypeFloat     uint8 = 40
	ValueTypeError     uint8 = 50
	ValueTypeString    uint8 = 60
	ValueTypeCellRange uint8 = 70
	ValueTypeArray     uint8 = 80
)

// Cell content alignment
const (
	HorizontalAlignLeft   string = "GNM_HALIGN_LEFT"   // 1
	HorizontalAlignRight  string = "GNM_HALIGN_RIGHT"  // 2
	HorizontalAlignCenter string = "GNM_HALIGN_CENTER" // 3
	HorizontalAlignFill   string = "GNM_HALIGN_FILL"   // 5
	VerticalAlignTop      string = "GNM_VALIGN_TOP"    // T
	VerticalAlignCenter   string = "GNM_VALIGN_CENTER" // C
	VerticalAlignBottom   string = "GNM_VALIGN_BOTTOM" // B
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

// Version  is the document generator application version info
type Version struct {
	Epoch uint16 `xml:",attr"`
	Major uint16 `xml:",attr"`
	Minor uint16 `xml:",attr"`
	Full  string `xml:",attr"`
}

type Calculation struct {
	ManualRecalculation bool    `xml:"ManualRecalc,attr"`
	EnableIteration     bool    `xml:",attr"`
	MaxIterations       uint    `xml:",attr"`
	FloatRadix          uint8   `xml:",attr"`
	FloatDigits         uint8   `xml:",attr"`
	IterationTolerance  float64 `xml:",attr"`
	DateConvention      string  // "Apple:1904" or "ODF:1899"
}

// Attributes is a document configuration root element, direct child of Workbook.
type Attributes struct {
	Attributes []Attribute `xml:"Attribute"`
}

// Attribute is a single document attribute.
type Attribute struct {
	Type  uint8  `xml:"type,omitempty"` // must be 4
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
	Width  uint16 `xml:",attr,omitempty"`
	Height uint16 `xml:",attr,omitempty"`
}

// Sheets is the document sheets list root element, direct child of Workbook.
type Sheets struct {
	Sheets []Sheet `xml:"Sheet"`
}

// Sheet is a single worksheet with its own configuration.
type Sheet struct {
	DisplayFormulas     uint8  `xml:",attr"`
	HideZero            uint8  `xml:",attr"`
	HideGrid            uint8  `xml:",attr"`
	HideColHeader       uint8  `xml:",attr"`
	HideRowHeader       uint8  `xml:",attr"`
	DisplayOutlines     uint8  `xml:",attr"` // default true
	OutlineSymbolsBelow uint8  `xml:",attr"` // default true
	OutlineSymbolsRight uint8  `xml:",attr"` // default true
	Visibility          string `xml:",attr"`
	GridColor           string `xml:",attr"`
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
	Margins     Margins `xml:"Margins"`
	Orientation string  `xml:"orientation"`
	Paper       string  `xml:"paper"`
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
	HorizontalAlign string `xml:"HAlign,attr"`
	VerticalAlign   string `xml:"VAlign,attr"`
	WrapText        uint8  `xml:",attr"`
	ShrinkToFit     uint8  `xml:",attr"`
	Rotation        uint8  `xml:",attr"`
	Orient          uint8  `xml:",attr"`
	Shade           uint8  `xml:",attr"`
	Indent          uint8  `xml:",attr"`
	Locked          uint8  `xml:",attr"`
	Hidden          uint8  `xml:",attr"`
	Fore            string `xml:",attr"`
	Back            string `xml:",attr"`
	PatternColor    string `xml:",attr"`
	Format          string `xml:",attr"`
	Font            *Font  `xml:",omitempty"`
	// StyleBorder  *StyleBorder `xml:"StyleBorder,omitempty"`
}

// Font
type Font struct {
	Unit          float32 `xml:",attr"`
	Bold          uint8   `xml:",attr"`
	Italic        uint8   `xml:",attr"`
	Underline     uint8   `xml:",attr"`
	StrikeThrough uint8   `xml:",attr"`
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
	Style uint8 `xml:",attr"`
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
	Hidden      uint8   `xml:",attr,omitempty"`
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
	Hidden       uint8   `xml:",attr,omitempty"`
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
	Column int    `xml:"Col,attr"`
	Row    int    `xml:",attr"`
	Style  int    `xml:",attr,omitempty"`
	Type   uint8  `xml:"ValueType,attr,omitempty"`
	Format string `xml:"ValueFormat,attr,omitempty"`
	Value  string `xml:",chardata"`
}

// SheetLayout
type SheetLayout struct {
	TopLeft string `xml:",attr"`
}

// UIData
type UIData struct {
	SelectedTab uint16 `xml:",attr"`
}
