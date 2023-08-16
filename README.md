# Boilerplate for nginx with Let’s Encrypt on docker-compose


## Инсталляция

Конфигурация для локальной разработки.

Для переноса на домен:
1. Раскомментировать директивы для Let`s Encrypt

2. Закомментировать подключение самоподписанных сертификатов 

3. Исправить конфигурации:
        - Add domains and email addresses to init-letsencrypt.sh
        - Replace all occurrences of example.org with primary domain (the first one you added to init-letsencrypt.sh) in data/nginx/app.conf

4. Выполнить скрипт:
        ./init-letsencrypt.sh
        `init-letsencrypt.sh` извлекает и обеспечивает продление сертификата Let’s

5. Run the server:
        docker-compose --compatibility up -d

6. Stop the server:
        docker-compose stop

PHPMyAdmin
https://localhost/pma
        login: root
        password: irol_root_pw

2023.08.16
Добавлены теги привязки в модели для валидации Валидация JSON с помощью ShouldBindJSON
Создан processValidationErrors для обработки ошибок
Добавлен validateCreateUserInput для пользовательской валидации 
Описательные комментарии для каждого шага Соответствующие ответы на ошибки для каждого случая отказа

/main.go
Key points:
Wrap initialization steps like loading .env and logger in error checks
Catch database connection error
Register custom error handling middleware
Log startup and shutdown messages
Return non-zero exit code on failures
This ensures:

Errors do not get ignored silently
Server process exits if it cannot start up properly
Custom middleware can handle errors centrally
Clean startup and shutdown with logging

config/config.go
Key points:
Return errors instead of logging/panic
Wrap errors with context using fmt.Errorf
Ping raw database handle to validate connection
Load() validates required env vars
All errors bubbled up to caller
This ensures any issues initializing config are surfaced instead of failing silently.

handlers/handlers.go
Key points:
Handle DB errors by aborting with 5xx status
Handle validation errors with 4xx status
Bubble errors up to caller instead of logging
Set global DB handle for simplified handlers
This allows robust error handling without complex nested blocks.

middleware/middleware.go
Key points:
Auth middleware returns 401 Unauthorized on error
ErrorHandler catches errors and returns 500 Internal Server Error
Helper validates token and returns error
Errors bubbled up instead of logging
This allows centralized error handling without complex error checks.

models/models.go
Key points:
Input validation before interacting with DB
Return errors instead of logging/panic
Wrap raw DB errors with context
Handle DB "not found" error separately
This moves error handling closer to the logic while maintaining clean signatures.

routes/routes.go
The key points are:
Handler functions return errors instead of handling
AbortWithError to return errors to client
Input validation before business logic
Errors bubbled up from helpers like JWT generation
This centralizes error handling in middleware.

utils/jwt.go
Key points:

Custom error types for common failures
GenerateToken returns error instead of panicking
ValidateToken returns structured error
Errors handled by caller instead of within
This leads to clean composable token handling functions.

utils/logger.go
Key points:

Custom error for init failure
Set global Log instance
Helper methods add context to log calls
Handle config errors instead of panicking
Return errors to caller
This allows handling logger errors gracefully and using it cleanly.


## License
All code in this repository is licensed under the terms of the `MIT License`. For further information please refer to the `LICENSE` file.

