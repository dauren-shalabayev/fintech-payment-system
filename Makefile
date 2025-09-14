# Makefile для Fiber Book Management API

# Переменные
BINARY_NAME=fiber-app
BUILD_DIR=./cmd/server
MAIN_FILE=./cmd/server/main.go

# Цвета для вывода
GREEN=\033[0;32m
YELLOW=\033[1;33m
RED=\033[0;31m
NC=\033[0m # No Color

.PHONY: help build run clean test deps dev

# Показать справку
help:
	@echo "$(GREEN)Fiber Book Management API$(NC)"
	@echo ""
	@echo "$(YELLOW)Доступные команды:$(NC)"
	@echo "  $(GREEN)build$(NC)     - Собрать приложение"
	@echo "  $(GREEN)run$(NC)       - Запустить сервер"
	@echo "  $(GREEN)dev$(NC)       - Запустить в режиме разработки"
	@echo "  $(GREEN)clean$(NC)     - Очистить собранные файлы"
	@echo "  $(GREEN)test$(NC)      - Запустить тесты"
	@echo "  $(GREEN)deps$(NC)      - Установить зависимости"
	@echo "  $(GREEN)help$(NC)      - Показать эту справку"

# Собрать приложение
build:
	@echo "$(YELLOW)Сборка приложения...$(NC)"
	@go build -o $(BINARY_NAME) $(BUILD_DIR)
	@echo "$(GREEN)✓ Приложение собрано: $(BINARY_NAME)$(NC)"

# Запустить сервер
run: build
	@echo "$(YELLOW)Запуск сервера...$(NC)"
	@echo "$(GREEN)Сервер доступен на: http://localhost:3000$(NC)"
	@./$(BINARY_NAME)

# Режим разработки (автоперезагрузка)
dev:
	@echo "$(YELLOW)Запуск в режиме разработки...$(NC)"
	@echo "$(GREEN)Установите air для автоперезагрузки: go install github.com/cosmtrek/air@latest$(NC)"
	@air

# Очистить собранные файлы
clean:
	@echo "$(YELLOW)Очистка...$(NC)"
	@rm -f $(BINARY_NAME)
	@rm -f storage.db
	@echo "$(GREEN)✓ Очистка завершена$(NC)"

# Запустить тесты
test:
	@echo "$(YELLOW)Запуск тестов...$(NC)"
	@go test ./...

# Установить зависимости
deps:
	@echo "$(YELLOW)Установка зависимостей...$(NC)"
	@go mod tidy
	@go mod download
	@echo "$(GREEN)✓ Зависимости установлены$(NC)"

# Проверить код
lint:
	@echo "$(YELLOW)Проверка кода...$(NC)"
	@golangci-lint run

# Форматировать код
fmt:
	@echo "$(YELLOW)Форматирование кода...$(NC)"
	@go fmt ./...
	@echo "$(GREEN)✓ Код отформатирован$(NC)"

# Показать информацию о проекте
info:
	@echo "$(GREEN)Fiber Book Management API$(NC)"
	@echo "$(YELLOW)Версия Go:$(NC) $(shell go version)"
	@echo "$(YELLOW)Архитектура:$(NC) $(shell go env GOOS)/$(shell go env GOARCH)"
	@echo "$(YELLOW)Модуль:$(NC) $(shell go list -m)"
