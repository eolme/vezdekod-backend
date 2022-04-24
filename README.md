# BackMemes

## Запуск

### В консольном режиме
```bash
go run . --mode=cli
```

### В режиме сервера
```bash
go run . --mode=api
```

- `/` - пользовательский интерфейс
- `/admin` - панель администрации

## Бинарники

```bash
./main-linux

./main-darwin
```

## Seed

```bash
TOKEN=<vk_api_access_token> go run .
```

> Процесс очень долгий, не рекомендуется, база уже засижена
