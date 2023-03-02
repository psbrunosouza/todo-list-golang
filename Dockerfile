# Imagem base do Golang
FROM golang:latest

# Criação do diretório do workspace e definição do mesmo como diretório de trabalho
RUN mkdir -p /app
WORKDIR /app

# Copia todos os arquivos e pastas do projeto para dentro do workspace
COPY . .

# Criação do diretório bin
RUN mkdir -p /app/bin

# Executa o comando "go mod tidy" para garantir que todas as dependências estejam corretamente instaladas
RUN go mod tidy

# Compila o projeto e cria o binário
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/todoapp ./...

# Define o comando padrão a ser executado ao rodar a imagem
CMD ["./bin/todoapp"]