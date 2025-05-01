# Desafio Go - API de Consulta de CEP

Este projeto é uma aplicação em Go que realiza a busca de informações de um CEP utilizando **duas APIs públicas** simultaneamente:

- [BrasilAPI](https://brasilapi.com.br/)
- [ViaCEP](https://viacep.com.br/)

A aplicação dispara requisições concorrentes e retorna a **primeira resposta recebida**. Caso nenhuma das APIs responda em até 1 segundo, a aplicação retorna um erro de timeout.

## 📝 Requisitos

Neste desafio, você precisará utilizar conceitos de **Multithreading** e **APIs** para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

- **BrasilAPI**: `"https://brasilapi.com.br/api/cep/v1/{cep}"`
- **ViaCEP**: `http://viacep.com.br/ws/{cep}/json/`

**Requisitos:**

- Acatar a resposta da API que entregar o resultado mais rápido e **descartar a resposta mais lenta**.
- Exibir o resultado no **command line** com os dados do endereço, bem como indicar qual API enviou a resposta.
- Limitar o tempo de resposta a **1 segundo**. Caso contrário, exibir o erro de **timeout**.

---

---

## Funcionalidade

- Recebe um CEP como argumento de linha de comando
- Consulta os serviços BrasilAPI e ViaCEP em paralelo
- Retorna a primeira resposta bem-sucedida
- Timeout automático de 1 segundo para evitar travamentos

---

#### Exemplo de uso

```bash
go run challenge/challenge.go 89120000
```

#### Exemplo de resposta

```json
{"CEP":"89120000", "Address": "" ,"Neighborhood": "", "City":"Timbó", "State":"SC", "OriginRequest":"BrasilAPI"}
```

---