#### Домашнее задание по лекции "Gitlab CI"


#### 1. Линтеры.
Добавить проверку кода тестов из репозитория в существующий пайплайн с помощью линтера.

* Линтер: ```flake8```
* Тесты лежат в ```framework/tests```

* Линтер может быть сконфигурирован следующим образом:
```sh
--max-line-length=120 --ignore=W605,W504
```
* Стадия: ```check``` (должна выполняться в самом начале)
* Список обнаруженных ошибок линтера нужно сохранить в файл, который будет доступен для скачивания в течение семи дней.
* Джоба должна запускаться автоматически при каждом запуске пайплайна.


