package common

/*
  - A `character` is a Unicode code point.
  - A line containing no characters, or line containing only spaces (U+0020) or
    tabs (U+0009), is called a blank line
*/

const (
	Null           = '\u0000'
	LineFeed       = '\u000A'
	FormFeed       = '\u000C'
	CarriageReturn = '\u000D'

	// A tab is regarded as 4 space if conversion is needed
	// More context required for https://spec.commonmark.org/0.30/#example-5
	Tab = '\u0009'

	// Unicode "Zs" category "Seprarator,Space" characters
	Space                   = '\u0020'
	NoBreakSpace            = '\u00A0'
	OGHamSpaceMark          = '\u1680'
	ENQuad                  = '\u2000'
	EMQuad                  = '\u2001'
	ENSpace                 = '\u2002'
	EMSpace                 = '\u2003'
	ThreePerEMSpace         = '\u2004'
	FourPerEMSpace          = '\u2005'
	SixPerEMSpace           = '\u2006'
	FigureSpace             = '\u2007'
	PunctuationSpace        = '\u2008'
	ThinSpace               = '\u2009'
	HairSpace               = '\u200A'
	NarrowNoBreakSpace      = '\u202F'
	MediumMathematicalSpace = '\u205F'
	IdeographicSpace        = '\u3000'

	EscChar = '\\'
)
