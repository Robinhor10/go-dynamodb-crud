FROM golang:1.20-alpine

WORKDIR /app

# Copiar o código fonte
COPY . .

# Inicializar o módulo Go e instalar dependências
RUN go mod init myapp
RUN go mod tidy

# Compilar a aplicação
RUN go build -o /myapp

EXPOSE 8080

CMD ["/myapp"]
