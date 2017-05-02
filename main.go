package main

import "github.com/storvik/pcloud-cli/commands"

var (
	CommitHash string
	BuildTime  string

	BaseURL      = "https://api.pcloud.com"
	ClientID     = "wMJTDKXtja"
	ClientSecret = "bCS3k9W89t0zL51qpcL2Ck3bjnF7"
)

func main() {
	commands.Execute(CommitHash, BuildTime, BaseURL, ClientID, ClientSecret)
}
