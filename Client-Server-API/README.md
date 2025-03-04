# Desafio Go - Cotação do Dólar

Este repositório contém a implementação de dois sistemas em Go: `client.go` e `server.go`. O objetivo deste desafio é criar uma comunicação HTTP entre o cliente e o servidor, onde o cliente solicita a cotação do dólar e o servidor consulta uma API de câmbio, armazena a cotação no banco de dados SQLite e retorna o valor da cotação para o cliente.

## Requisitos

### Arquivos

- **client.go**: O cliente realiza uma requisição HTTP no servidor (`server.go`) solicitando a cotação do dólar.
- **server.go**: O servidor consome a API de câmbio do dólar e retorna o valor da cotação no formato JSON para o cliente.

### Funcionalidades

1. **Requisição HTTP**:
   - O cliente realiza uma requisição para o servidor na URL `/cotacao`.
   - O servidor consulta a API de câmbio no endpoint `https://economia.awesomeapi.com.br/json/last/USD-BRL`.
   - O servidor armazena as cotações recebidas no banco de dados SQLite.
   
2. **Timeouts e Contextos**:
   - O `server.go` deve usar o package `context` para:
     - **Timeout da requisição HTTP**: Máximo de 200ms para chamar a API de cotação do dólar.
     - **Timeout de Persistência**: Máximo de 10ms para persistir os dados no banco SQLite.
   - O `client.go` deve:
     - Usar o package `context` para ter um timeout máximo de 300ms para receber o resultado do servidor.
   - Caso qualquer um dos contextos ultrapasse o tempo limite, deve ser gerado um erro no log.

3. **Armazenamento da Cotação**:
   - O servidor deve salvar cada cotação recebida no banco de dados SQLite.
   - O cliente deve salvar a cotação atual em um arquivo `cotacao.txt` com o formato:
     ```
     Cotação do dólar (dd/MM/yyyy HH:mm:ss): {valor}
     ```

4. **Estrutura do Projeto**:
   - **Servidor HTTP**:
     - O servidor escuta na porta `8080`.
     - O endpoint do servidor é `/cotacao`, que retorna o valor atual do câmbio do dólar no formato JSON.
   - **Banco de Dados SQLite**:
     - O servidor deve usar SQLite para persistir os dados da cotação do dólar.
     - A tabela no banco de dados deve ser criada com as colunas adequadas para armazenar o valor da cotação e a data/hora da consulta (tabela `cotacao`).

### Endpoints

- **GET /cotacao**: Retorna o valor atual do câmbio de Dólar para Real (campo `bid` do JSON da API) no formato JSON.

### Exemplo de Resposta do Endpoint

```json
{
  "bid": "5.4245"
}
```

### Usabilidade

- Rodar o arquivo `main.go` da pasta **server** e esperar a **API** subir na porta `8080`;
- Rodar o arquivo `client.go` da pasta **client** e avaliar o arquivo `cotacao.txt` com o valor da cota atual;