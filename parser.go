package golexem

import (
	"strconv"
	"strings"
)

const digits = "0123456789"

type genericToken struct {
	Value       string
	ValueNumber float64
	LineNumber  int
}

type ETC *genericToken
type STRING *genericToken
type NUMBER *genericToken
type COMMENT *genericToken

// Gets token number or -1 if fail
func GetTokenLine(tok any) int {
	if etc, ok := tok.(ETC); ok {
		return etc.LineNumber
	} else if str, ok := tok.(STRING); ok {
		return str.LineNumber
	} else if num, ok := tok.(NUMBER); ok {
		return num.LineNumber
	} else if cmt, ok := tok.(COMMENT); ok {
		return cmt.LineNumber
	}
	return -1
}

func NewToken(val string, valn float64) *genericToken {
	return &genericToken{
		Value:       val,
		ValueNumber: valn,
		LineNumber:  0,
	}
}

func Parse(src string) []any {
	ptr := 0
	toks := make([]any, 0, 8)
	line := 1
	for {
		if ptr >= len(src) {
			break
		}
		if src[ptr] == ' ' || src[ptr] == '\t' || src[ptr] == '\n' || src[ptr] == '\r' {
			if src[ptr] == '\n' {
				line += 1
			}
			ptr += 1
			continue
		}
		str, cnt := parseString(src[ptr:])
		if cnt > 0 {
			str.LineNumber = line
			toks = append(toks, str)
			ptr += cnt
		}
		num, cnt := parseFloat(src[ptr:])
		if cnt > 0 {
			num.LineNumber = line
			toks = append(toks, num)
			ptr += cnt
		}
		cmt, cnt := parseComment(src[ptr:])
		if cnt > 0 {
			cmt.LineNumber = line
			toks = append(toks, cmt)
			ptr += cnt
		}
		etc, cnt := parseEtc(src[ptr:])
		if cnt > 0 {
			etc.LineNumber = line
			toks = append(toks, ETC(etc))
			ptr += cnt
		}
	}
	return toks
}

func parseEtc(s string) (tok ETC, count int) {
	sb := strings.Builder{}
	sb.Grow(32)
	for _, c := range s {
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			break
		}
		sb.WriteRune(c)
	}
	return ETC(NewToken(sb.String(), 0)), sb.Len()
}

func parseString(s string) (tok STRING, count int) {
	if len(s) < 2 {
		return nil, 0
	}
	if !(s[0] == '\'' || s[0] == '"' || s[0] == '`') {
		return nil, 0
	}
	ptr := 1
	end := rune(s[0])
	esc := false
	sb := strings.Builder{}
	sb.Grow(64)
	for _, c := range s[1:] {
		if esc {
			esc = false
			if c == 'n' {
				sb.WriteRune('\n')
			} else if c == 't' {
				sb.WriteRune('\t')
			} else if c == 'r' {
				sb.WriteRune('\r')
			} else if c == '0' {
				sb.WriteRune(0)
			} else {
				sb.WriteRune(c)
			}
			ptr += 1
			continue
		}
		if c == '\\' {
			ptr += 1
			esc = true
			continue
		}
		if c == '\n' {
			return nil, 0
		}
		if c == end {
			ptr += 1
			break
		}
		ptr += 1
		sb.WriteRune(c)
	}
	return STRING(NewToken(sb.String(), 0)), ptr
}

func parseFloat(s string) (result NUMBER, count int) {
	sb := strings.Builder{}
	sb.Grow(8)
	dotAllow := true
	ptr := 0
	for i, c := range s {
		if c == '.' && i == 0 && dotAllow {
			sb.WriteString("0.")
			dotAllow = false
			ptr += 1
		} else if c == '-' && i == 0 {
			ptr += 1
			sb.WriteRune('-')
		} else if strings.Contains(digits, string(c)) {
			sb.WriteRune(c)
			ptr += 1
		} else if c == '.' && dotAllow {
			dotAllow = false
			ptr += 1
			sb.WriteRune('.')
		} else {
			break
		}
	}
	if n, err := strconv.ParseFloat(sb.String(), 64); err == nil {
		return NUMBER(NewToken("", n)), ptr
	} else {
		return nil, 0
	}
}

func parseComment(s string) (result COMMENT, count int) {
	if len(s) < 1 {
		return nil, 0
	}
	if s[0] != '#' {
		return nil, 0
	}
	ptr := 0
	for _, c := range s {
		if c == '\n' {
			break
		}
		ptr += 1
	}
	newStr := strings.TrimPrefix(s[1:ptr], " ")
	return COMMENT(NewToken(newStr, 0)), ptr
}
