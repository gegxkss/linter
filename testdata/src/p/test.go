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
