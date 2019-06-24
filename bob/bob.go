package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	allCapsRemark := strings.ToUpper(remark)
	isAllCaps := allCapsRemark == remark && allCapsRemark != strings.ToLower(allCapsRemark)
	isQuestion := strings.HasSuffix(remark, "?")

	if remark == "" {
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

// To do: figure out "      "
//https://gistpages.com/posts/go-lang-check-if-a-string-is-empty
