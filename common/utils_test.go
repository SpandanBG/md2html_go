package common

import (
	"fmt"
	"testing"
)

func isEqual[T comparable](t *testing.T, a, b T) {
	if a != b {
		t.Fatal("LHS: ", a, " != RHS: ", b)
	}
}

func isNil[T comparable](t *testing.T, a T) {
	var nilVaue T
	if a != nilVaue {
		t.Fatal(a, " is not nil ")
	}
}

func isNotNil[T comparable](t *testing.T, a T) {
	var nilVaue T
	if a == nilVaue {
		t.Fatal(a, " is not nil ")
	}
}

func TestEscapeRawToHTML(t *testing.T) {
	for _, test := range []struct {
		name   string
		input  rune
		output string
	}{
		{
			name:   "should turn & to &amp;",
			input:  '&',
			output: "&amp;",
		},
		{
			name:   "should turn \" to &quot;",
			input:  '"',
			output: "&quot;",
		},
		{
			name:   "should turn ' to &#39;",
			input:  '\'',
			output: "&#39;",
		},
		{
			name:   "should turn < to &lt;",
			input:  '<',
			output: "&lt;",
		},
		{
			name:   "should turn > to &gt;",
			input:  '>',
			output: "&gt;",
		},
		{
			name:   "should any other to itself",
			input:  'a',
			output: "a",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			actualOuput := EscapeRawToHTML(test.input)

			isEqual[string](t, test.output, string(actualOuput))
		})
	}
}

func TestSplitByLines(t *testing.T) {
	for _, test := range []struct {
		name   string
		input  string
		output []string
	}{
		{
			name:   "should return no lines for empty content",
			input:  "",
			output: []string{},
		},
		{
			name:   "should return one line for single line content ending with line feed",
			input:  fmt.Sprintf("hello, world%c", LineFeed),
			output: []string{"hello, world"},
		},
		{
			name:   "should return one line for single line content ending with carriage return",
			input:  fmt.Sprintf("hello, world%c", CarriageReturn),
			output: []string{"hello, world"},
		},
		{
			name:   "should return one line for single line content ending with carriage return and line feed",
			input:  fmt.Sprintf("hello, world%c%c", CarriageReturn, LineFeed),
			output: []string{"hello, world"},
		},
		{
			name:   "should return one line for single line content",
			input:  "hello, world",
			output: []string{"hello, world"},
		},
		{
			name:   "should return 2 line for 2 line content seperated by line feed",
			input:  fmt.Sprintf("hello,%cworld", LineFeed),
			output: []string{"hello,", "world"},
		},
		{
			name:   "should return 2 line for 2 line content seperated by carriage return",
			input:  fmt.Sprintf("hello,%cworld", CarriageReturn),
			output: []string{"hello,", "world"},
		},
		{
			name:   "should return 2 line for 2 line content seperated by carriage return",
			input:  fmt.Sprintf("hello,%cworld", CarriageReturn),
			output: []string{"hello,", "world"},
		},
		{
			name:   "should return 2 line for 2 line content seperated by carriage return followed by line feed",
			input:  fmt.Sprintf("hello,%c%cworld", CarriageReturn, LineFeed),
			output: []string{"hello,", "world"},
		},
		{
			name:   "should return an empty line for as single line ending",
			input:  fmt.Sprintf("%c", LineFeed),
			output: []string{""},
		},
		{
			name:   "should return 2 empty line for 2 line endings",
			input:  fmt.Sprintf("%c%c", LineFeed, LineFeed),
			output: []string{"", ""},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			actualOutput := SplitByLines(test.input)

			isEqual[int](t, len(test.output), len(actualOutput))
			for i, eachActual := range actualOutput {
				isEqual[string](t, test.output[i], eachActual)
			}
		})
	}
}

