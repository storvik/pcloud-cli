package main

import "github.com/storvik/pcloud-cli/commands"

var (
	VERSION    string
	BUILD      string
	BUILD_TIME string

	BASE_URL      = "https://api.pcloud.com"
	CLIENT_ID     = "wMJTDKXtja"
	CLIENT_SECRET = "bCS3k9W89t0zL51qpcL2Ck3bjnF7"
)

func main() {
	commands.Execute(VERSION, BUILD, BUILD_TIME, BASE_URL, CLIENT_ID, CLIENT_SECRET)
}
