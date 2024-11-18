package main

import (
	"fmt"
	"log"
	"strings"

	"google.golang.org/api/sheets/v4"
)

// TODO is this really how I should do it lol

type Session struct {
	SessionName string `json:"name"`
	StartTime		string `json:"start_time"`
	LastActive	int			`json:"last_active"`
}

type Sessions struct {
	Sessions []Session 
}

func GetSessions() []string {
	var args = []string{"list-sessions", "-F#S"}

	out, cmdErr, err := RunCmd(args...)

	if cmdErr != "" {
		fmt.Println(cmdErr)
	}

	if(err != nil) {
		fmt.Println(err)
		return nil
	}

	sessions := strings.Split(out, "\n")

	return sessions
}

func UpdateSessionsLocally(cmdOut string, config Config) {
	// activeSessions := strings.Split(cmdOut, "\n")
	// repoInteractions := make(map[string][]interface{})

	f, err := getWriteFile(config)
	if err != nil {
		log.Fatal("can't get write file. exiting ", err)
	}
	
	fmt.Println(f)
}

func UpdateSessionsWithGoogle(cmdOut string, config Config, srv *sheets.Service){
	activeSessions := strings.Split(cmdOut, "\n")
	repoInteractions := make(map[string][]interface{})

	for _, session := range activeSessions {
		sessionParts := strings.Fields(session)

		if len(sessionParts) < 2 {
			continue
		}

		sessionName := sessionParts[0]
		lastInteracted := sessionParts[1]

		for _, repo := range config.EnabledRepos {
			if repo == sessionName {
				repoInteractions[repo] = append(repoInteractions[repo], lastInteracted)
			}
		}
	}

	readRange := "Sheet1!1:1"
	resp, err := srv.Spreadsheets.Values.Get(config.SheetsApiOpts.SpreadsheetId, readRange).Do()
	if err != nil {
		fmt.Println("error reading from sheets: ", err)
		return
	}

	existingHeader := []string{}
	if len(resp.Values) > 0 {
		for _, value := range resp.Values[0] {
			existingHeader = append(existingHeader, fmt.Sprintf("%v", value))
		}
	}

	// update header if enabled_repos change
	newHeader := existingHeader
	for _, repo := range config.EnabledRepos {
		if !contains(existingHeader, repo) {
			newHeader = append(newHeader, repo)
		}
	}

	if len(existingHeader) < len(newHeader) || !slicesEqual(existingHeader, newHeader) {
		var vrHeader sheets.ValueRange
		vrHeader.Values = append(vrHeader.Values, convertToInterfaceSlice(newHeader))

		_, err := srv.Spreadsheets.Values.Update(
			config.SheetsApiOpts.SpreadsheetId,
			"Sheet1!1:1",
			&vrHeader,
			).ValueInputOption("RAW").Do()

		if err != nil {
			fmt.Println("error appending header: ", err)
			return
		}
	}

	// prevents overwriting existing data
	columnIndex := make(map[string]int)
	for i, repo := range newHeader {
		columnIndex[repo] = i
	}

	maxRows := 0
	for _, times := range repoInteractions {
		if len(times) > maxRows {
			maxRows = len(times)
		}
	}

	var vrData sheets.ValueRange
	for i := 0; i < maxRows; i++ {
		row := make([]interface{}, len(newHeader)) // Pre-fill row with empty cells
		for repo, times := range repoInteractions {
			if i < len(times) {
				row[columnIndex[repo]] = times[i]
			}
		}
		vrData.Values = append(vrData.Values, row)
	}

	if len(vrData.Values) > 0 {
		_, err := srv.Spreadsheets.Values.Append(
			config.SheetsApiOpts.SpreadsheetId,
			"Sheet1!A2",
			&vrData,
			).ValueInputOption("RAW").Do()
		if err != nil {
			fmt.Println("error appending: ", err)
		} 
	}
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func slicesEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}


// Helper function to convert a slice of strings to a slice of interfaces
func convertToInterfaceSlice(slice []string) []interface{} {
	interfaceSlice := make([]interface{}, len(slice))
	for i, v := range slice {
		interfaceSlice[i] = v
	}
	return interfaceSlice
}
