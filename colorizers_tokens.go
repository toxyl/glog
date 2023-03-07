package glog

import (
	"strconv"
	"strings"
	"time"
)

var tokenStringSeparators = []string{",", ".", ";", ":", "?", "!", ")", "}", "]", "'", "\""}
var tokenStringSeparatorsOpen = []string{"(", "{", "[", "'", "\""}

func isTokenStringSeparator(v string) bool {
	for _, s := range tokenStringSeparators {
		if s == v {
			return true
		}
	}
	return false
}

func isTokenStringSeparatorOpen(v string) bool {
	for _, s := range tokenStringSeparatorsOpen {
		if s == v {
			return true
		}
	}
	return false
}

type token struct {
	value         string
	color         int
	bold          bool
	italic        bool
	underline     bool
	strikeThrough bool
	next          *token
	prev          *token
}

func (t *token) String() string {
	res := ""
	if t.bold {
		res += Bold()
	}
	if t.italic {
		res += Italic()
	}
	if t.underline {
		res += Underline()
	}
	if t.strikeThrough {
		res += StrikeThrough()
	}
	if t.color > 0 {
		return res + Wrap(t.value, t.color)
	}
	if t.color < 0 {
		return res + Auto(t.value)
	}
	return res + t.value
}

func (t *token) first() *token {
	ct := t
	for ct.prev != nil {
		ct = ct.prev
	}
	return ct
}

func (t *token) last() *token {
	ct := t
	for ct.next != nil {
		ct = ct.next
	}
	return ct
}

// append moves to the last token and adds a next one
func (t *token) append(token *token) *token {
	ct := t.last()
	token.prev = ct
	ct.next = token
	ct = ct.next
	lastChar := string(ct.value[len(ct.value)-1])
	firstChar := string(ct.value[0])
	if ct.color != -1 && isTokenStringSeparator(lastChar) {
		ct.value = ct.value[:len(ct.value)-1]
		ct.insert(newToken(lastChar, -1, token.bold, token.italic, token.underline, token.strikeThrough))
	}
	if ct.color != -1 && isTokenStringSeparatorOpen(firstChar) {
		ct.value = ct.value[1:]
		ct.prev.insert(newToken(firstChar, -1, token.bold, token.italic, token.underline, token.strikeThrough))
	}

	return ct
}

// insert a token after the current position
func (t *token) insert(token *token) *token {
	ct := t
	ctn := t.next
	token.prev = ct
	token.next = ctn
	ct.next = token
	return ct
}

func (t *token) mergeWithNext(glue string) *token {
	if t.next != nil {
		t.value += glue + t.next.value
		t.next = t.next.next
	}
	return t
}

func (t *token) Merge(glue string) string {
	colors := []int{}
	glues := []string{}
	values := []string{}
	stringContent := ""
	ct := t.first()
	for ct != nil {
		for ct.next != nil && ct.next.color == ct.color {
			ct = ct.mergeWithNext(glue)
		}
		if ct.next != nil && isTokenStringSeparator(ct.next.value) {
			glues = append(glues, "")
		} else {
			if ct.next != nil {
				if isTokenStringSeparatorOpen(string(ct.value[len(ct.value)-1])) {
					glues = append(glues, "")
				} else {
					glues = append(glues, glue)
				}
			} else {
				glues = append(glues, "")
			}
		}
		colors = append(colors, ct.color)
		values = append(values, ct.value)

		if ct.color == -1 {
			stringContent += ct.value
		}
		ct = ct.next
	}
	col := stringColorCache.Get(stringContent)

	for i := range colors {
		if colors[i] == -1 {
			colors[i] = col
		}
	}

	output := ""
	for i := range values {
		output += Wrap(values[i], colors[i]) + Wrap(glues[i], col)
	}
	return output
}

func newToken(value string, color int, bold, italic, underline, strikeThrough bool) *token {
	return &token{
		value:         value,
		color:         color,
		bold:          bold,
		italic:        italic,
		underline:     underline,
		strikeThrough: strikeThrough,
		next:          nil,
		prev:          nil,
	}
}

type tokenSign int

const (
	sign_NEGATIVE = -1
	sign_ZERO     = 0
	sign_POSITIVE = 1
)

// getSign will take a string that is known (!!!) to be parsable as float
// and returns -1 if the value is negative, 1 if it's positive and 0 in all
// other cases.
func getSign(token string) tokenSign {
	v, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return sign_ZERO // we consider errors to be the zero-value, check your inputs!
	}
	if v > 0 {
		return sign_POSITIVE
	}
	if v < 0 {
		return sign_NEGATIVE
	}
	return sign_ZERO
}

