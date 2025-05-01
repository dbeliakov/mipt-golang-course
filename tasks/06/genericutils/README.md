# Generic Utilities

Реализуйте набор универсальных функций для работы с обобщенными типами (generics), которые часто требуются в проектах.

## Требуемые функции

```go
// Filter - фильтрация элементов слайса по условию
func Filter[T any](s []T, pred func(T) bool) []T

// GroupBy - группировка элементов по ключу
func GroupBy[T any, K comparable](s []T, grouper func(T) K) map[K][]T

// MaxBy - поиск максимального элемента по критерию
func MaxBy[T any](s []T, less func(a, b T) bool) T

// Repeat - создание слайса с повторяющимся значением
func Repeat[T any](val T, times int) []T

// JSONParse - парсинг JSON в указанный тип
func JSONParse[T any](data []byte) (T, error)

// Dedup - удаление дубликатов из слайса
func Dedup[T comparable](s []T) []T
```

Вам могут пригодиться:
- [Go Generics Tutorial](https://go.dev/doc/tutorial/generics)
- [Type Parameters Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)
