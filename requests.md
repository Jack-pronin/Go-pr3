# Тестовые HTTP-запросы (PZ3)

Данные запросы использовались для проверки работы HTTP API сервиса управления задачами.
Тестирование выполнялось с помощью Postman, ниже приведены эквивалентные curl-запросы.

---

## 1. Проверка состояния сервера
GET /health

```bash
curl -i http://localhost:8080/health
```

## 2. Создание задачи
POST /tasks

```bash
curl -i -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{"title":"Купить молоко"}'
```

## 3. Получение списка задач

GET /tasks
```bash
curl -i http://localhost:8080/tasks
```
## 4. Фильтрация задач по query-параметру

GET /tasks?q=молоко
```bash
curl -i "http://localhost:8080/tasks?q=молоко"
```
## 5. Получение задачи по идентификатору

GET /tasks/{id}
```bash
curl -i http://localhost:8080/tasks/1
```
## 6. Ошибка: задача не найдена

GET /tasks/9999
```bash
curl -i http://localhost:8080/tasks/9999
```
## 7. Изменение задачи

PATCH /tasks/{id}
```bash
curl -i -X PATCH http://localhost:8080/tasks/1
```
## 8. Удаление задачи

DELETE /tasks/{id}
```bash
curl -i -X DELETE http://localhost:8080/tasks/1
```
## 9. Ошибка Content-Type

POST /tasks с неверным Content-Type
```bash
curl -i -X POST http://localhost:8080/tasks \
-H "Content-Type: application/javascript" \
-d 'test'
```
## 10. Ошибка валидации

POST /tasks — неверная длина title
```bash
curl -i -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{"title":"a"}'
```