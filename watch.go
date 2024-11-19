package main

import (
	"fmt"
	"time"

	"google.golang.org/api/sheets/v4"
)

func SetHooks(session string) {
	// not yet used
	fmt.Println("setting hooks")
	args := []string{"set-hook","-t", session, "client-attached", "display-message hi"}

	out, cmdErr, err := RunTmux(args...)
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
	withGoogle := false
	if srv != nil {
		withGoogle = true
	}

	var prevOut string = ""
	for {
		out, cmdErr, err := RunTmux("list-sessions", "-F", "#{session_name} #{session_activity}")

		if cmdErr != "" {
			fmt.Println(cmdErr)
			continue
		}

		if err != nil {
			fmt.Println(err)
			continue
		}
		
		if withGoogle {
			UpdateSessionsWithGoogle(out, config, srv)
		} else {
			UpdateSessionsLocally(config, out, prevOut)
		}

		prevOut = out

		time.Sleep(5 * time.Minute)
	}
}
