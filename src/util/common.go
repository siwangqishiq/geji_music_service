package util

func Is[T any](value bool, left T, right T) T {
	if value {
		return left
	} else {
		return right
	}
}
