package main

import (
	"bitcly/reader"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

const (
	base        = "reader/data"
	encodesFile = "encodes.csv"
	decodesFile = "decodes.json"
)

func main() {
	if err := loadEnv(); err != nil {
		slog.Error("Error loading environment:", err)
		os.Exit(-1)
	}

	hashes, err := reader.GetHashes(base, encodesFile)
	if err != nil {
		slog.Error("Error reading file of hash associations:", err)
		os.Exit(-1)
	}

	clicks, err := reader.GetClicks(base, decodesFile)
	if err != nil {
		slog.Error("Error reading file of click events:", err)
		os.Exit(-1)
	}

	counts, err := clicks.Process(hashes)
	if err != nil {
		slog.Error("Error processing the result:", err)
		os.Exit(-1)
	}

	slog.Info("Successfully processed click counts per URL")
	fmt.Println(counts)
}

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	level := slog.LevelInfo
	switch os.Getenv("LOG_LEVEL") {
	case "-4":
		level = slog.LevelDebug
	case "0":
	case "4":
		level = slog.LevelWarn
	case "8":
		level = slog.LevelError
	default:
		slog.Error("LOG LEVEL value not supported, defaulting to INFO", "level", os.Getenv("LOG_LEVEL"))
	}

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	slog.SetDefault(slog.New(handler))

	return nil
}
