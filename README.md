# avito-intership-backend-2024
Репозиторий с решением тестового задания для стажера backend в Avito.
---
## Решение
В рамках решения тестового задания применялся исключительно язык **Go**. В качестве постоянного хранилища для баннеров используется **Postgres**, закешированные баннеры обслуживаются с помощью **redis**. Решение реализовывает базовые условия [задания](https://github.com/avito-tech/backend-trainee-assignment-2024/blob/main/README.md), полностью удовлетворяющее [API](https://github.com/avito-tech/backend-trainee-assignment-2024/blob/main/api.yaml) для взаимодействия с баннерами. Отдельного сервиса авторизации нет, его логика сильно упрощена. В ходе решения я старался придерживаться чистой архитектуры кода. Для тестов можно воспользоваться заполненным [API](https://github.com/mvp-mogila/avito-intership-backend-2024/blob/main/testApi.json)

## Запуск приложения

### Конфигурация
Для успешного запуска в следующих конфигурационных файлах должна быть заполнена нужная информация:
1. Makefile: postgres_user, postgres_password, host, port
2. config.yaml: основной конфиг приложения
3. docker-compose.yaml: файл для заупска контейнеров

### Команды Makefile:
- containers-up - запуск котейнеров с хранилищами
- containers-down - остановка конйтенеров
- migrations-up - установка утилиты goose и накат миграций в работающий контейнер postgres
- migrations-down - откат миграций
- run - запуск контейнеров и приложения

