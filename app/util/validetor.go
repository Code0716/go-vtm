package util

import (
	"net/mail"
	"regexp"
)

// IsValidUUID is validate uuid v4
func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

// ValidEmailAddress mail.ParseAddress wraper
func ValidEmailAddress(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
