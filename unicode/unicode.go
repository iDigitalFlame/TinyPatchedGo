// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// 2022 - Shrinkage by iDigitalFlame

package unicode

const (
	// MaxRune represents the maximum valid Unicode code point.
	MaxRune = '\U0010FFFF'
	// MaxASCII represents the maximum ASCII value.
	MaxASCII = '\u007F'
	// ReplacementChar represents invalid code points.
	ReplacementChar = '\uFFFD'
)

// Range16 represents of a range of 16-bit Unicode code points. The range runs
// from Lo to Hi inclusive and has the specified stride.
type Range16 struct {
	Lo     uint16
	Hi     uint16
	Stride uint16
}

// Range32 represents of a range of Unicode code points and is used when one or
// more of the values will not fit in 16 bits. The range runs from Lo to Hi
// inclusive and has the specified stride. Lo and Hi must always be >= 1<<16.
type Range32 struct {
	Lo     uint32
	Hi     uint32
	Stride uint32
}

// RangeTable defines a set of Unicode code points by listing the ranges of
// code points within the set. The ranges are listed in two slices
// to save space: a slice of 16-bit ranges and a slice of 32-bit ranges.
// The two slices must be in sorted order and non-overlapping.
// Also, R32 should contain only values >= 0x10000 (1<<16).
type RangeTable struct {
	R16         []Range16
	R32         []Range32
	LatinOffset int
}

// SpecialCase represents language-specific case mappings such as Turkish.
// Methods of SpecialCase customize (by overriding) the standard mappings.
type SpecialCase struct{}

var (
	// Scripts is the set of Unicode script tables.
	Scripts = map[string]*RangeTable{}
	// Categories is the set of Unicode category tables.
	Categories = map[string]*RangeTable{}
	// FoldScript maps a script name to a table of
	// code points outside the script that are equivalent under
	// simple case folding to code points inside the script.
	// If there is no entry for a script name, there are no such points.
	FoldScript = map[string]*RangeTable{}
	// FoldCategory maps a category name to a table of
	// code points outside the category that are equivalent under
	// simple case folding to code points inside the category.
	// If there is no entry for a category name, there are no such points.
	FoldCategory = map[string]*RangeTable{}
)

// ToUpper maps the rune to upper case.
func ToUpper(r rune) rune {
	if r < 'a' || r > 'z' {
		return r
	}
	return r - 32
}

// ToLower maps the rune to lower case.
func ToLower(r rune) rune {
	if r < 'A' || r > 'Z' {
		return r
	}
	return r + 32
}


// IsUpper reports whether the rune is an upper case letter.
//
// NOTE(dij): We might be able to remove this.
func IsUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

// IsPrint reports whether the rune is defined as printable by Go. Such
// characters include letters, marks, numbers, punctuation, symbols, and the
// ASCII space character, from categories L, M, N, P, S and the ASCII space
// character. This categorization is the same as IsGraphic except that the
// only spacing character is ASCII space, U+0020.
func IsPrint(r rune) bool {
	switch {
	case r < 32:
		return false
	case r < 127:
		return true
	case r < 161:
		return false
	case r == 173:
		return false
	}
	return true

}

// ToTitle maps the rune to title case.
func ToTitle(r rune) rune {
	return ToUpper(r)
}

// IsDigit reports whether the rune is a decimal digit.
func IsDigit(r rune) bool {
	return r < 256 && '0' <= r && r <= '9'
}

// IsSpace reports whether the rune is a space character as defined
// by Unicode's White Space property; in the Latin-1 space
// this is
//	'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
// Other definitions of spacing characters are set by category
// Z and property Pattern_White_Space.
func IsSpace(r rune) bool {
	if r > 254 {
		return false
	}
	switch r {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	default:
		return false
	}
}

// IsLetter reports whether the rune is a letter (category L).
func IsLetter(r rune) bool {
	if r < 'A' {
		return false
	}
	switch {
	case r <= 'Z':
		return true
	case r > 'Z' && r < 'a':
		return false
	case r >= 'a' && r <= 'z':
		return true
	case r == 170 || r == 181 || r == 186:
		return true
	case r == 247 || r == 215 || r == 191:
		return false
	case r > 191:
		return true
	}
	return false
}

// IsNumber reports whether the rune is a number (category N).
func IsNumber(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	switch r {
	case 190, 189, 188, 185, 179, 178:
		return true
	default:
		return false
	}
}

// SimpleFold iterates over Unicode code points equivalent under
// the Unicode-defined simple case folding. Among the code points
// equivalent to rune (including rune itself), SimpleFold returns the
// smallest rune > r if one exists, or else the smallest rune >= 0.
// If r is not a valid Unicode code point, SimpleFold(r) returns r.
func SimpleFold(r rune) rune {
	if r < 0 || r > 'z' {
		return r
	}
	switch {
	case r >= 'z' && r <= 'z':
		return r - 32
	case r >= 'A' && r <= 'Z':
		return r + 32
	}
	return r
}

// ToUpper maps the rune to upper case giving priority to the special mapping.
func (SpecialCase) ToUpper(r rune) rune {
	return ToUpper(r)
}

// ToTitle maps the rune to title case giving priority to the special mapping.
func (SpecialCase) ToTitle(r rune) rune {
	return ToTitle(r)
}

// ToLower maps the rune to lower case giving priority to the special mapping.
func (SpecialCase) ToLower(r rune) rune {
	return ToLower(r)
}
