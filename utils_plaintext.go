package glog

import (
	"bytes"
	"regexp"
	"unicode"
	"unicode/utf8"
)

var (
	reANSI = regexp.MustCompile(`\x1b\[(?:[0-9]{1,4}(?:;[0-9]{0,4})*)?[mABCDHfKsu]`)
)

func StripANSI(str string) string {
	return reANSI.ReplaceAllStringFunc(str, func(s string) string {
		return ""
	})
}

func ReplaceRunes(str string, replacement string, list []rune) string {
	var buf bytes.Buffer
	for _, r := range str {
		m := false
		for _, ru := range list {
			if ru == r {
				m = true
				break
			}
		}
		if m {
			buf.WriteString(replacement)
			continue
		}
		buf.WriteRune(r)
	}
	return buf.String()
}

func ReplaceEmojis(str string, replacement string) string {
	var buf bytes.Buffer
	for _, r := range str {
		if unicode.In(r, unicode.So, unicode.Sk) {
			buf.WriteString(replacement)
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func RemoveNonPrintable(str string) string {
	str = StripANSI(str)
	var buf bytes.Buffer
	for _, r := range str {
		if r == 0x00 {
			continue
		}
		if r >= 0x20 && r <= 0x7E || unicode.IsSpace(r) || (r >= 0x1F600 && r <= 0x1F64F) || (r >= 0x1F300 && r <= 0x1F5FF) || (r >= 0x1F680 && r <= 0x1F6FF) || (r >= 0x2600 && r <= 0x26FF) || (r >= 0x2700 && r <= 0x27BF) || (r >= 0x1F900 && r <= 0x1F9FF) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func PlaintextString(str string) string {
	str = RemoveNonPrintable(str)
	return str
}

func PlaintextStringLength(str string) int {
	str = ReplaceRunes(str, " ", []rune{'μ', 'µ'}) // for padding calculations we actually need to count emojis as two
	str = ReplaceEmojis(str, " ")
	str = RemoveNonPrintable(str)

	return utf8.RuneCountInString(str)
}

func plaintextStringLengthForPadding(str string) int {
	str = ReplaceRunes(str, " ", []rune{'μ', 'µ'}) // for padding calculations we actually need to count emojis as two chars
	str = ReplaceEmojis(str, "  ")                 // for padding calculations we actually need to count emojis as two chars
	str = RemoveNonPrintable(str)

	return utf8.RuneCountInString(str)
}
