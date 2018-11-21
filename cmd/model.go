package cmd

import (
	"encoding/json"
)

type ResultCmd struct {
	ExitCode int    `json:"exitCode"`
	Stdout   string `json:"stdout,omitempty"`
}

func (r *ResultCmd) String() string {
	json, _ := json.Marshal(r)
	return string(json)
}
