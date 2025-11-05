//go:build linux

package main

import (
	"github.com/scncore/scnorion-agent/internal/logger"
)

func main() {
	// Instantiate logger
	l := logger.New()

	// Instantiate service
	s := NewService(l)

	s.Execute()
}
