// The acronym package converts a phrase to its acronym

package acronym

import "strings"

// Abbreviate takes a phrase and returns an acronym
func Abbreviate(s string) string {
	sanitizedString := strings.ReplaceAll(s, "-", " ")
	list := strings.Fields(sanitizedString)

	letterList := []string{}
	for _, word := range list {
		x := strings.ToUpper(word)
		letterList = append(letterList, string(x[0]))
	}
	acronym := strings.Join(letterList, "")
	return acronym
}
