package main

import (
	"encoding/json"
	"os"
)

type SheetsApiOpts struct {
	IsEnabled					bool 	`json:"enabled"`
	SpreadsheetId 	string	`json:"id"`
	WriteRange			string 	`json:"write_range"`
}

type Config struct {
	EnabledRepos 		[]string 				`json:"enabled_repositories"`
	SheetsApiOpts 	*SheetsApiOpts	`json:"sheets_api_options"`
}

var (
	homeDir = os.Getenv("HOME")
	configPath = "/.config/tmux-watcher/config.json"
)


func GetConfig() (Config, error) {
	b, err := os.ReadFile(homeDir + configPath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(b, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
