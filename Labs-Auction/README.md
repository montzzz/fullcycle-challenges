# Labs-Auction

O objetivo é oferecer uma base sólida para implementar um fluxo completo de leilões, desde a criação até o fechamento automático e o registro das ofertas recebidas.

### Funcionalidades

A aplicação expõe os seguintes endpoints:

#### Leilões (/auction)

- `GET` /auction → Listar todos os leilões.
- `GET` /auction/:auctionId → Consultar detalhes de um leilão específico.
- `POST` /auction → Criar um novo leilão.
- `GET` /auction/winner/:auctionId → Consultar o lance vencedor de um leilão.

#### Lances (/bid)

- `POST` /bid → Criar um novo lance em um leilão.
- `GET` /bid/:auctionId → Listar todos os lances de um leilão.

#### Usuários (/user)
- `GET` /user/:userId → Consultar informações de um usuário.

### Configuração  

Toda a configuração é gerenciada por variáveis de ambiente.
Para desenvolvimento local, crie um arquivo .env dentro do diretório cmd/auction.

#### Variáveis de Ambiente (.env)

- BATCH_INSERT_INTERVAL: Intervalo para inserção em lote dos lances no banco de dados.
- MAX_BATCH_SIZE: Quantidade máxima de lances por inserção em lote.
- AUCTION_INTERVAL: Duração após a qual um leilão é automaticamente encerrado (ex.: 20s, 5m, 1h).
- MONGO_INITDB_ROOT_USERNAME: Usuário root do container MongoDB.
- MONGO_INITDB_ROOT_PASSWORD: Senha root do container MongoDB.
- MONGODB_URL: String de conexão com o MongoDB.
- MONGODB_DB: Nome do banco de dados.

#### Exemplo .env

```bash
BATCH_INSERT_INTERVAL=20s
MAX_BATCH_SIZE=4
AUCTION_INTERVAL=20s

MONGO_INITDB_ROOT_USERNAME=admin
MONGO_INITDB_ROOT_PASSWORD=admin
MONGODB_URL=mongodb://admin:admin@mongodb:27017/auctions?authSource=admin
MONGODB_DB=auctions
```

### Testar a API

#### Inicializar o docker-compose
```bash
docker compose up --build -d
```

#### Criar um leilão

Enviar uma solicitação POST para criar um novo leilão. Para testar o fechamento automático, defina AUCTION_INTERVAL para uma breve duração

```bash
curl -X POST http://localhost:8080/auction \
-H 'Content-Type: application/json' \
-d '{
  "product_name": "Vintage Watch",
  "category": "Accessories",
  "description": "A beautiful vintage watch.",
  "condition": 1
}'
```

#### Verificar fechamento automático

Depois da duração do leilão, envie uma solicitação GET para verificar o status do leilão:

```bash
curl -X GET http://localhost:8080/auction/{{YOUR_AUCTION_ID}}
```

##### Resposta esperada
```json
{
  "id":"{{YOUR_AUCTION_ID}}",
  "product_name":"Vintage Watch",
  "category":"Accessories",
  "description":"A beautiful vintage watch.",
  "condition":1,
  "status":1,
  "timestamp":"2025-08-24T00:49:20Z"
}
```