# Colorize

Одним из первых способов запечатлеть цветную фотографию, применненный, например, для создания [цветных фотографий царской России](https://moscow.zagranitsa.com/article/2256/pervye-tsvetnye-fotografii-tsarskoi-rossii-prokudi), состоял в следующем: через разные цветофильтры (например, красный, зеденый и синий) создавалось три изображения, которые в последующем проецировались вместе с использованием тех же цветофильтров.

В данном задании Вам необходимо написать реализации функций, которые для цивровых изображений будут делать подобные процедуры: раскладывать цветное изображение на три черно-белых (по компонентам RGB), и обратно - собирать из трех черно-белых изображений цветное.

Вам необходимо реализовать функции `DecomposeGRB` и `ComposeGRB`. Сигнатуры функций оперируют объектами `image.Image`, которые чвляются интерфейсами (то есть скрывают фактический тип), однако Вам надо использовать (и можно рассчитывать, что это верно и для входных данных) типы [image.Gray](https://pkg.go.dev/image#Gray) и [image.RGBA](https://pkg.go.dev/image#RGBA). Альфа канал во всех примерах считаем равным максимальному значению (т.е. изображение не прозрачное).

Для помощи в отладке, в директории `cmd/02/colorize` Вы можете найти утилиту, которая вызывает соответствующие функции и выполняет соответствующие операции, сохраняя файлы на диск.

Задача считается сделанной, если тесты в директории (`go test`, для более полного вывода - флаг `-v`) проходят.