func TestIsBlankLine(t *testing.T) {
	for _, test := range []struct {
		name   string
		input  string
		output bool
	}{
		{
			name:   "should return true for empty line as blank line",
			input:  "",
			output: true,
		},
		{
			name:   "should return true for line with just spaces as blank line",
			input:  fmt.Sprintf("%c%c", Space, Space),
			output: true,
		},
		{
			name:   "should return true for line with just tabs as blank line",
			input:  fmt.Sprintf("%c%c", Tab, Tab),
			output: true,
		},
		{
			name:   "should return true for line with space and tab as blank line",
			input:  fmt.Sprintf("%c%c", Space, Tab),
			output: true,
		},
		{
			name:   "should return false for line with some character as blank line",
			input:  "asdf",
			output: false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			actualOutput := IsBlankLine(test.input)

			isEqual[bool](t, test.output, actualOutput)
		})
	}
}

func TestIsCharWhitespace(t *testing.T) {
	for _, test := range []struct {
		name   string
		input  rune
		output bool
	}{
		{
			name:   "should return false for non whitespace character",
			input:  'a',
			output: false,
		},
		{
			name:   "should return true for LineFeed",
			input:  LineFeed,
			output: true,
		},
		{
			name:   "should return true for FormFeed",
			input:  FormFeed,
			output: true,
		},
		{
			name:   "should return true for CarriageReturn",
			input:  CarriageReturn,
			output: true,
		},
		{
			name:   "should return true for Space",
			input:  Space,
			output: true,
		},
		{
			name:   "should return true for NoBreakSpace",
			input:  NoBreakSpace,
			output: true,
		},
		{
			name:   "should return true for OGHamSpaceMark",
			input:  OGHamSpaceMark,
			output: true,
		},
		{
			name:   "should return true for ENQuad",
			input:  ENQuad,
			output: true,
		},
		{
			name:   "should return true for EMQuad",
			input:  EMQuad,
			output: true,
		},
		{
			name:   "should return true for ENSpace",
			input:  ENSpace,
			output: true,
		},
		{
			name:   "should return true for EMSpace",
			input:  EMSpace,
			output: true,
		},
		{
			name:   "should return true for ThreePerEMSpace",
			input:  ThreePerEMSpace,
			output: true,
		},
		{
			name:   "should return true for FourPerEMSpace",
			input:  FourPerEMSpace,
			output: true,
		},
		{
			name:   "should return true for SixPerEMSpace",
			input:  SixPerEMSpace,
			output: true,
		},
		{
			name:   "should return true for FigureSpace",
			input:  FigureSpace,
			output: true,
		},
		{
			name:   "should return true for PunctuationSpace",
			input:  PunctuationSpace,
			output: true,
		},
		{
			name:   "should return true for ThinSpace",
			input:  ThinSpace,
			output: true,
		},
		{
			name:   "should return true for HairSpace",
			input:  HairSpace,
			output: true,
		},
		{
			name:   "should return true for NarrowNoBreakSpace",
			input:  NarrowNoBreakSpace,
			output: true,
		},
		{
			name:   "should return true for MediumMathematicalSpace",
			input:  MediumMathematicalSpace,
			output: true,
		},
		{
			name:   "should return true for IdeographicSpace",
			input:  IdeographicSpace,
			output: true,
		},
		{
			name:   "should return false for U+1999",
			input:  '\u1999',
			output: false,
		},
		{
			name:   "should return false for U+200B",
			input:  '\u200B',
			output: false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			isEqual[bool](t, test.output, IsCharWhiteSpace(test.input))
		})
	}
}

func TestSecureNullChar(t *testing.T) {
	for _, test := range []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "should return same string when no null char is present",
			input:  "asdf asdf",
			output: "asdf asdf",
		},
		{
			name:   "should repalce null with replacement char",
			input:  string([]rune{'a', Null}),
			output: string([]rune{'a', ReplacementCha}),
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			actualOuput := SecureNullChar(test.input)

			isEqual[string](t, test.output, actualOuput)
		})
	}
}
