package logic

import "unicode"

func TrimSpaceAndSlash(r rune) bool {
	return unicode.IsSpace(r) || (r == rune('/'))
}
