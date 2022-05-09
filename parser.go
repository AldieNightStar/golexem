package golexem

import (
	"strconv"
	"strings"
)

const digits = "0123456789"

type ETC string

func Parse(src string) []any {
	ptr := 0
	toks := make([]any, 0, 8)
	for {
		if ptr >= len(src) {
			break
		}
		if src[ptr] == ' ' || src[ptr] == '\t' || src[ptr] == '\n' || src[ptr] == '\r' {
			ptr += 1
			continue
		}
		str, cnt := parseString(src[ptr:])
		if cnt > 0 {
			toks = append(toks, str)
			ptr += cnt
		}
		num, cnt := ParseFloat(src[ptr:])
		if cnt > 0 {
			toks = append(toks, num)
			ptr += cnt
		}
		cnt = ParseComment(src[ptr:])
		if cnt > 0 {
			ptr += cnt
		}
		etc, cnt := parseEtc(src[ptr:])
		if cnt > 0 {
			toks = append(toks, ETC(etc))
			ptr += cnt
		}
	}
	return toks
}

func parseEtc(s string) (result string, count int) {
	sb := strings.Builder{}
	sb.Grow(32)
	for _, c := range s {
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			break
		}
		sb.WriteRune(c)
	}
	return sb.String(), sb.Len()
}

func parseString(s string) (result string, count int) {
	if len(s) < 2 {
		return "", 0
	}
	if !(s[0] == '\'' || s[0] == '"' || s[0] == '`') {
		return "", 0
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
			return "", 0
		}
		if c == end {
			ptr += 1
			break
		}
		ptr += 1
		sb.WriteRune(c)
	}
	return sb.String(), ptr
}

func ParseFloat(s string) (result float64, count int) {
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
		return n, ptr
	} else {
		return 0, 0
	}
}

func ParseComment(s string) (count int) {
	if len(s) < 1 {
		return 0
	}
	if s[0] != '#' {
		return 0
	}
	ptr := 0
	for _, c := range s {
		ptr += 1
		if c == '\n' {
			break
		}
	}
	return ptr
}
