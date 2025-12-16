# Практическая работа №3
## Пронин Евгений Анатольевич ЭФМО-01-25


Примеры запросов:

ПОЛНЫЙ ПИСОК в requests.md

## Запуск проекта
### Зпуск проекта
```bash
make run
```
### Зпуск тестов
```bash
make test
```
### Собрать бинарник
```bash
make build
```

## Покрытие юнит-тестов

<img width="485" height="70" alt="image" src="https://github.com/user-attachments/assets/88070636-d137-41bf-8e27-866af887369a" />

## Структура проекта

```
├── Makefile
├── cmd
│   └── server
│       └── main.go
├── go.mod
├── internal
│   ├── api
│   │   ├── handlers.go
│   │   ├── handlers_test.go
│   │   ├── middleware.go
│   │   └── responses.go
│   └── storage
│       └── memory.go
└── requests.md
```

## GET /health

<img width="1040" height="482" alt="Снимок экрана 2025-12-15 в 21 29 47" src="https://github.com/user-attachments/assets/baddb71e-c71b-4385-9f2b-80f0001c79ec" />

## POST /tasks \ -H "Content-Type: application/json" \ -d '{"title":"Купить молоко"}'

<img width="1042" height="499" alt="Снимок экрана 2025-12-15 в 21 32 39" src="https://github.com/user-attachments/assets/a2e4788e-9ba2-447c-8136-5b345f4e8082" />

## GET /tasks

<img width="1031" height="546" alt="Снимок экрана 2025-12-15 в 21 33 27" src="https://github.com/user-attachments/assets/f244d54d-906c-44c8-a57c-d6e78863b9e4" />


## GET /tasks?q=молоко

<img width="1044" height="531" alt="Снимок экрана 2025-12-15 в 21 33 56" src="https://github.com/user-attachments/assets/873ac29e-10d2-4761-b00c-dff3d769ecc4" />


## GET /tasks/1

<img width="1039" height="447" alt="Снимок экрана 2025-12-15 в 21 34 20" src="https://github.com/user-attachments/assets/c4bd9a11-b1e6-42ad-9536-c4102502b555" />


## GET /tasks/9999

<img width="1022" height="418" alt="Снимок экрана 2025-12-15 в 21 34 46" src="https://github.com/user-attachments/assets/ff2bba58-ce24-41aa-99bb-035eff1e3425" />


## PATCH /tasks/1

<img width="1038" height="433" alt="Снимок экрана 2025-12-15 в 21 35 59" src="https://github.com/user-attachments/assets/0d5d921e-8533-42a9-b5f2-34db8e07cb1c" />


## DELETE /tasks/1

<img width="1043" height="362" alt="Снимок экрана 2025-12-15 в 21 36 32" src="https://github.com/user-attachments/assets/c33ef1ab-120d-4227-8ca5-c365202a137c" />


## POST /tasks \ -H "Content-Type: application/javascript" \ -d 'test'

<img width="1018" height="490" alt="image" src="https://github.com/user-attachments/assets/67ad161e-c57f-401f-941f-381c8523e504" />


## POST tasks \ -H "Content-Type: application/json" \ -d '{"title":"a"}'


<img width="1032" height="516" alt="image" src="https://github.com/user-attachments/assets/424585c0-b57a-462d-b648-88c243f1c95a" />
