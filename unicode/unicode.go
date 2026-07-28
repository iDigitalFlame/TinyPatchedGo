// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// 2022 - Shrinkage by iDigitalFlame

// Package unicode provides data and functions to test some properties of
// Unicode code points.
package unicode

const (
	MaxRune         = '\U0010FFFF' // Maximum valid Unicode code point.
	ReplacementChar = '\uFFFD'     // Represents invalid code points.
	MaxASCII        = '\u007F'     // maximum ASCII value.
	MaxLatin1       = '\u00FF'     // maximum Latin-1 value.
)

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

// RangeTable defines a set of Unicode code points by listing the ranges of
// code points within the set. The ranges are listed in two slices
// to save space: a slice of 16-bit ranges and a slice of 32-bit ranges.
// The two slices must be in sorted order and non-overlapping.
// Also, R32 should contain only values >= 0x10000 (1<<16).
type RangeTable struct {
	R16         []Range16
	R32         []Range32
	LatinOffset int // number of entries in R16 with Hi <= MaxLatin1
}

// Range16 represents of a range of 16-bit Unicode code points. The range runs from Lo to Hi
// inclusive and has the specified stride.
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

// CaseRange represents a range of Unicode code points for simple (one
// code point to one code point) case conversion.
// The range runs from Lo to Hi inclusive, with a fixed stride of 1. Deltas
// are the number to add to the code point to reach the code point for a
// different case for that character. They may be negative. If zero, it
// means the character is in the corresponding case. There is a special
// case representing sequences of alternating corresponding Upper and Lower
// pairs. It appears with a fixed Delta of
//
//	{UpperLower, UpperLower, UpperLower}
//
// The constant UpperLower has an otherwise impossible delta value.
type CaseRange struct {
	Lo    uint32
	Hi    uint32
	Delta [4]rune
}

// SpecialCase represents language-specific case mappings such as Turkish.
// Methods of SpecialCase customize (by overriding) the standard mappings.
type SpecialCase []CaseRange

// ToUpper maps the rune to upper case.
func ToUpper(r rune) rune {
	if r < 'a' || r > 'z' {
		return r
	}
	return r - 32
}

// IsUpper reports whether the rune is an upper case letter.
//
// NOTE(dij): We might be able to remove this.
func IsUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

// IsDigit reports whether the rune is a decimal digit.
func IsDigit(r rune) bool {
	return r < 256 && '0' <= r && r <= '9'
}

// IsLower reports whether the rune is a lower case letter.
func IsLower(r rune) bool {
	return r >= 'a' && r <= 'z'
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

// IsSpace reports whether the rune is a space character as defined
// by Unicode's White Space property; in the Latin-1 space
// this is
//
//	'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
//
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

// ToLower maps the rune to lower case.
func ToLower(r rune) rune {
	if r < 'A' || r > 'Z' {
		return r
	}
	return r + 32
}

// ToTitle maps the rune to title case.
func ToTitle(r rune) rune {
	return ToUpper(r)
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

// IsGraphic reports whether the rune is defined as a Graphic by Unicode.
// Such characters include letters, marks, numbers, punctuation, symbols, and
// spaces, from categories L, M, N, P, S, Zs.
func IsGraphic(r rune) bool {
	if r >= 0x20 && r <= 0x7E {
		return true
	}
	if r >= 0xA1 && r <= 0xAC {
		return true
	}
	return r >= 0xAE
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

// Is reports whether the rune is in the specified table of ranges.
func Is(v *RangeTable, r rune) bool {
	x := v.R16
	if len(x) > 0 && r <= rune(x[len(x)-1].Hi) {
		return is16(x, uint16(r))
	}
	y := v.R32
	if len(y) > 0 && r >= rune(y[0].Lo) {
		return is32(y, uint32(r))
	}
	return false
}
func is16(v []Range16, r uint16) bool {
	if len(v) <= 18 || r <= MaxLatin1 {
		for i := range v {
			x := &v[i]
			if r < x.Lo {
				return false
			}
			if r <= x.Hi {
				return x.Stride == 1 || (r-x.Lo)%x.Stride == 0
			}
		}
		return false
	}
	l, h := 0, len(v)
	for l < h {
		e := l + (h-l)/2
		x := &v[e]
		if x.Lo <= r && r <= x.Hi {
			return x.Stride == 1 || (r-x.Lo)%x.Stride == 0
		}
		if r < x.Lo {
			h = e
		} else {
			l = e + 1
		}
	}
	return false
}
func is32(v []Range32, r uint32) bool {
	if len(v) <= 18 {
		for i := range v {
			x := &v[i]
			if r < x.Lo {
				return false
			}
			if r <= x.Hi {
				return x.Stride == 1 || (r-x.Lo)%x.Stride == 0
			}
		}
		return false
	}
	l, h := 0, len(v)
	for l < h {
		e := l + (h-l)/2
		x := &v[e]
		if x.Lo <= r && r <= x.Hi {
			return x.Stride == 1 || (r-x.Lo)%x.Stride == 0
		}
		if r < x.Lo {
			h = e
		} else {
			l = e + 1
		}
	}
	return false
}

// ToLower maps the rune to lower case giving priority to the special mapping.
func (SpecialCase) ToLower(r rune) rune {
	return ToLower(r)
}

// ToTitle maps the rune to title case giving priority to the special mapping.
func (SpecialCase) ToTitle(r rune) rune {
	return ToTitle(r)
}

// ToUpper maps the rune to upper case giving priority to the special mapping.
func (SpecialCase) ToUpper(r rune) rune {
	return ToUpper(r)
}
