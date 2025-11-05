//go:build darwin

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/scncore/scnorion-agent/internal/agent"
	"github.com/scncore/scnorion-agent/internal/logger"
)

type scnorionService struct {
	Logger *logger.scnorionLogger
}

func NewService(l *logger.scnorionLogger) *scnorionService {
	return &scnorionService{
		Logger: l,
	}
}

func (s *scnorionService) Execute() {
	// Get new agent
	a := agent.New()

	// Start agent
	a.Start()

	// Keep the connection alive for service
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-done

	// Stop agent
	log.Println("[INFO]: service has received the stop or shutdown command")
	s.Logger.Close()
	a.Stop()
}
