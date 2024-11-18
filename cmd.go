package main

import (
	"bytes"
	"os/exec"
)

func RunCmd(args ...string) (string, string, error) {
	tmux, err := exec.LookPath("tmux")
	if err != nil {
		return "", "", err
	}

	cmd := exec.Command(tmux, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	
	err = cmd.Run()
	out, cmdErr := string(stdout.Bytes()), string(stderr.Bytes())

	return out, cmdErr, err
}
