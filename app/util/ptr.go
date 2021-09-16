package util

import "time"

// BoolPtr returns pointer to bool value.
func BoolPtr(v bool) *bool {
	return &v
}

// IntPtr returns pointer to int value.
func IntPtr(v int) *int {
	return &v
}

// Int64Ptr returns pointer to int value.
func Int64Ptr(v int64) *int64 {
	return &v
}

// StrPtr returns pointer to string value.
func StrPtr(v string) *string {
	return &v
}

// TimePtr returns pointer to time.Time value.
func TimePtr(v time.Time) *time.Time {
	return &v
}

// Float32Ptr returns pointer to float32 value.
func Float32Ptr(v float32) *float32 {
	return &v
}
