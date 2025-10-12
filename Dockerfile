FROM golang:1.25.1-alpine AS builder

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
