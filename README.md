# Основная информация
Язык написания: GoLang

База данных: Postgres (с помощью облачной PaaS-платформы Heroku) + драйвер: <a href="https://github.com/lib/pq">github.com/lib/pq</a>

ID реализованы с помощью uuid v.4 (библиотека: <a href="https://github.com/google/uuid">google/uuid</a>)

### Таблица all_users (таблица всех пользователей)
| id               | username | createdAt |
|------------------|----------|------------|
| TEXT PRIMARY KEY | TEXT     | INT        |
### Таблица chats (таблица чатов)
| id               | name | users | createdAt |
|------------------|------|:-----:|-----------|
| TEXT PRIMARY KEY | TEXT | TEXT  | INT       |
### Таблица messages (таблица всех сообщений)
| id               | chat | author | text | createdAt |
|------------------|------|:------:|------|-----------|
| TEXT PRIMARY KEY | TEXT | TEXT   | TEXT | INT       |

## Реализация API методов (7 штук)
#### + http://localhost:9000/users/add?username=(имя пользователя)
#### + fwq
