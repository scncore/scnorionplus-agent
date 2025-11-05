//go:build windows

package logger

import (
	"log"
	"os"
	"path/filepath"
)

func New() *scnorionLogger {
	logger := scnorionLogger{}

	// Get executable path to store logs
	ex, err := os.Executable()
	if err != nil {
		log.Fatalf("could not get executable info: %v", err)
	}
	wd := filepath.Dir(ex)

	logPath := filepath.Join(wd, "logs", "scnorion-log.txt")
	logger.LogFile, err = os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("could not create log file: %v", err)
	}

	log.SetOutput(logger.LogFile)
	log.SetPrefix("scnorion-agent: ")
	log.SetFlags(log.Ldate | log.Ltime)

	return &logger
}
