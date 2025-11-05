//go:build darwin

package report

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"

	scnorion_nats "github.com/scncore/nats"
)

func (r *Report) getApplicationsInfo(debug bool) error {
	var appData SPApplicationsDataType
	r.Applications = []scnorion_nats.Application{}

	if debug {
		log.Println("[DEBUG]: applications info has been requested")
	}

	out, err := exec.Command("system_profiler", "-json", "SPApplicationsDataType").Output()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(out, &appData); err != nil {
		return err
	}

	for _, a := range appData.SPApplicationsDataType {
		if strings.Contains(a.Path, "Library") {
			continue
		}

		app := scnorion_nats.Application{}
		app.Name = a.Name
		app.Version = a.Version
		if len(a.SignedBy) > 0 {
			signer := strings.Split(a.SignedBy[0], ":")
			if len(signer) > 1 {
				app.Publisher = strings.TrimSpace(signer[1])
			} else {
				app.Publisher = a.ObtainedFrom
			}
		} else {
			app.Publisher = a.ObtainedFrom
		}
		r.Applications = append(r.Applications, app)
	}

	log.Println("[INFO]: desktop apps information has been retrieved from package manager")

	return nil
}

type SPApplicationsDataType struct {
	SPApplicationsDataType []ApplicationsDataType `json:"SPApplicationsDataType"`
}

type ApplicationsDataType struct {
	Name         string   `json:"_name"`
	ArchKind     string   `json:"arch_kind"`
	LastModified string   `json:"lastModified"`
	ObtainedFrom string   `json:"obtained_from"`
	Path         string   `json:"path"`
	SignedBy     []string `json:"signed_by"`
	Version      string   `json:"version"`
}
