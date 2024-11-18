package tmuxsheets

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const (
	redirectUrl = "http://localhost:8080/oauth2callback"
)

type ClientOpts struct {
	// location of the credentials.json
	CredentialsPath	string
	// location of token.json
	TokenPath 			string
}

func NewClientOpts() *ClientOpts {
	return &ClientOpts{
		CredentialsPath: "./.secrets/credentials.json",
		TokenPath: "./.secrets/token.json",
	}
}

func Setup(ctx context.Context,opts *ClientOpts) (*sheets.Service, error) {
	b, err := os.ReadFile(opts.CredentialsPath)
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, err
	}

	config.RedirectURL = redirectUrl

	client := getClient(config, opts)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return srv, err
	}

	return srv, nil
}

func getClient(config *oauth2.Config, opts *ClientOpts) *http.Client {
	tokFile := opts.TokenPath
	token, err := getTokenFromFile(tokFile) 
	if err != nil {
		token = getTokenFromWeb(config)
		saveToken(opts.TokenPath, token)
	}


	return config.Client(context.Background(), token)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authUrl := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	pr, pw := io.Pipe()

	http.HandleFunc("/oauth2callback", func(w http.ResponseWriter, r *http.Request) {
		handleOAuthCallback(w, r, pw)
	})

	go func() {
		http.ListenAndServe(":8080", nil)
		time.Sleep(2 * time.Second)
	}()


	fmt.Println("Go to the following link in your browser to authorize")

	d := color.New(color.FgBlue, color.Bold)
	d.Printf("%v\n", authUrl)

	scanner := bufio.NewScanner(pr)
	
	var authCode string 
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			authCode = strings.Split(line, "=")[1]
			break
		}
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retreive token from the web: %v", err)
	}

	return tok
}

func getTokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)

	return token, err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
