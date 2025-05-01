# Mock Helper

Для подмены поведения произвольных интерфейсов часто используют моки, сгенерированные автоматически.
В этой задаче вместо генератора моков напишем библиотеку, которая поможет писать их вручную. Библиотека должна позволять:

1. Задавать ожидаемые вызовы методов с конкретными аргументами
2. Возвращать заданные значения для этих вызовов
3. Проверять, что все ожидаемые вызовы были выполнены
4. Поддерживать подачу произвольных значений (если мы не уверены, с каким аргументом будет вызван метод, или если это не важно)

## Пример использования

```go
func TestUserService(t *testing.T) {
    mh := mockhelper.NewMockHelper(t)

    // Необязательно, должен быть вызван в любом случае!
    defer mh.Verify()
    
    mh.ExpectCall("GetUser", 123).Return(&User{Name: "Alice"}, nil)
    
    // Ожидаем вызов DeleteUser с любым аргументом
    mh.ExpectCall("DeleteUser", mockhelper.Any()).Return(nil)
    
    service := NewUserService(&MockDB{mh: mh})
    
    user, err := service.GetUser(123)
    require.NoError(t, err)
    
    err = service.DeleteUser(456)
    require.NoError(t, err)
}
```

Больше примеров можно найти в тестах.


Для проверки того, что все ожидаемые вызовы были совершены, можно запускать `defer mh.Verify()` после каждого создания нового хелпера, но об этом можно забыть. 
Чтобы избавиться от надобности вызывать defer в тестах, можно воспользоваться методом `Cleanup()` у `*testing.T`, который сам вызовет нужную функцию перед завершением теста.

Если во время теста был вызван неожиданный метод, нужно сразу завершать этот тест, так как мы не знаем, что должен возвращать метод, который мы не ожидали увидеть.

Вам могут пригодиться:
- [reflect.DeepEqual](https://pkg.go.dev/reflect#DeepEqual)
- [testing.T.Cleanup](https://pkg.go.dev/testing#T.Cleanup)
- [testing.T.Errorf](https://pkg.go.dev/testing#T.Errorf)
- [testing.T.FailNow](https://pkg.go.dev/testing#T.FailNow)
