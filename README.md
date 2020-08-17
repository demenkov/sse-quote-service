# GO sse-quote service
### Установка
При билде у вас должны быть ключи доступа к репозиторию. Иначе го не сможет добавить себе в модули данный сервис.

### Настройка
Отредактировать параметры в env файле
###
    REDIS_HOST={указать хост}
    REDIS_PORT={указать порт}
    REDIS_PASSWORD={пароль если есть}
    SOCKET_PORT={порт на котором будет слушать сервис}
    REST_CACHE_INTERVAL={время которое будет храниться кеш, по дефолду 60 секунд(указывать в секундах)}
    SSE_CLOUD_URL={хост провайдера}
    CURL_CLOUD_PKEY={приватный ключ провайдера}
    
### Урлы по которым слушает сервис
    /sse/
### Каналы по которым слушает сервис
    sse-request {в этот канал отправляем}
    sse-response {этот канал слушаем}
### Формат(пример) запроса JSON
    {
        "symbols": "twtr", // string {можно указывать через замяпую twtr,spy,appl и.д}
    }
### Запуск приложения
    go run main.go (или пользоваться докерфайлом)
    