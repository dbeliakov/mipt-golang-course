# Асинхронный загрузчик

В данной задаче вам нужно написать асинхронный загрузчик произвольных `URL`. 

Загрузчик принимает на вход произвольное количество произвольных `URL` и возвращает хэш-таблицу с контентом по ключу `URL` `map[string][]byte`.

Помимо таблицы с контентом для страниц, скачать которые удалось успешно, загрузчик должен возвращать слайс ошибок `[]error` для каждой страницы, которую скачать не удалось. 
Все ошибки в слайсе должны иметь тип `DownloadError` и иметь заполненные поля `url` и `err`.

## http.Response

Также в этой задаче вам предстоит поработать с [http.Response](https://pkg.go.dev/net/http#Response) -- результатом выполнения http-запроса.
Прочитать из него тело ответа можно с помощью `Response.Body`, реализующего [io.ReadCloser](https://pkg.go.dev/io#ReadCloser).

При работе с `http.Response` нужно помнить про пару нюансов:
- тело запроса `Response.Body` обязательно нужно закрывать, иначе соединение не вернется в пул и данные запроса не освободятся
- тело запроса `Response.Body` нужно всегда читать до конца (даже когда само тело не нужно) по тем же причинам
  
Пример:

```go
client := http.DefaultClient
resp, err := client.Get("https://google.com")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

data, err := io.ReadAll(resp.Body)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))
```

Вам могут пригодиться:
- [sync.WaitGroup](https://pkg.go.dev/sync#WaitGroup)
- [io.ReadAll](https://pkg.go.dev/io#ReadAll)
