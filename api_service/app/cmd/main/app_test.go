package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/theartofdevel/notes_system/api_service/internal/config"
	"github.com/theartofdevel/notes_system/api_service/pkg/logging"
)

func TestMainFunction(t *testing.T) {
	// Переопределяем функции, которые вызываются в main(), для упрощения тестирования
	logging.Init = func() {}
	config.GetConfig = func() *config.Config {
		return &config.Config{
			Listen: config.ListenConfig{
				Type:   "tcp",
				BindIP: "127.0.0.1",
				Port:   "8080",
			},
			UserService:     config.ServiceConfig{URL: "http://localhost:8001"},
			CategoryService: config.ServiceConfig{URL: "http://localhost:8002"},
			NoteService:     config.ServiceConfig{URL: "http://localhost:8003"},
			TagService:      config.ServiceConfig{URL: "http://localhost:8004"},
		}
	}
	logging.GetLogger = func() logging.Logger {
		return logging.Logger{}
	}

	// Запускаем main() в отдельной горутине, чтобы он не блокировал тесты
	go main()

	// Даем время для инициализации сервера
	time.Sleep(2 * time.Second)

	// Отправляем тестовый запрос к серверу
	resp, err := http.Get("http://127.0.0.1:8080/metrics")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK; got %v", resp.Status)
	}
}

func TestStartFunction(t *testing.T) {
	logger := logging.GetLogger()
	router := httprouter.New()
	cfg := &config.Config{
		Listen: config.ListenConfig{
			Type:   "tcp",
			BindIP: "127.0.0.1",
			Port:   "8081",
		},
	}

	go start(router, logger, cfg)

	time.Sleep(2 * time.Second)

	// Отправляем тестовый запрос к серверу
	resp, err := http.Get("http://127.0.0.1:8081/metrics")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK; got %v", resp.Status)
	}
}