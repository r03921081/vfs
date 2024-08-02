package util

import (
	"fmt"
	"r03921081/vfs/constant"
	"regexp"
)

var (
	ValidCommand = regexp.MustCompile(fmt.Sprintf(`^[a-zA-Z0-9\+\-\*/_\@\[\]\(\)\{\}\.\s\-]{1,%d}$`, constant.MaxLengthCommand))

	ValidName = regexp.MustCompile(fmt.Sprintf(`^[a-zA-Z0-9]{1,%d}$`, constant.MaxLengthName))

	// Description validation
	ValidDescription = regexp.MustCompile(fmt.Sprintf(`^[a-zA-Z0-9\+\-\*/_\@\[\]\(\)\{\}\.\s\-]{1,%d}$`, constant.MaxLengthDescription))
)

func IsValidInput(input string, r *regexp.Regexp) bool {
	return r.MatchString(input)
}
