package macos

import (
	"bytes"
	"strconv"
	"strings"
)

func strip(value string) string {
	newValue, err := strconv.Unquote(strings.TrimSpace(value))
	if err == nil {
		return strings.TrimRight(strings.TrimLeft(newValue, "<"), ">")
	}
	return strings.TrimRight(strings.TrimLeft(value, "<"), ">")
}

func stripValueBytes(b []byte) string {
	b = bytes.TrimSpace(b)
	length := len(b)
	if length < 2 {
		return string(b)
	}
	if b[0] == '<' && b[length-1] == '>' {
		b = b[1 : length-1]
	}
	v, err := strconv.Unquote(string(b))
	if err == nil {
		return v
	}
	return string(b)
}
