# Read and Write

Потоковые чтение и запись очень часто встречаются в Go:
- https://pkg.go.dev/os#File.ReadFrom
- https://pkg.go.dev/net/http#Response
- https://pkg.go.dev/net/http#Header.Write
- https://pkg.go.dev/compress/gzip#Writer


Попробуйте потренироваться в реализации потоковых чтений и записей с помощью нескольких небольших реализаций.

## CountReader

`CountReader` -- бесконечный поток байт, который возвращает байты от 0 до 9:
```go
r := rw.NewCountReader()
p := make([]byte, 3)
_, _ = r.Read(p) // p contains [0, 1, 2]
_, _ = r.Read(p) // p contains [3, 4, 5]
...
_, _ = r.Read(p) // p contains [9, 0, 1]
...
```

## LimitReader

`LimitReader` должен оборачивать другой `io.Reader` и возвращать ошибку `ErrLimitExceeded` в случае, если было прочитано больше n байт.

## ConcatReader

`ConcatReader` должен конкатенировать потоки байт из произвольного числа `io.Reader`. Когда n-ый `io.Reader` заканчивается, `ConcatReader` должен продолжать писать в переданный ему слайс, пока в нем осталось место:

```go
r1 := strings.NewReader("a")
r2 := strings.NewReader("bcdef")
cr := rw.NewConcatReader(r1, r2)
p := make([]byte, 3)
_, _ = cr.Read(p) // p contains abc
```

## HexWriter
`HexWriter` должен оборачивать другой `io.Writer` и записывать в него поток байт, который перед этим кодируется в hex

## TeeWriter
`TeeWriter` должен принимать произвольное количество `io.Writer` и мультиплексировать в них поступающий поток байт. Если в процессе одной из записей произойдет ошибка, `TeeWriter` должен перестать писть в последующие `io.Writer`.

```go
tr := rw.NewTeeWriter(os.Stdout, os.Stderr)
_, _ = fmt.Fprintf(tr, "Hello, World!")
```

Для реализации рекомендуется изучить контракты, которым должны следовать интерфейсы [io.Reader](https://pkg.go.dev/io#Reader) и [io.Writer](https://pkg.go.dev/io#Writer).

В решении могут пригодиться:
- [hex.EncodingLen](https://pkg.go.dev/encoding/hex#EncodedLen)
- [hex.Encode](https://pkg.go.dev/encoding/hex#Encode)
