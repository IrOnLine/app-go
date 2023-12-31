# Golang : Use gorm with mysql in gin

This repository guides to how ORM can be implemented in Golang. After cloning the code, follow below steps to let the project run on your system.

1. Run `go build main.go`

2. Run `go run main.go`

curl -X POST -H "Content-Type: application/json" -d '{"Name":"roman", "Email":"abro@yandex.ru"}' http://localhost:8080/users 

After running the commands, you can find server started on `8080` .

2023.08.16
Добавлены теги привязки в модели для валидации Валидация JSON с помощью ShouldBindJSON
Создан processValidationErrors для обработки ошибок
Добавлен validateCreateUserInput для пользовательской валидации 
Описательные комментарии для каждого шага 
Соответствующие ответы на ошибки для каждого случая отказа

/main.go
#(!)
Оберните шаги инициализации, такие как загрузка .env и логгера, в проверки ошибок
Перехват ошибок подключения к базе данных
Регистрация пользовательского промежуточного ПО для обработки ошибок
Регистрировать сообщения о запуске и завершении работы
Возвращать ненулевой код выхода при ошибках
Это гарантирует:
- Ошибки не игнорируются втихую
- Серверный процесс завершается, если он не может запуститься должным образом
- Пользовательское промежуточное ПО может централизованно обрабатывать ошибки
- Чистый запуск и завершение работы с протоколированием

config/config.go
#(!)
Возвращать ошибки вместо логирования/паники
Обернуть ошибки контекстом с помощью fmt.Errorf
Пинговать необработанный хэндл базы данных для проверки соединения
Load() проверяет необходимые env-вары
Все ошибки передаются вызывающей стороне.
Это гарантирует, что любые проблемы с инициализацией конфигурации будут выявлены, а не завершится молчанием.

handlers/handlers.go
#(!)
Обработка ошибок БД путем прерывания со статусом 5xx
Обработка ошибок валидации с помощью статуса 4xx
Выдавать ошибки вызывающей стороне вместо протоколирования
Устанавливать глобальный хэндл БД для упрощенных обработчиков.
Это позволяет обеспечить надежную обработку ошибок без сложных вложенных блоков.

middleware/middleware.go
#(!)
Промежуточное ПО Auth возвращает ошибку 401 Unauthorized
ErrorHandler перехватывает ошибки и возвращает 500 Internal Server Error
Хелпер проверяет токен и возвращает ошибку
Ошибки выводятся на экран вместо логирования
Это позволяет централизованно обрабатывать ошибки без сложных проверок.

models/models.go
#(!)
Валидация ввода перед взаимодействием с БД
Возврат ошибок вместо логирования/паники
Обернуть необработанные ошибки БД контекстом
Отдельно обрабатывать ошибку "не найдено" в БД.
Это позволяет приблизить обработку ошибок к логике, сохраняя при этом чистоту сигнатур.

routes/routes.go
#(!)
Функции-обработчики возвращают ошибки вместо обработки
AbortWithError для возврата ошибок клиенту
Валидация ввода перед бизнес-логикой
Ошибки передаются от вспомогательных функций, таких как генерация JWT.
Это позволяет централизовать обработку ошибок в промежуточном ПО.

utils/jwt.go
#(!)
Пользовательские типы ошибок для распространенных сбоев
GenerateToken возвращает ошибку вместо паники
ValidateToken возвращает структурированную ошибку
Ошибки обрабатываются вызывающей стороной, а не внутри
Это приводит к созданию чистых композитных функций обработки токенов.

utils/logger.go
#(!)
Пользовательская ошибка при сбое инициализации
Установка глобального экземпляра журнала
Вспомогательные методы добавляют контекст к вызовам журнала
Обработка ошибок конфигурации вместо паники
Возвращать ошибки вызывающей стороне
Это позволяет изящно обрабатывать ошибки логгера и чисто использовать его.

