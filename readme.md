# TODOlist

TODOlist - это пет проект Alif Academy: Golang, написанный на Golang с использованием Gin-gonic, JWT токенов, GORM и Postgresql в качестве хранилища.

## Установка

1. Клонируйте репозиторий
2. Откройте файл `/db/postgres.go` и измените конфигурацию базы данных на свою
3. Запустите проект с помощью команды `go run cmd/main.go`

## Сборка

Чтобы собрать проект в исполняемый файл, выполните следующие шаги:

### Windows
1. Откройте терминал и перейдите в корневую директорию проекта
2. Выполните команду `go build -o todolist.exe cmd/main.go`
3. Исполняемый файл `todolist.exe` будет создан в корневой директории проекта

### Linux и macOS
1. Откройте терминал и перейдите в корневую директорию проекта
2. Выполните команду `go build -o todolist cmd/main.go`
3. Исполняемый файл `todolist` будет создан в корневой директории проекта

## Запуск

Чтобы запустить исполняемый файл после сборки, выполните следующие шаги:

### Windows
1. Откройте терминал и перейдите в директорию, где находится исполняемый файл
2. Выполните команду `todolist.exe`

### Linux и macOS
1. Откройте терминал и перейдите в директорию, где находится исполняемый файл
2. Выполните команду `chmod +x todolist` для предоставления прав на исполнение файла
3. Запустите исполняемый файл с помощью команды `./todolist`

Проект запустится и будет доступен по указанному в конфигурации адресу и порту.

## Использование

Проект имеет следующие роуты:

### Задачи
- `POST /tasks/` - создание задачи
- `GET /tasks/` - получение всех задач пользователя
- `GET /tasks/completed` - получение завершенных задач
- `PATCH /tasks/:id/reassign` - переназначение задачи (только для администраторов)
- `GET /tasks/:id` - получение задачи по ID (только для администраторов)
- `PATCH /tasks/:id` - изменение статуса завершения задачи (только для администраторов)
- `DELETE /tasks/:id` - удаление задачи (только для администраторов)

### Пользователи
- `GET /users/` - получение всех пользователей (только для администраторов)
- `GET /users/:id` - получение пользователя по ID (только для администраторов)
- `PUT /users/:id` - обновление пользователя (только для администраторов)
- `DELETE /users/:id` - удаление пользователя (только для администраторов)

### Аутентификация
- `POST /auth/sign-up` - регистрация нового пользователя
- `POST /auth/sign-in` - вход в систему
