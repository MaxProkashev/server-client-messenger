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

В любом методе вместо ... следует писать, либо id в формате uuid, либо новое имя пользователя или чата

#### - ```http://localhost:9000/users/add?username=...```
возвращает id нового пользователя, записывает лог на стороне сервиса, создает row в таблице all_users
#### - ```http://localhost:9000/chats/add?name=...;users=...,...```
возвращает id нового чата, записывает лог на стороне сервиса, создает row в таблице chats
#### - ```http://localhost:9000/messages/add?chat=...;author=...;text=...```
возвращает id нового сообщения, записывает лог на стороне сервиса, создает row в таблице messages
#### - ```http://localhost:9000/users/show```
выводит список всех пользователей (id,username,date в формате time.Time), записывает лог
#### - ```http://localhost:9000/chats/get?user=...```
выводит список всех чатов конкретного пользователя, выводит все колонки из базы данных(id,name,users - вертикальным списком,date в формате time.Time), логирует
#### - ```http://localhost:9000/messages/get?chat=...```
выводит список всех сообщений конкретного чата, выводит все колонки из базы данных(id,chat,author,text,date в формате time.Time), логирует
#### - ```http://localhost:9000/drop```
сбрасывает все таблицы и создает их заново

## Способ запуска

клонировать репу --> в консоле запустить команду **go run *.go** --> либо либо

--> в браузере 

--> в другой консоле (curl, ..)

##### Пример запуска:
```
>>> curl http://localhost:9000/users/add?username=qwe
New user ID: 4c5a5ce5-e797-4594-a247-16ef42d3ff82
>>> curl http://localhost:9000/users/add?username=ewq
New user ID: d9950723-74a7-491a-8f71-c92985ec9ae8
>>> curl http://localhost:9000/users/add?username=123
New user ID: 59ea5ed5-ebf6-4387-bdcb-8867ba6df0e8
>>> curl http://localhost:9000/users/show
All Users ID:

ID: 4c5a5ce5-e797-4594-a247-16ef42d3ff82
Username: qwe
Create time: Mon Aug 17 19:44:50 MSK 2020

ID: d9950723-74a7-491a-8f71-c92985ec9ae8
Username: ewq
Create time: Mon Aug 17 19:44:55 MSK 2020

ID: 59ea5ed5-ebf6-4387-bdcb-8867ba6df0e8
Username: 123
Create time: Mon Aug 17 19:45:18 MSK 2020
>>> curl http://localhost:9000/chats/add?name=newchat;users=4c5a5ce5-e797-4594-a247-16ef42d3ff82,d9950723-74a7-491a-8f71-c92985ec9ae8
New chat ID: c42df210-ee16-40c4-ab5e-3b3692275aed
>>> curl http://localhost:9000/messages/add?chat=c42df210-ee16-40c4-ab5e-3b3692275aed;author=4c5a5ce5-e797-4594-a247-16ef42d3ff82;text=hello1
New message ID: b2a4a5ef-cd91-440e-b4f9-3acef417cd75
>>> curl http://localhost:9000/messages/add?chat=c42df210-ee16-40c4-ab5e-3b3692275aed;author=d9950723-74a7-491a-8f71-c92985ec9ae8;text=hello2
New message ID: 679b0734-f263-4794-8bde-521a2af71f4a
>>> curl http://localhost:9000/chats/get?user=4c5a5ce5-e797-4594-a247-16ef42d3ff82
All Chats User ID 4c5a5ce5-e797-4594-a247-16ef42d3ff82:

ID: c42df210-ee16-40c4-ab5e-3b3692275aed
Name: newchat
Users:
4c5a5ce5-e797-4594-a247-16ef42d3ff82
d9950723-74a7-491a-8f71-c92985ec9ae8

Create time: Mon Aug 17 19:49:07 MSK 2020

>>> curl http://localhost:9000/messages/get?chat=c42df210-ee16-40c4-ab5e-3b3692275aed
All Message Chat ID c42df210-ee16-40c4-ab5e-3b3692275aed:

ID: b2a4a5ef-cd91-440e-b4f9-3acef417cd75
Chat: c42df210-ee16-40c4-ab5e-3b3692275aed
Author: 4c5a5ce5-e797-4594-a247-16ef42d3ff82
Text: hello1
Create time: Mon Aug 17 19:51:19 MSK 2020

ID: 679b0734-f263-4794-8bde-521a2af71f4a
Chat: c42df210-ee16-40c4-ab5e-3b3692275aed
Author: d9950723-74a7-491a-8f71-c92985ec9ae8
Text: hello2
Create time: Mon Aug 17 19:51:36 MSK 2020
```

##### пример лога:

```
$ go run *.go
2020/08/17 19:44:19 [OK] Create all_users table
2020/08/17 19:44:19 [OK] Create chats table
2020/08/17 19:44:19 [OK] Create messages table
2020/08/17 19:44:31 [OK] Show user request
2020/08/17 19:44:50 [OK] New user 4c5a5ce5-e797-4594-a247-16ef42d3ff82
2020/08/17 19:44:55 [OK] New user d9950723-74a7-491a-8f71-c92985ec9ae8
2020/08/17 19:45:18 [OK] New user 59ea5ed5-ebf6-4387-bdcb-8867ba6df0e8
2020/08/17 19:47:39 [OK] Show user request
....
```
