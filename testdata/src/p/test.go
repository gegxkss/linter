package p

import (
	"log"
	"log/slog"
)

func CheckValidMessages() {
	slog.Info("starting server on port 8080")
	log.Fatal("failed to connect to database")
	slog.Debug("")
	slog.Info("lowercase")
}

func CheckInvalidMEssage() {
	log.Print("Starting server on port 8080")   // want "message must start with a lowercase letter"
	slog.Error("Failed to connect to database") // want "message must start with a lowercase letter"
	slog.Warn("Something went wrong")           // want "message must start with a lowercase letter"
	log.Printf("Connection failed")             // want "message must start with a lowercase letter"
}

func CheckIgnored() {
	message := "Hello"
	log.Print(message)
}

func checkLanguageAndSymbols() {
	slog.Error("ошибка подключения к базе данных")    // want "message must be english letters without special symbols only"
	slog.Error("Oшибка подключения к базе данных")    // want "message must be english letters without special symbols only" "message must start with a lowercase letter"
	log.Print("server started!🚀")                     // want "message must be english letters without special symbols only"
	log.Fatal("connection failed!!!")                 // want "message must be english letters without special symbols only"
	slog.Warn("Warning: something went wrong...")     // want "message must be english letters without special symbols only"  "message must start with a lowercase letter"
	slog.Error("Oшибка подключения к базе данных!!!") // want "message must be english letters without special symbols only"  "message must start with a lowercase letter"
	log.Fatal("connection.")                          // want "message must be english letters without special symbols only"
	log.Print("starting server on port 8080")
	slog.Error("failed to connect to database")
	log.Print("starting server on port 8080.") // want "message must be english letters without special symbols only"
	slog.Info("user password 12345")
	slog.Info("user password: 12345") // want "message must be english letters without special symbols only" "the message contains sensitive data"
}

func checkSensetiveData() {
	slog.Info("password: 12345")       // want "the message contains sensitive data" "message must be english letters without special symbols only"
	slog.Info("token=abc123")          // want "the message contains sensitive data" "message must be english letters without special symbols only"
	slog.Info("userPassword: qwerty")  // want "the message contains sensitive data" "message must be english letters without special symbols only"
	slog.Debug("api_key=secretkey123") // want "the message contains sensitive data" "message must be english letters without special symbols only"

	log.Print("user authenticated successfully")
	slog.Info("token validated")
	slog.Debug("api request completed")
}
