package util

func Is[T any](value bool, left T, right T) T {
	if value {
		return left
	} else {
		return right
	}
}

// AddAll 将 src 的所有元素追加到 dst，返回新切片
func AddAll[T any](dst, src []T) []T {
	return append(dst, src...)
}
