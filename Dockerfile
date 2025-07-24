FROM golang:1.22-alpine

WORKDIR /app

# Copia os arquivos do projeto
COPY . .

# Inicializa o módulo (ignora erro se já existir)
RUN [ -f go.mod ] || go mod init bankdb

# Instala dependências (idempotente)
RUN go get github.com/lib/pq
RUN go mod tidy

# Compila
RUN go build -o app main.go crudbanco.go

CMD ["./app"]
