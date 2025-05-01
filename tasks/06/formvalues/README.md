# Form Values Unpacker

Вам необходимо реализовать декодированиe данных HTML-форм (form values) в структуры Go.

Реализуйте функцию:
```go
func Unpack(req *http.Request, ptr interface{}) error
```

Которая должна:
- Парсить form values из HTTP-запроса
- Заполнять переданную структуру (ptr) соответствующими значениями
- Поддерживать базовые типы: string, int, bool, []string
- Использовать тег  для указания имен полей формы
- Возвращать ошибки при:
  - Несоответствии типов
  - Передаче не-указателя
  - Неподдерживаемых типах полей
  
Вам могут пригодиться:
- [http.Request](https://pkg.go.dev/net/http#Request)
- [http.Request.ParseForm](https://pkg.go.dev/net/http#Request.ParseForm)
- [reflect.ValueOf](https://pkg.go.dev/reflect#ValueOf)
- [reflect.Value](https://pkg.go.dev/reflect#Value)
- [reflect.Type](https://pkg.go.dev/reflect#Type)
- [strconv](https://pkg.go.dev/strconv)
