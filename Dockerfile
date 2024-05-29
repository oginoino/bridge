# Use uma imagem base do Go
FROM golang:1.22.3

# Defina o diretório de trabalho dentro do container
WORKDIR /app

# Copie os arquivos do projeto para o diretório de trabalho
COPY . .

# Baixe as dependências do Go e compile a aplicação
RUN go mod download
RUN go build -o bridge ./main.go

# Execute a aplicação
CMD ["./bridge"]
