package cmd

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"syscall"
)

func Run(command string) (*ResultCmd, error) {

	log.Debugf("Cmd: sh -c %s", command)

	cmd := exec.Command("sh", "-c", command)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	resultCmd := &ResultCmd{}

	if exitError, ok := err.(*exec.ExitError); ok {
		waitStatus := exitError.Sys().(syscall.WaitStatus)
		resultCmd.ExitCode = waitStatus.ExitStatus()
	}

	resultCmd.Stdout = string(cmdOutput.Bytes())

	log.Debugf("Result: %s", resultCmd)

	return resultCmd, nil
}
