# Fiber Book Management API

Простое REST API для управления книгами, построенное на Go Fiber с JWT аутентификацией и SQLite базой данных.

## 🏗️ Архитектура проекта

```
fiber-golang-json-api/
├── cmd/                          # Точка входа приложения
│   └── server/
│       └── main.go              # Главный файл приложения
│
├── internal/                     # Внутренние пакеты проекта
│   ├── config/                  # Конфигурации
│   │   └── config.go           # Настройки приложения
│   ├── database/               # Подключение к БД
│   │   └── database.go         # Инициализация БД
│   ├── models/                 # Модели данных
│   │   └── models.go          # User, Book модели
│   ├── repository/             # Работа с базой данных (CRUD)
│   │   ├── user_repository.go
│   │   └── book_repository.go
│   ├── service/                # Бизнес-логика
│   │   ├── user_service.go
│   │   └── book_service.go
│   ├── handler/                # HTTP обработчики (Fiber)
│   │   ├── auth_handler.go
│   │   ├── book_handler.go
│   │   └── download_handler.go
│   ├── middleware/             # Fiber middleware
│   │   └── auth_middleware.go
│   └── util/                  # Утилитарные функции
│       └── jwt.go            # JWT токены
│
├── pkg/                       # Пакеты для других проектов
├── go.mod
├── go.sum
└── storage.db                # SQLite база данных
```

## 🚀 Быстрый старт

### Предварительные требования

- Go 1.22+ 
- Git

### Установка и запуск

1. **Клонируйте репозиторий:**
```bash
git clone <repository-url>
cd fiber-golang-json-api
```

2. **Установите зависимости:**
```bash
go mod tidy
```

3. **Запустите сервер (выберите один из способов):**

**Способ 1 - через go run (для разработки):**
```bash
go run ./cmd/server/main.go
```

**Способ 2 - через сборку (для продакшена):**
```bash
go build -o fiber-app ./cmd/server
./fiber-app
```

**Способ 3 - через Makefile:**
```bash
make run
```

Сервер запустится на `http://localhost:3000`

## 📚 API Документация

### Аутентификация

#### Регистрация пользователя
```bash
POST /auth/register
Content-Type: application/x-www-form-urlencoded

username=testuser&password=testpass123
```

**Ответ:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### Вход пользователя
```bash
POST /auth/login
Content-Type: application/x-www-form-urlencoded

username=testuser&password=testpass123
```

**Ответ:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Управление книгами (требует аутентификации)

#### Получить все книги
```bash
GET /book/
Authorization: Bearer <token>
```

**Фильтрация:**
```bash
GET /book/?status=reading&author=Толстой&year=1869
Authorization: Bearer <token>
```

**Доступные фильтры:**
- `status` - статус книги (to_read, reading, read)
- `author` - автор (частичное совпадение)
- `title` - название (частичное совпадение)
- `year` - год издания

#### Получить книгу по ID
```bash
GET /book/{id}
Authorization: Bearer <token>
```

#### Создать книгу
```bash
POST /book/
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Война и мир",
  "author": "Лев Толстой",
  "year": 1869,
  "status": "to_read"
}
```

#### Обновить книгу
```bash
PUT /book/{id}
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Новое название",
  "status": "read"
}
```

#### Удалить книгу
```bash
DELETE /book/{id}
Authorization: Bearer <token>
```

### Экспорт данных (без аутентификации)

#### Скачать все книги в JSON
```bash
GET /download/?format=json
```

#### Скачать все книги в CSV
```bash
GET /download/?format=csv
```

## 🛠️ Разработка

### Структура проекта

- **cmd/server/main.go** - точка входа приложения
- **internal/config/** - конфигурация приложения
- **internal/database/** - подключение к базе данных
- **internal/models/** - модели данных (User, Book)
- **internal/repository/** - слой доступа к данным
- **internal/service/** - бизнес-логика
- **internal/handler/** - HTTP обработчики
- **internal/middleware/** - middleware для Fiber
- **internal/util/** - утилитарные функции

### Переменные окружения

Создайте файл `.env` (опционально):

```env
PORT=3000
DB_DSN=storage.db
JWT_SECRET=your-secret-key
```

### Тестирование API

Используйте curl или любой HTTP клиент (Postman, Insomnia):

```bash
# Регистрация
curl -X POST http://localhost:3000/auth/register \
  -d "username=testuser&password=testpass123"

# Создание книги
curl -X POST http://localhost:3000/book/ \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"title": "Test Book", "author": "Test Author", "year": 2024, "status": "to_read"}'

# Получение книг
curl -X GET http://localhost:3000/book/ \
  -H "Authorization: Bearer <token>"

# Скачивание в CSV
curl -X GET "http://localhost:3000/download/?format=csv"
```

## 📊 Модели данных

### User
```go
type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Username string `json:"username" gorm:"unique"`
    Password string `json:"-"` // не возвращается в JSON
}
```

### Book
```go
type Book struct {
    ID     uint   `json:"id" gorm:"primaryKey"`
    Title  string `json:"title"`
    Status Status `json:"status" gorm:"default:to_read"`
    Author string `json:"author"`
    Year   int    `json:"year"`
    UserID uint   `json:"user_id"`
}
```

### Status
```go
type Status string

const (
    StatusToRead  Status = "to_read"
    StatusReading Status = "reading"
    StatusRead    Status = "read"
)
```

## 🔧 Технологии

- **Go 1.22+** - язык программирования
- **Fiber v2** - веб-фреймворк
- **GORM** - ORM для работы с базой данных
- **SQLite** - база данных
- **JWT** - аутентификация
- **bcrypt** - хеширование паролей

## 📝 Лицензия

MIT License

## 🤝 Вклад в проект

1. Fork репозиторий
2. Создайте feature branch (`git checkout -b feature/amazing-feature`)
3. Commit изменения (`git commit -m 'Add amazing feature'`)
4. Push в branch (`git push origin feature/amazing-feature`)
5. Откройте Pull Request

## 📞 Поддержка

Если у вас есть вопросы или проблемы, создайте issue в репозитории.
