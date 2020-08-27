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
* в сервисе встречаются странные на первый взгляд названия полей `brandname` (т.к. часто используется переменная `brand`) и `automodel` (т.к. `model` совпадает с названием одного из пакетов)
* взаимодействие с БД осуществляется с помощью объектов, которые реализует интерфейсы. Если в будущем не планируется переход на другую БД, то интерфейсы в данном случае могут быть избыточными.

### Структура проекта
~~~
.
├── cmd
│   └── main.go         - точка входа 
├── config
│   └── config.go       - конфиги для запуска (переменные окружения)
├── internal
│   ├── contract        - пакет с интерфейсами по работе с данными
│   │   ├── auto.go  
│   │   └── brand.go  
│   ├── core            - пакет с логикой работы с данными
│   │   └── core.go     
│   ├── handler         - http ручки и миделвары
│   │   └── handler.go  
│   └── mongo           - пакет, реализующий взаимодействие с БД
│       └── mongo.go    
├── model
│   ├── auto.go     - структура записи автомобиля
│   ├── brand.go    - структура записи бренда
│   └── error.go    - возможные ошибки
├── go.mod
├── go.sum
└── README.md       - ты здесь :)
~~~
## О работе сервиса

Перед добавлением записи автомобиля в коллекции с брендом должна находится запись содержащая информацию о модели.

Пример
~~~JSON
  {
    "_id": "5f46d2883336e5f091656374",
    "brandname": "lada",
    "automodels": ["kalina", "niva"]
  }
~~~

### Создание бренда

~~~shell script
curl --location --request POST '127.0.0.1:8080/api/v1/brands/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "brandname": "lada"
}'
~~~

При создании можно сразу указать модели бренда в поле `brandname`.

### Удаление бренда
Внимание: При удалении бренда не будут удалены машины этого бренда.
~~~shell script
curl --location --request DELETE '127.0.0.1:8080/api/v1/brands/kia'
~~~


### Создание марки
На момент создания марки бренд должен уже существовать.
~~~shell script
curl --location --request POST '127.0.0.1:8080/api/v1/brands/lada/models/niva'
~~~

### Удаление марки
~~~shell script
curl --location --request DELETE '127.0.0.1:8080/api/v1/brands/lada/models/niva'
~~~

### Получение записи об одном автомобиле
~~~shell script
curl --location --request GET '127.0.0.1:8080/api/v1/autos/4'
~~~

### Получение записи о нескольких автомобилях
Для пагинации в запросе указывается GET-параметры `limit` (по умолчанию 10) и `offset` (по умолчанию 0).

~~~shell script
curl --location --request GET '127.0.0.1:8080/api/v1/autos/?offset=0&limit=3' 
~~~
### Создание автомобиля
~~~shell script
curl --location --request POST '127.0.0.1:8080/api/v1/autos/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "brandname": "lada", 
    "automodel": "niva",
    "price": 1000,
    "status": "stock",
    "mileage": 5000
}'
~~~
### Изменение автомобиля
При изменении должны быть переданы все поля.
~~~shell script
curl --location --request PUT '127.0.0.1:8080/api/v1/autos/4' \
--header 'Content-Type: application/json' \
--data-raw '{
    "brandname": "lada", 
    "automodel": "niva",
    "price": 1000,
    "status": "sold out",
    "mileage": 7000
}'
~~~
### Удаление автомобиля
~~~shell script
curl --location --request DELETE '127.0.0.1:8080/api/v1/autos/4'
~~~

## Что можно доработать
* Сделать более точное соответствие `jsonapi` с указанием ссылок и метаинформации 
* Проставить индексы в БД
* Выдавать в API сообщения об ошибках, не содержащие сведенья о системе
* Добавить автотестов с моками монги
* Сделать авторизацию (хотя бы `basic`) и роли пользователям
* Собирать docker контейнер  (3 стадии: prebuild/build/service)