//go:build windows

package main

import (
	"log"
	"time"

	"github.com/scncore/scnorion-agent/internal/agent"
	"github.com/scncore/scnorion-agent/internal/logger"
	"golang.org/x/sys/windows/svc"
)

type scnorionService struct {
	Logger *logger.scnorionLogger
}

func NewService(l *logger.scnorionLogger) *scnorionService {
	return &scnorionService{
		Logger: l,
	}
}

func (s *scnorionService) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.StartPending}
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

	// Get new agent
	a := agent.New()

	// Start agent
	a.Start()

	// service control manager
loop:
	for {
		select {
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				changes <- c.CurrentStatus
				time.Sleep(100 * time.Millisecond)
				changes <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				log.Println("[INFO]: service has received the stop or shutdown command")
				s.Logger.Close()
				a.Stop()
				break loop
			default:
				log.Println("[WARN]: unexpected control request")
				return true, 1
			}
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return true, 0
}
