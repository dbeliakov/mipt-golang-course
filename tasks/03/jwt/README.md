# JWT

> :warning: **Для получения всех баллов за задачу необходимо пройти Code Review**.
> Для этого необходимо сделать коммит с решением в отдельную ветку, сделать из этой ветки Pull Request в master и добавить dbeliakov в этот PR.

Необходимо реализовать мини-пакет для работы с Json Web Tokens (JWT). Простое объяснение JWT можно
почитать в [Пять простых шагов для понимания JSON Web Tokens (JWT)](https://habr.com/ru/post/340146/).

Все экспортируемые сущности пакета уже объявлены, новые типы, функции и пр. должны быть неэкспортируемыми. Запрещено использование любых
пакетов кроме стандартной библиотеки.

Для настройки параметров работы функции используются специальные функции (см. [opts.go](./opts.go)).

> Обратите внимание на переменную timeFunc. Получение текущего времени нужно делать через нее,
> иначе в тестах не удастся переопределить время

> При решении задачи постарайтесь минимизировать операции конвертации `[]byte` в
> `string` и обратно (они копируют контент) и операции конкатенации строк (используйте `bytes.Buffer`)

> Выносите повторяющиеся куски кода в отдельные функции, константам давайте осмысленные имена, старайтесь сделать код поддерживаемым

> В процессе решения удобно для отладки токенов использовать [jwt.io](https://jwt.io/)

Функция `Encode` принимает на вход пользовательские данные и опции для конфигурирования токена,
возвращает сам токен или ошибку. Список ошибок:
* `ErrConfigurationMalformed` - конфигурация противоречива (например, заданы оба `TTL` и `Expires`, либо `Expires` меньше, чем `Now`)
* `ErrInvalidSignMethod` - переданный метод не попадает ни под одну из перечисленных констант для метода подписи

Функция `Decode` принимает на вход токен, пользовательские данные для заполнения и опции для конфигурирования,
возвращает ошибку или nil в случае успеха. Список ошибок:
* `ErrInvalidToken` - переданный токен не удовлетворяет формату `JWT` (количество компонент, невалидный `base64` или `json`, ...)
* `ErrSignMethodMismatched` - метод подписи в конфигурации и в токене не соответствуют друг другу
* `ErrInvalidSignMethod` - переданный метод не попадает ни под одну из перечисленных констант для метода подписи
* `ErrSignatureInvalid` - контент токена не соответствует его подписи
* `ErrTokenExpired` - время жизни токена закончилось

#### Пакеты, которые могут пригодиться при решении задачи
* `bytes`
* `strings`
* `crypto/hmac`
* `crypto/sha256`
* `crypto/sha512`
* `encoding/base64`
* `encoding/json`
* `hash`
* `time`

#### Функции, которые могут пригодиться при решении задачи
* `json.Marshal`, `json.Unmarshal`
* методы `Add`, `Before`, `After`, `Unix` у объекта типа `time.Time`
* `time.Unix`
* `strings.Split`, `bytes.Split`
* для `base64` кодирования используйте `base64.RawURLEncoding`
* для создания объекта для вычисления подписи используйте `hmac.New(sha***.New, key)`