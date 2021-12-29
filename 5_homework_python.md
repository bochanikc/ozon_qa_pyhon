## Домашнее задание к лекции "Внутренности Python"

Версия Python: 3.8

### 1. Реализовать декоратор с параметром.

Декоратор должен проверять наличие прав у пользователя.
В случае авторизации пользователя с недостаточными правами - должен вызываться PermissionError, который требуется определить самостоятельно.
В решении следует использовать следующий фрагмент кода:

```python
permitted_role = ['admin']

def check_access():
    # your code is here
    pass


@check_access('admin')
def user_login():
    return "Добро пожаловать!"

print(user_login())
# Результат работы: Добро пожаловать!


@check_access('user')
def user_login():
    return "Добро пожаловать!"

print(user_login())
# Результат работы: PermissionError: Access denied for user "user"
```


### 2. Разработать скрипт.

* Скрипт должен обращаться к API superheroapi.com (ручка */id*.).
* В качестве ID должно использоваться случайное значение (1-700)
* Из полученного json необходимо получить изображение (*image/url*) и сохранить на диск.

Таким образом, при запуске скрипта должен запускаться генератор, который вернет случайный ID, с которым необходимо отправить запрос в API.
После завершения работы скрипта в директории должна быть сохраненная картинка.

####Пример:

Результат работы генератора: ```ID: 720```

Запрос в API: 
```https://superheroapi.com/api/<access-token>/720```

```json
{
    "response": "success",
    "id": "720",
    "name": "Wonder Woman",
    "powerstats": {
        "intelligence": "88",
        "strength": "100",
        "speed": "79",
        "durability": "100",
        "power": "100",
        "combat": "100"
    },
    "image": {
        "url": "https://www.superherodb.com/pictures2/portraits/10/100/807.jpg"
    }
    "biography": {
        "full-name": "Diana Prince",
        "alter-egos": "No alter egos found.",
    ...
}
```

На диск должно быть сохранено изображение:
```https://www.superherodb.com/pictures2/portraits/10/100/807.jpg```

PS. Для авторизации понадобится логин через FB-аккаунт, если у вас его нет - решим в индивидуальном порядке.

### 2.1. Дополнительное задание.
* Необходимо сохранять картинку под названием, которое находится в полученном json в поле *biography/full-name*.
(если значение содержит пробелы - их нужно заменить на символ "_")
  
#### Пример:
На диск должно быть сохранено изображение по ссылке:
```https://www.superherodb.com/pictures2/portraits/10/100/807.jpg```

С названием:
```Diana_Prince.jpg```

### 3. Разработать генератор товаров.

Требуется разработать скрипт, который будет возвращать json-файл с N-товаров.
(N - рандомное число, которое скрипт принимает на вход)

Правила генерации для полей json выглядят следующим образом:
* **id** - номер по порядку.

* **article** - рандомная строка формата ```OZON_<category>_<A-Z0-9>```

* **category** - одно из значений:
```
outerwear - верхняя одежда
shirt     - рубашка
socks     - носки
```

* **weight** - вес товара, в граммах (integer). Выставляется по следующей логике:
```
outerwear: 1000 - 10000
shirt    : 100 - 2000
socks    : 10 - 300
```

* **type** - тип товара. Заполняется как:
```
outerwear: один из - coat (пальто), jacket (куртка) или windbreaker (ветровка)
shirt    : один из - casual (повседневная), dress (парадная) или loosefit (свободного кроя)
socks    : не заполняется (поля не должно быть в json)
```

* **colour** - рандомный цвет. Без ограничений.

* **textile** - состав (массив). Состоит из полей:
```
1. material: один из - cotton, silk, polyester, viscose, acrylic.
2. percent : процент содержания материала в составе товара
```

* **description** - рандомный текст. Уникальный для разных товаров.


Для генерации случайных данных **рекомендуется пользоваться библиотекой ```Faker```**.


Примерный вариант ответа может выглядеть следующим образом:

```Для N=2:```

```json
[{
    "id": 1,
    "article": "OZON_shirt_AB03",
    "category": "shirt",
    "weight": 500,
    "type": "casual",
    "colour": "LightSeaGreen",
    "textile": [
        {
            "material": "cotton",
            "percent": 60
        },
        {
            "material": "acrylic",
            "percent": 10
        },
        {
            "material": "viscose",
            "percent": 30
        }
    ],
    "description": "Give side human pressure. Social hear country own marriage. Technology event capital body language about loss."
},
{
    "id": 2,
    "article": "OZON_socks_KS24",
    "category": "socks",
    "weight": 250,
    "colour": "Gray",
    "textile": [
        {
            "material": "cotton",
            "percent": 90
        },
        {
            "material": "polyester",
            "percent": 10
        }
    ],
    "description": "Sit action per. Though century total old allow while middle enjoy. Myself body law order toward whatever."
}]
```

### 3.1. Дополнительное задание.

Требуется записать полученный массив с товарами в excel-файл.
Пример excel-файла можно посмотреть в файле ```example.xlsx```

Перед записью в файл требуется привести массив ```textile``` к строке:

```
60% cotton, 30% viscose, 10% acrylic
90% cotton, 10% polyester
```

При этом материал **должен быть отсортирован** в порядке убывания процента содержания (первый - максимальный)!

Для работы с форматом excel **рекомендуется пользоваться библиотекой ```openpyxl```**.