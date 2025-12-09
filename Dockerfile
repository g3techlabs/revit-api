# =======================
# Build Stage
# =======================
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Instalar dependências necessárias (git, gcc)
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilar aplicação
RUN go build -o build/main -ldflags "-s -w" ./src

# =======================
# Final Stage
# =======================
FROM alpine:latest

WORKDIR /app

# Copiar binário da aplicação
COPY --from=builder /app/build/main .

EXPOSE 8080

CMD ["./main"]
