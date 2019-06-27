package bob

import (
	"strings"
)

// Hey if a function that returns the comments of a teenager named Bob.
func Hey(remark string) string {

	trimmedRemark := strings.TrimSpace(remark)
	allCapsRemark := strings.ToUpper(trimmedRemark)
	isAllCaps := allCapsRemark == trimmedRemark && allCapsRemark != strings.ToLower(allCapsRemark)
	isQuestion := strings.HasSuffix(trimmedRemark, "?")

	if trimmedRemark == "" {
		return "Fine. Be that way!"
	}
	if isQuestion && isAllCaps {
		return "Calm down, I know what I'm doing!"
	}
	if isAllCaps {
		return "Whoa, chill out!"
	}
	if isQuestion {
		return "Sure."
	}

	return "Whatever."
}
