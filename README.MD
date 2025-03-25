# Сервис сбора runtime-метрик

## О проекте

Сервер предназначен для сбора, хранения и получения runtime-метрик, отправляемых по HTTP-запросам. Поддерживаются метрики двух типов:
- **Gauge** (float64) — новое значение замещает предыдущее.
- **Counter** (int64) — новое значение добавляется к предыдущему.

## Запуск

Запуск сервера выполняется из корня проекта:

```sh
go run cmd/main.go
```

Сервер доступен по адресу:
```
http://localhost:8080
```

## Поддерживаемые эндпоинты

### Проверка доступности

- Запрос:
  ```bash
  curl http://localhost:8080/ping
  ```

- Ответ:
  ```
  Server is running
  ```

### Отправка метрик

Отправка gauge-метрики:
```bash
curl -X POST http://localhost:8080/update/gauge/temperature/24.5
```

Отправка counter-метрики:
```bash
curl -X POST http://localhost:8080/update/counter/requests/10
```

### Получение метрик

Получение gauge-метрики:
```bash
curl http://localhost:8080/value/gauge/temperature
```

Получение counter-метрики:
```bash
curl http://localhost:8080/value/counter/requests
```

## Структура проекта

```
yandex-practicum-metrics/
├── cmd/
│   └── main.go
├── internal/
│   ├── handlers/
│   │   └── handlers.go
│   ├── storage/
│   │   └── storage.go
│   ├── server/
│   │   └── server.go
│   └── database/
│       └── database.go
├── go.mod
└── go.sum
```

## Используемые технологии
- Go
- HTTP сервер из стандартной библиотеки Go

## Дальнейшее развитие
- Реализация постоянного хранения метрик в PostgreSQL
- Добавление логирования и мониторинга

