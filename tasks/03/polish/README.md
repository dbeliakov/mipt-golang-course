# Вычисление обратной польской нотации

Реализуйте функцию `Calculate(expr string) (int, error)`, 
которая вычисляет значение выражения, записанного в [обратной польской нотации](https://en.wikipedia.org/wiki/Reverse_Polish_notation).

На вход функция получает строку, состоящую из цифр и операторов `+`, `-`, `*`. 
Гарантируется, что все элементы выражения будут разделены одним пробелом. 

Помимо значения выражения функция возвращает ошибку в случае, если само выражение оказалось невалидно.
В таком случае вам нужно вернуть уже существующую ошибку `ErrInvalidExpression`. 
Примеры невалидных выражений вы можете найти в тестах.

Пример:
```go
res, err := Calculate("3 5 +")
fmt.Println(res, err)  // 8 nil
```

Вам могут пригодиться:
- [strings.Split](https://pkg.go.dev/strings#Split)
- [strconv.Atoi](https://pkg.go.dev/strconv#Atoi)
