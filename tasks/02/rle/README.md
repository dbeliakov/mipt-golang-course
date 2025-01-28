# RLE-сжатие

Реализуйте функцию `RLECompress`, которая производит RLE-сжатие строки: на вход подается строка, состоящая из часто повторяющихся символов латинского алфавита (a-z, A-Z). Функция должна вернуть строку, в которой все идущие подряд одинаковые символы заменены на один экземпляр символа и количество повторений:
`aabbc -> a2b2c1`.

В данной задаче есть бенчмарк, проверяющий, насколько эффективна ваша реализация алгоритма. Для того, чтобы не выделять лишнюю память, изучите [`strings.Builder`](https://pkg.go.dev/strings#Builder).

Результат бенчмарка авторского решения:
```
goos: darwin
goarch: arm64
pkg: github.com/dbeliakov/mipt-golang-course/tasks/02/rle
BenchmarkRLECompress
BenchmarkRLECompress-12           226419              4500 ns/op            5368 B/op         10 allocs/op
PASS
```

# Запуск бенчмарков и анализ профилей

Для запуска бенчмарков с профайлером на примере задачи rle можно использовать следующую команду:
```shell
go test -v -run='^$' -bench='.*' -memprofile=mem.out ./tasks/02/rle/...
```

Семплы профайлера будут находиться в `mem.out`. Анализировать полученные семплы удобно с помощью встроенной утилиты `pprof`:

```shell
go tool pprof mem.out
```
```
File: rle.test
Type: alloc_space
Time: Jan 25, 2025 at 4:17pm (MSK)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)
```

Утилита запускает нтерактивную среду, в которой можно анализировать полученные профили, используя различные команды. Чтобы узнать, какие команды есть у `pprof`, нужно вызвать `help`. Для того, чтобы узнать, как вызывать конкретную команду, нужно вызвать `help [имя команды]`.

## Полезные команды

Команда `top` покажет, какие функции выделяли больше всего памяти:
```
(pprof) top -cum
Showing nodes accounting for 22.30GB, 100% of 22.30GB total
Dropped 1 node (cum <= 0.11GB)
      flat  flat%   sum%        cum   cum%
         0     0%     0%    22.30GB   100%  github.com/dbeliakov/mipt-golang-course/tasks/02/rle.BenchmarkRLECompress
   22.30GB   100%   100%    22.30GB   100%  github.com/dbeliakov/mipt-golang-course/tasks/02/rle.RLECompress (inline)
         0     0%   100%    22.30GB   100%  testing.(*B).runN
         0     0%   100%    22.30GB   100%  testing.(*B).launch
```

Для того, чтобы проанализировать функцию, которая делает слишком много аллокаций, можно использовать команду `list`, которая покажет, на каких строчках происходили аллокации:
```
(pprof) list github.com/dbeliakov/mipt-golang-course/tasks/02/rle.RLECompress
Total: 22.30GB
ROUTINE ======================== github.com/dbeliakov/mipt-golang-course/tasks/02/rle.RLECompress in /-S/mipt-golang-course/tasks/02/rle/rle.go
   22.30GB    22.30GB (flat, cum)   100% of Total
         .          .     31:func RLECompress(input string) string {
         .          .     32:   // return input + input
         .          .     33:   for _, c := range input {
   22.30GB    22.30GB     34:           input += string(c)
         .          .     35:   }
         .          .     36:
         .          .     37:   return input
         .          .     38:}
```
