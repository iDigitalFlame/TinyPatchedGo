// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// 2022 - Shrinkage by iDigitalFlame

package unicode

//ReplacementChar represents an invalid code points.
const ReplacementChar = '\uFFFD'

// SpecialCase represents language-specific case mappings such as Turkish.
// Methods of SpecialCase customize (by overriding) the standard mappings.
type SpecialCase bool

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
