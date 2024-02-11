# Stage 1: Build the Go application
FROM golang:1.21-alpine as builder

WORKDIR /app

# Копируем go.mod и go.sum из корня проекта
COPY ../../go.mod ../../go.sum ./

# Загрузка зависимостей
RUN go mod download

# Копируем весь проект внутрь контейнера
COPY ../../ ./

# Сборка приложения
RUN go build -o ./bin/chat-api ./cmd/chat-api/main.go

# Stage 2: Create a minimal runtime container
FROM alpine

WORKDIR /app

# Копируем собранный бинарный файл из предыдущего этапа
COPY --from=builder /app/bin/chat-api .

# Указываем команду для запуска приложения
CMD ["./chat-api"]