package crlfuzz

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// GenerateURL should generate for potential vulnerability URLs
func GenerateURL(u string) ([]string, error) {
	var generated_urls []string

	parsed_url, e := url.Parse(u)

	if len(parsed_url.Path) == 0 {
		parsed_url.Path = "/"
	}

	fragment := parsed_url.Fragment
	parsed_url.Fragment = ""

	if e != nil {
		return nil, errors.New(e.Error())
	}

	for _, a := range appendList {
		for _, e := range escapeList {
			generated_urls = append(generated_urls, fmt.Sprint(parsed_url, a, DecodeUnicode(e), keyHeader, "%3a%20", valHeader, "#", fragment))
		}
	}

	return generated_urls, nil
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