func isTokenFloat(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func isTokenInt(token string) bool {
	_, err := strconv.ParseInt(token, 10, 64)
	return err == nil
}

func isTokenDate(token string) bool {
	_, err := time.Parse(LoggerConfig.DateFormat, token)
	return err == nil
}

func isTokenDateTime(token string) bool {
	_, err := time.Parse(LoggerConfig.DateTimeFormat, token)
	return err == nil
}

func isTokenDateTime12hr(token string) bool {
	_, err := time.Parse(LoggerConfig.DateTimeFormat12hr, token)
	return err == nil
}

func isTokenTime(token string) bool {
	_, err := time.Parse(LoggerConfig.TimeFormat, token)
	return err == nil
}

func isTokenTime12hr(token string) bool {
	_, err := time.Parse(LoggerConfig.TimeFormat12hr, token)
	return err == nil
}

func isTokenTimeFormat(token string) bool {
	return isTokenDate(token) || isTokenDateTime(token) || isTokenDateTime12hr(token) || isTokenTime(token) || isTokenTime12hr(token)
}

func getTokenColor(token string) int {
	token = strings.TrimSpace(token)
	if token == "" {
		return -1
	}

	lastChar := string(token[len(token)-1])
	firstChar := string(token[0])
	isLastSep := isTokenStringSeparator(lastChar)
	isFirstSep := isTokenStringSeparatorOpen(firstChar)
	if isLastSep {
		token = token[:len(token)-1]
		lastChar = string(token[len(token)-1])
	}
	if isFirstSep {
		token = token[1:]
	}

	tl := strings.ToLower(token)
	if tl == "false" {
		return LoggerConfig.ColorBoolFalse
	}

	if tl == "true" {
		return LoggerConfig.ColorBoolTrue
	}

	if tl == "nil" {
		return LoggerConfig.ColorNil
	}

	if isTokenInt(token) {
		switch getSign(token) {
		case sign_POSITIVE:
			return LoggerConfig.ColorIntPositive
		case sign_NEGATIVE:
			return LoggerConfig.ColorIntNegative
		}
		return LoggerConfig.ColorIntZero
	}

	if isTokenFloat(token) {
		switch getSign(token) {
		case sign_POSITIVE:
			return LoggerConfig.ColorFloatPositive
		case sign_NEGATIVE:
			return LoggerConfig.ColorFloatNegative
		}
		return LoggerConfig.ColorFloatZero
	}

	if lastChar == "%" && isTokenFloat(token[:len(token)-1]) {
		switch getSign(token[:len(token)-1]) {
		case sign_POSITIVE:
			return LoggerConfig.ColorPercentagePositive
		case sign_NEGATIVE:
			return LoggerConfig.ColorPercentageNegative
		}
		return LoggerConfig.ColorPercentageZero
	}

	if isTokenTimeFormat(token) {
		return LoggerConfig.ColorTime
	}

	// if isSep {
	// 	// this usually happens in text with lists and punctuation
	// 	// token = token[:len(token)-1]
	// 	return -1
	// }

	// TODO: figure out a good way to detect durations

	return -1 // we treat everything else as string
}

// var reWordBoundary = regexp.MustCompile(`\s`)

// Token takes a list of strings and will try the following for each one of them:
//
//  - Is int?
//  - Is float?
//  - Is percentage?
//  - Is bool?
//  - Is nil?
//  - Is date, time or datetime?
//  - Is duration?
//
// If neither is the case the token is assumed to be a string and split on spaces,
// each of the new tokens will be processed with Token(), i.e. processing can happen
// recursively. Once all tokens are processed the string will be rebuild using the color
// of the combined string-tokens left over and all others colored according to their type.
//
//
// Related config setting(s):
//
func Token(glue string, tokens ...string) string {
	var res *token
	for _, value := range tokens {
		tc := getTokenColor(value)
		token := newToken(value, tc, false, false, false, false)

		if tc == -1 {
			// this a string token, so we have to parse it recursively
			// let's first fix a case that can happen with Wikipedia quotes:
			// footnote references at the end of a sentence look like: .[7]
			value = strings.ReplaceAll(value, ".[", "."+glue+"[")
			tkns := strings.Split(value, glue)
			if len(tkns) > 1 {
				token.value = Token(glue, tkns...)
			}
		}

		if res == nil {
			res = token
			continue
		}

		res = res.append(token)
	}
	return res.Merge(glue)
}
