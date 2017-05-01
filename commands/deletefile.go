package commands

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/storvik/pcloud-cli/models"
)

var (
	deletefileCmd = &cobra.Command{
		Use:   "delete [path to file]",
		Short: "Delete file.",
		Long: `Delete given file from pCloud storage.
Paths containing spaces should be wrapped in double quotes.`,

		Run: deletefile,
	}
)

func init() {
	FileCmd.AddCommand(deletefileCmd)

	// Hidden / Aliased
}

func deletefile(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Invalid input. See 'pcloud-cli file delete --help'.")
		os.Exit(1)
	}

	parameters := url.Values{}
	if args[0][0] != 47 {
		parameters.Add("path", "/"+args[0])
	} else {
		parameters.Add("path", args[0])
	}

	pcloud := new(Pcloud)
	pcloud.Endpoint = "/deletefile"
	pcloud.Parameters = parameters
	pcloud.AccessToken = ACCESS_TOKEN

	resp, err := pcloud.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var response models.DeletefileResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		fmt.Println("Could not decode server response.")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("File deleted successfully")

	if verbose {
		fmt.Println("Name: " + response.Metadata.Name)
		fmt.Println("Modified: " + response.Metadata.Modified)
		fmt.Println("Size: " + strconv.Itoa(response.Metadata.Size))
		fmt.Println("Content type: " + response.Metadata.ContentType)
	}
}
