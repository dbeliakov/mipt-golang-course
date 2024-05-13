# Файловая система с вложенностью

Вам необходимо реализовать модель файловой системы с поддержкой вложенности

Требования:
1. Директория может включать в себя файлы и другие директории. При этом, если существует файл или директория по тому же пути, то добавление должно перезаписывать существующее.
2. Необходимо поддержать возможность создания файлов и директорий (и получения их) с указанием полного пути, разделенного '/'.
3. При добавлении файла или директории, все промежуточные директории должны быть созданы при необходимости.
4. У директории должна быть возможность посчитать, сколько весят вложенные в нее файлы (с учетом тех, которые лежат во вложенных директориях).