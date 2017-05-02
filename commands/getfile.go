package commands

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/storvik/pcloud-cli/helper"
	"github.com/storvik/pcloud-cli/models"
)

var (
	getfileCmd = &cobra.Command{
		Use:   "get [path to remote folder] [optional location for downloaded file]",
		Short: "Get remote file url and download it.",
		Long: `Get remote file url from server and place it in
specified location, if any. Downloads the file using the best
available server. Paths containing spaces should be wrapped in
double quotes.`,

		Run: getfile,
	}
)

func init() {
	FileCmd.AddCommand(getfileCmd)

	// Hidden / Aliased
}

func getfile(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Invalid input. See 'pcloud-cli copyfile --help'.")
		os.Exit(1)
	}

	parameters := url.Values{}
	if args[0][0] != 47 {
		parameters.Add("path", "/"+args[0])
	} else {
		parameters.Add("path", args[0])
	}

	pcloud := NewPcloud()
	pcloud.Endpoint = "/getfilelink"
	pcloud.Parameters = parameters
	pcloud.AccessToken = AccessToken

	resp, err := pcloud.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var response models.GetfileResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		fmt.Println("Could not decode server response.")
		fmt.Println(err)
		os.Exit(1)
	}

	filepath := filepath.Base(string(response.Path))
	if len(args) > 1 {
		filepath = args[1] + filepath
	}

	fileURL := "http://" + response.Hosts[0] + response.Path

	helper.DownloadFile(fileURL, filepath)

	fmt.Println("File successfully downloaded")

}
