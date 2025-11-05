package report

import (
	"log"
	"time"

	scnorion_utils "github.com/scncore/utils"
	"gopkg.in/ini.v1"
)

func (r *Report) getUpdateTaskInfo(debug bool) error {
	if debug {
		log.Println("[DEBUG]: update task info has been requested")
	}

	// Open ini file
	configFile := scnorion_utils.GetAgentConfigFile()
	cfg, err := ini.Load(configFile)
	if err != nil {
		return err
	}

	key, err := cfg.Section("Agent").GetKey("UpdaterLastExecutionTime")
	if err != nil {
		return err
	}

	r.AgentReport.LastUpdateTaskExecutionTime, err = time.ParseInLocation("2006-01-02T15:04:05", key.String(), time.Local)
	if err != nil {
		return err
	}

	key, err = cfg.Section("Agent").GetKey("UpdaterLastExecutionStatus")
	if err != nil {
		return err
	}
	r.AgentReport.LastUpdateTaskStatus = key.String()

	key, err = cfg.Section("Agent").GetKey("UpdaterLastExecutionResult")
	if err != nil {
		return err
	}
	r.AgentReport.LastUpdateTaskResult = key.String()

	return nil
}
