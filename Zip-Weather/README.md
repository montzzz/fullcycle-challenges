# Weather by CEP

Este projeto é uma API desenvolvida em **Go** que permite consultar a localização de um **CEP brasileiro** e, a partir disso, retornar a **temperatura atual** formatada em **Celsius, Fahrenheit e Kelvin**.

---

## Funcionalidades

- Recebe um **CEP válido** (8 dígitos) via query parameter.
- Consulta a localização correspondente ao CEP.
- Obtém a temperatura atual da cidade.
- Retorna a temperatura convertida em:
  - 🌡️ Celsius
  - 🌡️ Fahrenheit
  - 🌡️ Kelvin
- Tratamento de erros para entradas inválidas ou CEPs inexistentes.

---

## Endpoints

### Consultar temperatura por CEP

**Request:**

```http
GET http://localhost:8080/weather?cep=01001000
```

#### Sucesso (200)

```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.6
}
```

#### Erros

**422 Unprocessable Entity (formato incorreto):**

```json
{
  "message": "invalid zipcode"
}
```

**404 Not Found (CEP não encontrado):**

```json
{
  "message": "invalid zipcode"
}
```


# Uso via Go

#### Defina a variável de ambiente para a key do weatherapi  
    
- Poderá definir manualmente pelo arquivo `.env` que se encontra na raíz do projeto;
- Exportando diretamente no terminal, exemplo:  
    ```bash
        export WEATHER_API_KEY=sua_chave_aqui
    ```

#### Executando o app

```bash
go run cmd/main.go
```

# Uso via Docker

#### Build da imagem
```bash
docker build -t weather-api .
```

#### Executar a imagem
```bash
docker run --rm weather-api \
  -e WEATHER_API_KEY=sua_chave_aqui \
```

# Teste

Para teste, execute o curl abaixo e valide se a resposta vem corretamente com um CEP válido:  
```bash
curl "http://localhost:8080/weather?cep=01001000"
```

# Teste via endpoint cloudrun

O endpoint para acesso é o abaixo:
```bash
https://cloudrun-zipweather-3vhdzphkna-uc.a.run.app/weather?cep=01001000
```
