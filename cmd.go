package main

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/shirou/gopsutil/v3/process"
)

const (
	PROCESS_NAME = "tmux-watcher"
)

func RunTmux(args ...string) (string, string, error) {
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

func RunCmd(command string) error {
	switch command {
	case "stop":
		fmt.Println("stop received")
		processes, err := process.Processes()
		if err != nil {
			fmt.Println(err)
		}
		for _, p := range processes {
			n, err := p.Name()
			if err != nil {
				fmt.Println(err)
			}
			if n == PROCESS_NAME {
				err := p.Kill()
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}

	return nil
}
