## Обзор
Это репозиторий, где реализован код клиент-серверного приложения, генерирующего отчеты.

## Структура репозитория
Этот репозиторий состоит из следующих основных частей:

- [//build](/build) - файлы системы сборки и дополнения
- [//client](/client) - клиентская часть приложения (для возможности отправлять запросы)
- [//components](/components) - компоненты и приложения
- [//server](/server) - серверная часть приложения

## Стек
- docker-compose: v2.10.2
- go: v1.24

## Как запустить:
### Shell
```shell
  ./build.sh
```

### Вручную
```
go mod tidy
go run server/cmd/main.go
go run client/cmd/main.go
```