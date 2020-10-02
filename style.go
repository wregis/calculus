package calculus

import "fmt"

// Style is a set of view definitions that can be applied to a cell
type Style struct {
	Font         *Font
	NumberFormat string
}

// Font represents general font configuration
//
// Bold, Intalic, Strikethrough and Underline are flags. Size is what the name says.
// Name is the font name, like Arial or Helvetica.
//
// The color property can be a CSS color name, like *red* or *aliceblue*, a 6-digit
// hexadecimal RGB, like #ffd700 for *gold*, or a 8-digit hexadecimal RGBA (RGB with
// alpha/transparency), like #fff55080 for *coral* with 50% alpha.
type Font struct {
	Bold          bool
	Color         string
	Italic        bool
	Name          string
	Strikethrough bool
	Size          float32
	Underline     bool
}

// RGB updates the font color value by composing a RGB value
func (f *Font) RGB(r, g, b uint8) {
	f.Color = fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

// RGBA updates the font color value by composing a RGBA value
func (f *Font) RGBA(r, g, b, a uint8) {
	f.Color = fmt.Sprintf("#%02x%02x%02x%02x", r, g, b, a)
}
