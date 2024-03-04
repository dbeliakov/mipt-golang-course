# Кодировщик/Декодировщик

Создайте тип CustomEncoder на основе стандартного типа string, который будет изменять каждый символ во входящей строке на его Unicode номер, используя конкретный разделитель. 

Реализуйте для CustomEncoder методы Encode(string) string для кодирования строки и Decode(string) string для обратного преобразования. Используйте константы для определения различных режимов кодирования (например, с разделителями "-", "_", итд.).

При решении задачи могут помочь функции стaндартной библиотеки: `strings.Split`, `strings.Join`, `strconv.Itoa`, `strconv.Atoi`.