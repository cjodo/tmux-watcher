package main

import (
	"context"
	"fmt"
	"os"
	"tmux-watcher/pkg/tmuxsheets"

	"google.golang.org/api/sheets/v4"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("default behaviour")
	} else {
		command := os.Args[1]
		RunCmd(command)
	}

	ctx := context.Background()

	config, err := GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	var srv *sheets.Service
	if config.SheetsApiOpts.IsEnabled {
		opts := tmuxsheets.NewClientOpts()
		srv, err = tmuxsheets.Setup(ctx, opts)
		if err != nil {
			fmt.Println("error starting service: ", err);
		}
	}

	sessions := GetSessions()

	for _, repo := range config.EnabledRepos {
		for _, sessionName := range sessions {
			if repo == sessionName {
				SetHooks(sessionName)
			}
		}
	}

	MoniterSessions(sessions, config, srv)
}
