# Pi

> :warning: **Для получения всех баллов за задачу необходимо пройти Code Review**.
> Для этого необходимо сделать коммит с решением в отдельную ветку, сделать из этой ветки Pull Request в master и добавить dbeliakov в этот PR.

Необходимо реализовать функцию, которая [подсчитывает число pi методом Монте Карло](https://habr.com/ru/post/128454/).
Следующую случайную точку необходимо брать из генератора с интерфейсом `RandomPointGenerator`.
Вычисление должно происходить в `concurrent` горутинах, **общее число** итераций должно быть равно
`iterations`.