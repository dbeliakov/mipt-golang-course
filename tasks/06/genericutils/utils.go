package genericutils

func Filter[T any](s []T, pred func(T) bool) []T {
	panic("implement me")
}

func GroupBy[T any, K comparable](s []T, grouper func(T) K) map[K][]T {
	panic("implement me")
}

func MaxBy[T any](s []T, less func(a T, b T) bool) T {
	panic("implement me")
}

func Repeat[T any](val T, times int) []T {
	panic("implement me")
}

func JSONParse[T any](data []byte) (T, error) {
	panic("implement me")
}

func Dedup[T comparable](s []T) []T {
	panic("implement me")
}
