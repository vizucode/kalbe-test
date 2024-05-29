package formatting

import "regexp"

func PhoneFormatting(rawPhoneNumber string) (resp string) {
	regex := regexp.MustCompile(`^0`)
	return regex.ReplaceAllString(rawPhoneNumber, "62")
}
