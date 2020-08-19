package utils

import (
	"unicode/utf8"
)

const (
	lowerhex = "0123456789abcdef"
)

func RuneToASCII(r rune) string {
	buf := make([]byte, 0)
	if !utf8.ValidRune(r) {
		r = utf8.RuneError
	}
	switch {
	case r < ' ':
		buf = append(buf, `\x`...)
		buf = append(buf, lowerhex[byte(r)>>4])
		buf = append(buf, lowerhex[byte(r)&0xF])
	case r > utf8.MaxRune:
		r = 0xFFFD
		fallthrough
	case r < 0x10000:
		buf = append(buf, `\u`...)
		for s := 12; s >= 0; s -= 4 {
			buf = append(buf, lowerhex[r>>uint(s)&0xF])
		}
	default:
		buf = append(buf, `\U`...)
		for s := 28; s >= 0; s -= 4 {
			buf = append(buf, lowerhex[r>>uint(s)&0xF])
		}
	}
	return string(buf)
}
