# API Retries

При работе с каким-нибудь API достаточно часто возникают ошибки, которые зачастую надо ещё и ретраить.

Допустим, у нас есть простой интерфейс key-value кэша:
```go
type SimpleAPI interface {
	Get(key string) (val *Value, epoch uint64, err error)
	Set(key string, targetEpoch uint64, value *Value) error
}
```

Вам нужно написать функцию, которая принимает `context.Context`, реализацию интерфейса, ключ, функцию обновления значения и `UpdateOptions`. Функция должна по ключу получить значение из API, обновить его и сохранить. В процессе могут возникать разные виды ошибок:

- таймауты и ошибки сети (`ErrTimeout`, `ErrNetworkFault`) нужно ретраить
- если значение по ключу не нашлось (`APIError.Status() == StatusNotFound`) или произошла фатальная ошибка на стороне API (`APIError.Status() == StatusFatalError`), ошибку нужно вернуть из функции
- если кто-то уже успел обновить значение быстрее нас (`APIError.Status() == StatusValueTooOld`), нужно снова попробовать получить значение по ключу и обновлять уже новое значение (это тоже считается повтором попытки)

**Ограничение на повторы:** цикл не должен быть бесконечным. Используйте:

- отмену через `ctx.Done()` — если контекст отменён, верните `ctx.Err()`;
- **лимит повторов** в `UpdateOptions.MaxRetries`: максимальное число раз, когда вы делаете `continue` из-за сети или `StatusValueTooOld`. Если лимит исчерпан, верните ошибку `ErrRetryBudgetExhausted`. Нулевое значение `MaxRetries` можно трактовать как разумное значение по умолчанию (большое, но конечное).

**Колбэк:** поле `UpdateOptions.OnRetry` — функция `func(attempt int, err error)`, вызывается **перед** очередным повтором; `attempt` начинается с 1, `err` — ошибка, из-за которой выполняется retry.

Сигнатура:

```go
func UpdateValue(ctx context.Context, api SimpleAPI, key string, update ValueUpdater, opts UpdateOptions) error
```

Вам пригодятся:
- [context.Context](https://pkg.go.dev/context#Context)
- [errors.Is](https://pkg.go.dev/errors#Is)
- [errors.As](https://pkg.go.dev/errors#As)
