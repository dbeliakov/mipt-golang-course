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

`ConcatReader` должен конкатенировать потоки байт из произвольного числа `io.Reader`. Поведение согласуйте с контрактом [`io.Reader`](https://pkg.go.dev/io#Reader) и с логикой [`io.MultiReader`](https://pkg.go.dev/io#MultiReader) из стандартной библиотеки:

- читатели обрабатываются по очереди; при `EOF` у текущего переходите к следующему;
- если очередной `Read` возвращает часть данных и `EOF`, при следующем вызове `Read` нужно продолжить со следующего читателя;
- **короткие `Read`**: один вызов `Read` может вернуть меньше байт, чем запрошено — при последующих вызовах нужно добирать данные из той же цепочки (в т.ч. через `io.ReadAll`);
- **нулевой читатель** (сразу `EOF`, например пустая строка) пропускается, данные идут из следующих.

Пример:

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
`TeeWriter` должен принимать произвольное количество `io.Writer` и мультиплексировать в них поступающий поток байт. Поведение как у [`io.MultiWriter`](https://pkg.go.dev/io#MultiWriter):

- вызывайте `Write` у каждого writer по очереди с **одним и тем же** слайсом `p`;
- при **первой** ошибке от очередного `Write` сразу верните `(n, err)` **из этого вызова** и **не** вызывайте последующие writers;
- если все предыдущие записи прошли без ошибки, но очередной `Write` вернул `n < len(p)` при `err == nil`, верните `(n, io.ErrShortWrite)` (короткая запись).

```go
tr := rw.NewTeeWriter(os.Stdout, os.Stderr)
_, _ = fmt.Fprintf(tr, "Hello, World!")
```

Для реализации рекомендуется изучить контракты, которым должны следовать интерфейсы [io.Reader](https://pkg.go.dev/io#Reader) и [io.Writer](https://pkg.go.dev/io#Writer).

В решении могут пригодиться:
- [hex.EncodedLen](https://pkg.go.dev/encoding/hex#EncodedLen)
- [hex.Encode](https://pkg.go.dev/encoding/hex#Encode)
