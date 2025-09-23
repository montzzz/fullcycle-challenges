# Zip-Weather Gateway

O **Zip-Weather Gateway** é o serviço responsável por receber um CEP via API, validar o input e repassar para o serviço de backend (**zip-weather**), que faz a busca do endereço e da temperatura.  

Este projeto também possui **tracing distribuído** implementado usando **OpenTelemetry (OTEL)** e **Zipkin**, permitindo monitoramento completo do fluxo A → B → APIs externas.

---

## Funcionalidades principais

- Recebe CEP via POST `/weather`  
- Valida formato do CEP (`8 dígitos`)  
- Chama o `zip-weather` para buscar endereço e temperatura  
- Tracing distribuído com spans para:
  - **CEP lookup** (ViaCEP)
  - **Weather lookup** (WeatherAPI)
- Exporta traces para Zipkin ou outro backend OTEL compatível

---

## Endpoints

### POST `/weather`
**Request Body:**
```json
{
  "cep": "29902555"
}
```

## Response Body (exemplo):
```json
{
   "city":"Linhares",
   "temp_c": 9.9,
   "temp_f": 49.82,
   "temp_k": 282.9
}
```

# Uso via Docker-Compose

É possível testar os dois serviços (zip-weather-gateway + zip-weather) junto com o Zipkin:  
  - Certifique-se de ter o Docker e Docker Compose instalados;
  - Configure a **env** `ZIP_WEATHER_API_URL` para o `zip-weather-gateway` dentro do arquivo _docker-compose.yml_ (caso queira testar com um backend que não esteja rodando localmente);
  - Configure a **env** `WEATHER_API_KEY` para o zip-weather dentro do arquivo _docker-compose.yml_;

#### Rodando os apps
```bash
docker-compose up --build -d
```

- Caso queira parar todos os containers e removê-los rode: 
```bash
docker-compose down
```


# Testando a API

```bash
curl -X POST http://localhost:8080/weather \
  -H "Content-Type: application/json" \
  -d '{"cep": "29902555"}'
```

# Tracing

- Caso queira ver o tracing completo de cada request, poderá acessar o `Zipkin UI` pela url: `http://localhost:9411/zipkin/`;
- Poderá procurar por **serviceName**, onde vai conter os services `zip-gateway` e `zip-weather`;
  - `zip-gateway`: Mostra os dados dos traces que ocorreram nesse serviço de entrada (do handler até o get para o zip-weather);
  - `zip-weather`: Mostra os dados dos traces que ocorreram nesse serviço, juntamente com os spans para as apis externas utilizadas;
