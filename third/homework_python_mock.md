## Теория
### Чистота функций
- https://devman.org/encyclopedia/clean_code/decomposition_pure_functions/
### Форматы файлов
- csv: https://pyneng.readthedocs.io/ru/latest/book/17_serialization/1_csv.html и https://docs.python.org/3/library/csv.html
- Что такое json? http://www.json.org/json-ru.html
- Как работать с json? https://pyneng.readthedocs.io/ru/latest/book/17_serialization/2_json.html и http://docs.python-guide.org/en/latest/scenarios/json/ и https://docs.python.org/3/library/json.html
### tcp/ip
- Стек TCP/IP: https://ru.wikipedia.org/wiki/TCP/IP
- Протокол HTTP: https://ru.wikipedia.org/wiki/HTTP
- Протокол TCP: https://ru.wikipedia.org/wiki/TCP
- Протокол UDP: https://ru.wikipedia.org/wiki/UDP
- Протокол IP: https://ru.wikipedia.org/wiki/IP
### urllib / requests
- Модуль urllib: https://docs.python.org/3/library/urllib.html
- requests http://docs.python-requests.org/en/master/
- requests vs urllib https://stackoverflow.com/questions/2018026/what-are-the-differences-between-the-urllib-urllib2-and-requests-module
### loguru
- loguru: https://github.com/Delgan/loguru
### mock
- https://pypi.org/project/pytest-mock/
- https://docs.python.org/3/library/unittest.mock.html
- http://blog.plataformatec.com.br/2015/10/mocks-and-explicit-contracts/
## Практика
Необходимо реализовать простой парсер данных и написать тесты к нему.
1. На вход подается csv список email адресов, вида:
```csv
"Sincere@april.biz","Shanna@melissa.tv","anastasia.net"
```
2. Получаем по переданным почтам id пользователей с сайта: https://jsonplaceholder.typicode.com/users/
3. Для каждого пользователя необходимо скачать всю информацию по его posts, albums и todos, пример:
- https://jsonplaceholder.typicode.com/users/1/posts
- https://jsonplaceholder.typicode.com/users/1/albums
- https://jsonplaceholder.typicode.com/users/1/todos
4. В качестве результата необходимо сохранить всю информацию о каждом пользователе в виде xml файла вида:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<user>
  <id>1</id>
  <email>Sincere@april.biz</email>
  <posts>
    <post>
      <id>1</id>
      <title>sunt aut facere repellat provident occaecati excepturi optio reprehenderit</title>
      <body>quia et suscipit suscipit recusandae consequuntur expedita et cum reprehenderit molestiae ut ut quas totam nostrum rerum est autem sunt rem eveniet architecto</body>
    </post>
  </posts>
  <albums>
    <album>
      <id>1</id>
      <title>omnis laborum odio</title>
    </album>
  </albums>
  <todos>
    <todo>
      <id>1</id>
      <title>delectus aut autem</title>
      <completed>false</completed>
    </todo>
  </todos>
</user>
```
Если у пользователя нет какой-то части данных, например posts, то записываем просто <posts></posts>
5. Все файлы мы сохраняем в папку users/${userId}/, где ${userId} заменяем на id пользователя
6. Необходимо логировать:
- количество пользователей для запуска
- переход к парсингу каждого конкретного email, пример: Starts parsing for Sincere@april.biz
- время выполнения каждого запроса к внешнему источнику данных
- логируем ошибки, если они случаются
7. Тесты должны покрывать удачные и провальные случаи. Для практики будем всегда полностью мокировать логгер (logguru), файловую систему (сохранение и загрузка файлов) и сеть (requests).
Технические требования:
- используем встроенный модуль csv для работы с csv
- Используем poetry для управления зависимостями
- Используем loguru для логирования
- Используем requests для выполнения запросов
- Можно выбрать любую библиотеку для работы с xml
- Тесты пишем на pytest
- Способ мока - выбирайте сами