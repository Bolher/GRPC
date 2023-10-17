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

## Execução

Para executar o servidor, siga estas etapas:

Execute o servidor Go:

```bash
go run server.go

## Cliente (Python)

O cliente gRPC permite que você envie solicitações de validação de documentos para o servidor.

### Configuração

Para configurar o cliente Python, siga estas etapas:

Instale as dependências do Python:

```bash
pip install grpcio grpcio-tools

### Execução

Para executar o cliente, siga estas etapas:

```bash
python client.py

