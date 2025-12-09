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
RUN go build -o main .

# =======================
# Final Stage
# =======================
FROM alpine:3.19

WORKDIR /app

# Copiar binário da aplicação
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
