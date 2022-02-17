# Лекция 1. Введение

* [Слайды](https://dbeliakov.github.io/mipt-golang-course/lectures/01/lecture.slide.html)

* [Официальный сайт](https://golang.org/)
* [The Go Playground](https://play.golang.org/)
* [The Go Programming Language](https://www.gopl.io/) (есть на русском языке)

### Сдача заданий
1. Сделать приватный fork репозитория
2. Дать пользователю [dbeliakov](https://github.com/dbeliakov) доступ до репозитория
3. Прислать в чат Имя, Фамилию и ссылку на fork репозитория

### Приватный fork публичного репозитория

1. Склонировать репозиторий (с флагом `bare`)
```shell
git clone --bare git@github.com:dbeliakov/mipt-golang-course.git
```
2. Создать на GitHub приватный репозиторий
3. Сделать mirror-push склонированного репозитория в только что созданный приватный
```shell
cd mipt-golang-course.git
git push --mirror git@github.com:<your_username>/<your_repo>.git
```
4. Удалить склонированный репозиторий
```shell
cd ..
rm -rf mipt-golang-course.git
```
5. Склонировать приватный репозиторий
```shell
git clone git@github.com:<your_username>/<your_repo>.git
```
6. Добавить remote на оригинальный репозиторий, чтобы можно было получать обновления
```shell
git remote add upstream git@github.com:dbeliakov/mipt-golang-course.git
git remote set-url --push upstream DISABLE
```
7. Чтобы получить изменения из upstream, нужно выполнить
```shell
git fetch upstream
git rebase upstream/master
```

