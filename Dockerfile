# =======================
# Build Stage
# =======================
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Instalar dependências necessárias (git, gcc)
RUN apk add --no-cache git

# Instalar o Swag CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Gerar documentação Swagger dentro do container
RUN swag init -g src/main.go

# Compilar aplicação
RUN go build -o main .

# =======================
# Final Stage
# =======================
FROM alpine:3.19

WORKDIR /app

# Copiar binário da aplicação
COPY --from=builder /app/main .
# Copiar docs do Swagger
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./main"]
