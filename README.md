## Обновить версию пакета протокола 

```
go get -u github.com/wrs-news/golang-proto@v0.0.1
```

## Создать миграцию

```
migrate create -ext sql -dir migrations <migration_name>
```

## Выполнить/Откатить миграцию

```
migrate -path migrations -database <url> [up/down]
```