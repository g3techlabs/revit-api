FROM golang:1.25.1-alpine AS builder

# Instala o Air e dependências
RUN apk add --no-cache git bash && \
    go install github.com/air-verse/air@latest

WORKDIR /app

# Copia apenas os arquivos necessários primeiro (para cache eficiente)
COPY go.mod go.sum ./
RUN go mod download

# Copia o resto do código
COPY . .

# Expor a porta padrão do Fiber
EXPOSE 8080

# Comando de inicialização usando Air
CMD ["air", "-c", ".air.toml"]
