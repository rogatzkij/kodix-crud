# Тестовое задание для Kodix
## Задание
Разработать CRUD (REST API) для модели автомобиля, который имеет следующие поля:
1. Уникальный идентификатор (любой тип, общение с БД не является критерием чего-либо, можно сделать и in-memory хранилище на время жизни сервиса)
2. Бренд автомобиля (текст)
3. Модель автомобиля (текст)
4. Цена автомобиля (целое, не может быть меньше 0)
5. Статус автомобиля (В пути, На складе, Продан, Снят с продажи)
6. Пробег (целое)
Формат ответа api - json api (https://jsonapi.org/) 

## О реализации
* при составлении структуры проекта опирался на [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

## Проверка

### Создание бренда

~~~shell script
curl --location --request POST '127.0.0.1:8080/api/v1/brands/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "brandname": "kia"
}'
~~~

### Удаление бренда
Внимание: При удалении бренда не будут удалены машины этого бренда.
~~~shell script
curl --location --request DELETE '127.0.0.1:8080/api/v1/brands/kia'
~~~
