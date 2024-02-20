package crlfuzz

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// GenerateURL should generate for potential vulnerability URLs
func GenerateURL(u string) []string {
	var url []string

	if !strings.HasSuffix(u, "/") {
		u += "/"
	}

	for _, a := range appendList {
		for _, e := range escapeList {
			url = append(url, fmt.Sprint(u, a, DecodeUnicode(e), keyHeader, "%3a%20", valHeader))
		}
	}

	return url
}

func DecodeUnicode(s string) string {
	re := regexp.MustCompile(`%u[0-9A-Fa-f]{4}`)
	decodedString := re.ReplaceAllStringFunc(s, func(match string) string {
		unicodeStr := match[2:]
		unicodeInt, _ := strconv.ParseInt(unicodeStr, 16, 32)

		return string(rune(unicodeInt))
	})
	decodedString = strings.ReplaceAll(decodedString, "%u", "\\u")

	return decodedString
}
