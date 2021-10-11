package util

import (
	"net/mail"

	uuid "github.com/satori/go.uuid"
)

// IsValidUUID is validate uuid v4
func IsValidUUID(u string) bool {
	_, err := uuid.FromString(u)
	return err == nil
}

// ValidEmailAddress mail.ParseAddress wraper
func ValidEmailAddress(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
