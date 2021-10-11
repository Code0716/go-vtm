package util

import (
	uuid "github.com/satori/go.uuid"
)

// UUIDGenerator is generate uuid
func UUIDGenerator() string {
	return uuid.NewV4().String()
}
