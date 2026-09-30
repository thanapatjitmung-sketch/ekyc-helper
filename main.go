package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"time"
)

func init() {
	logDir := "log"

	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(err)
	}

	// วันที่รูปแบบ YYYYMMDD
	date := time.Now().Format("20060102")

	// logs/app_20260930.log
	logFileName := fmt.Sprintf("app_%s.log", date)
	logPath := filepath.Join(logDir, logFileName)

	file, err := os.OpenFile(
		logPath,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	logger := slog.New(
		slog.NewJSONHandler(file, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}),
	)

	logger.Info("application started")

	// logger.Info(
	// 	"process csv",
	// 	"file", "input.csv",
	// 	"total", 1000,
	// )

	// logger.Error(
	// 	"read csv failed",
	// 	"file", "input.csv",
	// 	"error", "invalid format",
	// )
}

func main() {
	caseReport := flag.String("case", "", "report case")
	_ = caseReport
	flag.Parse()

}
