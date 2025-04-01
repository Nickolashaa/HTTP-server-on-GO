# HTTP Server on GO

![Go](https://img.shields.io/badge/Go-1.21-blue)
![License](https://img.shields.io/badge/License-MIT-green)

## 📌 Описание

Этот проект — простой HTTP-сервер на Go, который поддерживает работу с пользователями и книгами через REST API.

## 🚀 Установка

1. **Скачайте и установите Golang** (если у вас его нет) с [официального сайта](https://go.dev/).
2. **Клонируйте репозиторий:**  
   ```sh
   git clone https://github.com/Nickolashaa/HTTP-server-on-GO.git
   cd HTTP-server-on-GO
   ```
3. **Установите зависимости:**  
   ```sh
   go mod download  # или go mod tidy
   ```
4. **Запустите сервер:**  
   ```sh
   go run .\main.go
   ```

## 📡 API Эндпоинты

### 🟢 GET

#### 🔹 Получить пользователя по ID
**Запрос:**  
`GET /users/{id}`  

**Пример:**  
```sh
curl -X GET http://localhost:8080/users/1
```
**Ответ:**  
```json
{
    "id": 1
}
```

#### 🔹 Получить список всех книг
**Запрос:**  
`GET /books`  

**Пример:**  
```sh
curl -X GET http://localhost:8080/books
```
**Ответ:**  
```json
{
    "0": { "Title": "СУПЕРСКАЯ КНИГА" },
    "1": { "Title": "ЕЩЕ ОДНА СУПЕРСКАЯ КНИГА" },
    "2": { "Title": "ПОСЛЕДНЯЯ СУПЕРСКАЯ КНИГА" }
}
```

#### 🔹 Получить книгу по ID
**Запрос:**  
`GET /books/{id}`  

**Пример (успех):**  
```sh
curl -X GET http://localhost:8080/books/1
```
**Ответ:**  
```json
{
    "Title": "ЕЩЕ ОДНА СУПЕРСКАЯ КНИГА"
}
```

**Пример (ошибка):**  
```sh
curl -X GET http://localhost:8080/books/4
```
**Ответ:**  
```
404 Not Found
```

---

### 🟠 POST

#### 🔹 Создать нового пользователя
**Запрос:**  
`POST /users`  

**Тело запроса:**  
```json
{
    "Name": "Иван",
    "email": "ivan@example.com"
}
```

**Пример:**  
```sh
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"Name": "Иван", "email": "ivan@example.com"}'
```
**Ответ:**  
```json
{
    "id": 0
}
```

#### 🔹 Добавить книгу
**Запрос:**  
`POST /books`  

**Тело запроса:**  
```json
{
    "title": "СУПЕРСКАЯ КНИГА"
}
```

**Пример:**  
```sh
curl -X POST http://localhost:8080/books -H "Content-Type: application/json" -d '{"title": "СУПЕРСКАЯ КНИГА"}'
```
**Ответ:**  
```json
{
    "id": 0
}
```

---

## 📜 Лицензия

Этот проект распространяется под лицензией MIT. Подробнее см. [LICENSE](LICENSE).

## ✒️ Автор

Разработчик: [Anti6eptik](https://github.com/Anti6eptik)  
Оригинальный репозиторий: [Nickolashaa](https://github.com/Nickolashaa)
