# Go Year Course

## Использует:
- [Golang](https://go.dev/)
- [net/http](https://pkg.go.dev/net/http)

##

## Как использовать
### Запуск приложения
```
go run ./cmd/main/main.go
```
> [!NOTE]
> Сервер поднимается по адресу **localhost:8080**

### Использование
Отправив POST запрос на ```/api/v1/calculate``` с телом JSON формата: 
```
{
    "expression": "выражение"
}
```
Вы получите ответ от сервера

## Примеры использования

Пример 1
#### **url:** ```http://localhost:8080/api/v1/calculate```
#### **Тело:**
```
{
    "expression": "2+2*2*2"
}
```
#### **Ответ**
***satus code: 200***
```
{
    "result": "10"
}
```

### Пример 2
#### **url:** ```http://localhost:8080/api/v1/calculate```
#### **Тело:**
```
{
    "expression": "smth"
}
```
#### **Ответ**
***satus code: 422***
```
{"error": "Expression is not valid"}

```

### Пример 3
#### **url:** ```http://localhost:8080/api/v1/calculate```
#### **Тело:**
```
Some data 
```
#### **Ответ**
***satus code: 500***
```
{"error": "Internal server error"}
```