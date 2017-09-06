package main

import "github.com/storvik/pcloud-cli/commands"

var (
	// CommitHash describing current build
	CommitHash string
	// BuildTime marking time when built
	BuildTime  string

	// BaseURL to pCloud API
	BaseURL      = "https://api.pcloud.com"
	// ClientID is pCloud ID of pcloud-cli
	ClientID     = "wMJTDKXtja"
	// ClientSecret is secret key needed to identify app
	ClientSecret = "bCS3k9W89t0zL51qpcL2Ck3bjnF7"
)

func main() {
	commands.Execute(CommitHash, BuildTime, BaseURL, ClientID, ClientSecret)
}
