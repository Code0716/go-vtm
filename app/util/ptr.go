package util

// LiteralToPtrGenerics returns pointer
func LiteralToPtrGenerics[T comparable](x T) *T {
	return &x
}
