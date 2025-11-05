//go:build windows

package main

import (
	"log"

	"github.com/scncore/scnorion-agent/internal/logger"
	"golang.org/x/sys/windows/svc"
)

func main() {
	// Instantiate logger
	l := logger.New()

	// Instantiate service
	s := NewService(l)

	// Run service
	err := svc.Run("scnorion-agent", s)
	if err != nil {
		log.Fatalf("could not run service: %v", err)
	}
}
