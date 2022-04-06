package collections


func Duplicate[T any](x T) []T {
	return []T{x, x}
}