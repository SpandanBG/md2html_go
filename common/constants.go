package common

/*
  - A `character` is a Unicode code point.
  - A line containing no characters, or line containing only spaces (U+0020) or
    tabs (U+0009), is called a blank line
*/

const (
	Null           = '\u0000'
	LineFeed       = '\u000A'
	CarriageReturn = '\u000D'
	Tab            = '\u0009'

	// Unicode "Zs" category "Seprarator,Space" characters
	Space                   = '\u0020'
	NoBreakSpace            = "U+00A0"
	OGHamSpaceMark          = "U+1680"
	ENQuad                  = "U+2000"
	EMQuad                  = "U+2001"
	ENSpace                 = "U+2002"
	EMSpace                 = "U+2003"
	ThreePerEMSpace         = "U+2004"
	FourPerEMSpace          = "U+2005"
	SixPerEMSpace           = "U+2006"
	FigureSpace             = "U+2007"
	PunctuationSpace        = "U+2008"
	ThinSpace               = "U+2009"
	HairSpace               = "U+200A"
	NarrowNoBreakSpace      = "U+202F"
	MediumMathematicalSpace = "U+205F"
	IdeographicSpace        = "U+3000"

	EscChar = '\\'
)
