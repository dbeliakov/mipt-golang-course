# JSON censor

JSON-сообщения в Go обычно представляются в виде слайса байт. 
Нужно пройти по всем **строковым значениям** в JSON и при совпадении с правилами цензора заменить **всё значение целиком** на `"***"` (не подстроку внутри JSON-строки — меняется вся строка как значение).

Поведение задаётся структурой `CensorConfig`:

- **`Needles`** — слайс запрещённых фрагментов (или целых строк — см. `Mode`). Пустые элементы игнорируются.
- **`Mode`**:
  - `MatchSubstring` — цензорить, если строка **содержит** любую из игл как подстроку.
  - `MatchWholeValue` — цензорить только если строка **целиком совпадает** с одной из игл.
- **`CaseInsensitive`** — если `true`, для `MatchSubstring` сравнение подстроки без учёта регистра (через приведение к нижнему регистру для ASCII); для `MatchWholeValue` — сравнение через `strings.EqualFold`.

Пример:

```go
input := []byte(`{"k": "v", "another_key": {"arr": ["word_to_censor"],
"inner_key": "what if i have a word_to_censor here???"}}`)

cfg := CensorConfig{
    Needles: []string{"word_to_censor"},
    Mode:    MatchSubstring,
}

res, _ := CensorJSON(input, cfg)
fmt.Println(string(res))
// {"k": "v", "another_key": {"arr": ["***"], "inner_key": "***"}}
```

Реализуйте `func CensorJSON(jsonData []byte, cfg CensorConfig) ([]byte, error)`.

Подобную задачу, возможно, можно решить с помощью регулярных выражений или ручного парсинга JSON-сообщения, но скорее всего это приведет к невероятно переусложненному коду. 
Для решения этой задачи можно вспомнить, что могут представлять собой JSON-значения:
- Числа - `int` и `double`
- Строки - `string`
- Массивы произвольных JSON-значений - `[]any`
- Наборы пар ключ-значение, где ключи являются строками - `map[string]any`

Для каждого из значений понятно, что с ним нужно делать — не трогать числа, обрабатывать строки по правилам выше, итерироваться по массивам и объектам рекурсивно.

Очень удобно, что интерфейсы в Go как раз обладают таким свойством (в частности `any`) с помощью [type switch](https://go.dev/tour/methods/16).

```go
var val any
err := json.Unmarshal(jsonMessage, &val)
```

Вам могут пригодиться:
- [json.Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal)
- [json.Marshal](https://pkg.go.dev/encoding/json#Marshal)
- [strings.Contains](https://pkg.go.dev/strings#Contains), [strings.EqualFold](https://pkg.go.dev/strings#EqualFold)
