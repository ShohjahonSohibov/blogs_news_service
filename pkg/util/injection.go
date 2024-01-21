package util

import (
	"regexp"
	"strings"
)

func IsSQLInjection(str string) bool {
	injectionPattern := `(?:--|\b(?:SELECT|INSERT|PG_SLEEP|UPDATE|DELETE|FROM|WHERE)\b)`

	// Case-insensitive matching
	re := regexp.MustCompile(injectionPattern)

	return re.MatchString(strings.ToUpper(str))
}

