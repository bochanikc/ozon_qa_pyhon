# Домашнее задание

## Проверить proto файл с помощью линтера protolint

### Подготовка:

Установить `protolint`

```bash
brew tap yoheimuta/protolint
brew install protolint
```

### Задание:

- Проверить линтером `echo.proto` файл
- Исправить 7 ошибок в `echo.proto` файле

### Критерии приемки задания:

- В git залит исправленный `echo.proto` файл
- Все ошибки исправлены

### Подсказка:

- Синтаксис файла `echo.proto` ДОЛЖЕН быть`proto3`

### Материалы:

- [Style Guide](https://developers.google.com/protocol-buffers/docs/style) для proto файлов
- [protolint](https://github.com/yoheimuta/protolint)

## Проверить proto файл с помощью линтера protolint

### Подготовка:

- Установить `grpcurl`

```bash
brew install grpcurl
```

- Развернуть учебный [пример](https://grpc.io/docs/platforms/web/quickstart/):

```bash
git clone https://github.com/grpc/grpc-web
cd grpc-web
docker-compose pull prereqs node-server envoy commonjs-client
docker-compose up -d node-server envoy commonjs-client
```

### Задание:

Выполнить команды `grpcurl`, чтобы:

- вывести список `list` методов сервиса
- вывести описание `describe` метода, запроса и ответа
- выполнить запрос метода `Echo`
- выполнить запрос метода `ServerStreamingEcho`

Сохранить лог выполненных вызовов из консоли в файл.

### Критерии приемки задания:

- В лог файле есть вызовы `list`, `describe`, `Echo`, `ServerStreamingEcho`

### Подсказка:

```bash
grpcurl -import-path ~/Downloads -proto echo.proto -plaintext localhost:9090 list
```

### Материалы:

- [grpcurl](https://github.com/fullstorydev/grpcurl)
