# Protocolo de Validação de Documentos com gRPC - Documentação

Este é um exemplo de um servidor e cliente que implementam um protocolo de validação de documentos usando gRPC. O servidor valida números de CPF ou CNPJ informados pelo cliente e responde com um código de validação. O servidor é escrito em Go e o cliente em Python, com a comunicação feita através do gRPC.

## O que é gRPC

[gRPC](https://grpc.io/) é um framework de código aberto que permite a comunicação eficiente e confiável entre aplicativos distribuídos usando um protocolo de serialização eficiente chamado Protocol Buffers (protobuf).

## Pré-requisitos

Para executar o código do servidor e do cliente gRPC, você precisa das seguintes dependências:

### Para o Servidor (Go)

- Go (Golang)
- Biblioteca gRPC para Go

### Para o Cliente (Python)

- Python 3
- Bibliotecas gRPC para Python

## Servidor (Go)

O servidor gRPC é configurado para ouvir na porta 50051.

### Configuração

Para configurar o servidor, siga estas etapas:

1. Instale as dependências do Go:

```bash
go get google.golang.org/grpc
```
## Execução

Para executar o servidor, siga estas etapas:

Execute o servidor Go:

```bash
go run server.go
```
## Cliente (Python)

O cliente gRPC permite que você envie solicitações de validação de documentos para o servidor.

### Configuração

Para configurar o cliente Python, siga estas etapas:

Instale as dependências do Python:

```bash
pip install grpcio grpcio-tools
```
### Execução

Para executar o cliente, siga estas etapas:

```bash
python client.py
```

# Como Funciona o Código

## Servidor (Go)

Quando o cliente envia uma solicitação, o servidor faz o seguinte:

Recebe a solicitação com um número de documento e um formato (0 para CPF, 1 para CNPJ).

Verifica o tamanho do número de documento para determinar se é um CPF (11 dígitos) ou um CNPJ (14 dígitos).

Com base no formato, o servidor decide qual algoritmo de validação aplicar.

Para CPF, ele usa a função isValidCPF para verificar se o número é válido.

Para CNPJ, ele usa a função isValidCNPJ para verificar se o número é válido.

Com base no resultado da validação, o servidor responde com um código de validação (0 para documento válido, 1 para documento inválido, 2 para tamanho de documento inválido, 3 para formato de documento inválido).

## Cliente (Python)

Ao iniciar o cliente, o seguinte acontece:

O cliente solicita que o usuário insira um número de documento.

O cliente também solicita que o usuário escolha o formato (0 para CPF, 1 para CNPJ).

Com base na escolha do usuário, o cliente envia a solicitação gRPC apropriada para o servidor.

O cliente recebe a resposta do servidor, que contém o código de validação.

O cliente exibe o código de validação no console.

Essa abordagem permite que o cliente interaja com o servidor e receba uma resposta que indica se o documento é válido e qual é o código de validação associado.

