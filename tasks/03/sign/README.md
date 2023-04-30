# Sign

Вам необходимо реализовать механизм подписи данных на основе двух разных алгоритмов: RSA и HMAC.

В файле [signer.go](signer.go) находятся интерфейсы Signer и Validator, имплементации которых находятся рядом в соответствующих файлах. Обе реализации принимают в качестве параметров метод хэширования, который необходимо использовать при рассчете подписи (для примеров использования см. тесты).

В случае HMAC создать объект для вычисления подписи можно с помощью [hmac.New](https://pkg.go.dev/crypto/hmac#New). В случае RSA для вычисления и фалидации подписи используйте [rsa.SignPSS](https://pkg.go.dev/crypto/rsa#SignPSS) и [rsa.VerifyPSS](https://pkg.go.dev/crypto/rsa#VerifyPSS).

В случае возникновения ошибок необходимо их вернуть наружу. В пакете также объявлена ошибка `ErrInvalidMethod` - ее необходимо возвращать в случае, когда на валидацию передана структура с указанным другим методом.