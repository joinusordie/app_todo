# Проект REST API по созданию TODO списков на Go

## В проекте разобраны следующие концепции:
- Разработка Веб-Приложений на Go, следуя дизайну REST API.
- Работа с фреймворком <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>.
- Подход Чистой Архитектуры в построении структуры приложения. Техника внедрения зависимости.
- Работа с БД Postgres. Запуск из Docker. Генерация файлов миграций. 
- Конфигурация приложения с помощью библиотеки <a href="https://github.com/spf13/viper">spf13/viper</a>. Работа с переменными окружения.
- Работа с БД, используя библиотеку <a href="https://github.com/jmoiron/sqlx">sqlx</a>.
- Регистрация и аутентификация. Работа с JWT. Middleware.
- Написание SQL запросов.
- Graceful Shutdown

### Для запуска приложения:

```
make build && make run
```
При первом запуске нужео сделать маграции к БД
```
make migrate
```

## Демонстрация работы
https://youtu.be/vVzGgmpGGiY

### Frontend-часть проекта можно посмотреть по ссылке:
https://github.com/ArtKolyach/todo-react
