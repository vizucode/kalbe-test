package formatting

import (
	"strings"
)

// will return index of value sorted by asc, if not found will return len of array
func StringSearch(arr []string, x string) (resp bool) {
	for _, res := range arr {
		if strings.EqualFold(res, x) {
			return true
		}
	}
	return false
}
