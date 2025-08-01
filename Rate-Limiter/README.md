# # Desafio Go - Rate Limiter

Um middleware de **Rate Limiting em Go** que controla requisições por **IP** ou por **Token de Acesso**, utilizando o **Redis** como mecanismo de persistência.

## Funcionalidades

- ✅ Limita o número de requisições por IP ou por Token
- ✅ Prioridade para tokens conhecidos (limitados individualmente)
- ✅ Redis como backend de contagem e bloqueio
- ✅ Configuração via `.env`
- ✅ Middleware único, fácil de integrar
- ✅ Estrutura modular (clean architecture style)

## Como Funciona

- Por padrão, todas as requisições são limitadas por IP.
- Se o header `API_KEY` for enviado **e o token for conhecido**, o limite será aplicado com base no token (e substituirá o limite por IP).
- Se o limite for excedido:
  - O IP ou Token será **bloqueado temporariamente**
  - Será retornado o status `429 Too Many Requests`

---

## Exemplo de `.env`

```env
# Limite padrão por IP
RATE_LIMIT_DEFAULT=10

# Tempo de bloqueio (em segundos)
BLOCK_DURATION=30

# Conexão Redis
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# Lista de tokens (token:limite), separados por ";"
TOKENS=token123:100;vipUser:500
```


## Executando o Projeto

1. Suba o Redis com Docker  
```docker-compose up -d```
2. Instale as dependências  
```go mod tidy```
3. Execute o servidor  
```go run cmd/main.go```


## Fazendo Requisições

#### Com token`
```bash
curl -H "API_KEY: token123" http://localhost:8080
```

#### Sem Token (usa o IP)
```bash
curl http://localhost:8080
```

#### Adicionando tokens

- Você pode adicionar tokens manualmente no **.env**;
  - Cada token pode ter seu valor para limite, seguindo o padrão __${token}:${limite}__, seguindo o exemplo:  ```TOKENS=token1:100;token2:50```

- Caso você tenha o **CLI** do `Redis` instalado e conectado ao servidor do docker-compose que está rodando, basta executar o comando abaixo:
```bash
redis-cli set API_KEY_LIMIT:tokenEspecial 1000
```
  - O prefixo no Redis deve ser __API_KEY_LIMIT:${TOKEN} {$LIMIT}__