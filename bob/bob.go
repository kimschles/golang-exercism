package bob

import (
	"strconv"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	if remark == "" {
		return "Fine. Be that way!"
	} else if remark == strings.ToUpper(remark) && "false" == strconv.FormatBool(strings.ContainsAny(remark, "?")) {
		return "Whoa, chill out!"
	} else if remark == strings.ToUpper(remark) && "true" == strconv.FormatBool(strings.ContainsAny(remark, "?")) {
		return "Calm down, I know what I'm doing!"
	} else if "true" == strconv.FormatBool(strings.ContainsAny(remark, "?")) {
		return "Sure."
	} else {
		return "Whatever."
	}
}

// To do: figure out uppercase conditions
