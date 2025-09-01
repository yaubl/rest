package snowflake

import (
	"strconv"
)

// simple utility to check if a string
// is a valid discord snowflake.
func IsSnowflake(s string) bool {
	if s == "" {
		return false
	}

	if s[0] == '-' || s[0] == '+' {
		return false
	}

	_, err := strconv.ParseUint(s, 10, 64)
	return err == nil
}
