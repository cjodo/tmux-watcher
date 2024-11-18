package main

import (
	"fmt"
	"time"

	"google.golang.org/api/sheets/v4"
)

type Session struct {
	Name 					string
	LastModified  time.Time
}

func SetHooks(session string) {
	args := []string{"set-hook","-t", session, "client-attached", "display-message hi"}

	out, cmdErr, err := RunCmd(args...)
	if cmdErr != "" {
		fmt.Println(cmdErr)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(out)
}

func MoniterSessions(sessions []string, config Config, srv *sheets.Service) {
	if srv != nil {
		fmt.Println("sheets service enabled")
	}
	for {
		out, cmdErr, err := RunCmd("list-sessions", "-F", "#{session_name} #{session_activity}")
		if cmdErr != "" {
			fmt.Println(cmdErr)
			continue
		}

		if err != nil {
			fmt.Println(err)
			continue
		}
		
		UpdateSessions(out, config, srv)

		time.Sleep(5 * time.Minute)
	}
}
