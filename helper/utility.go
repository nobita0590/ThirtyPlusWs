package helper

import (
	"unicode"
	"strings"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func ReplaceUTF8Character(input string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, input)
	result = strings.Replace(strings.Trim(strings.ToLower(result)," "),"đ","d",-1)
	return  result
}

func PrettyUrl(title string) string {
	if title == "" {
		return title
	}
	title = strings.Replace(ReplaceUTF8Character(title)," ","-",-1)
	title = strings.Replace(title,"/","",-1)
	return title
}
func StringInSlice(a string, list []string) int {
	for k, b := range list {
		if strings.ToLower(b) == strings.ToLower(a) {
			return k
		}
	}
	return -1
}