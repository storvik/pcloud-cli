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
	renamefileCmd = &cobra.Command{
		Use:   "rename [source file] [destination file]",
		Short: "Rename / Move source file.",
		Long: `Rename / Move file to new location.
Paths containing spaces should be wrapped in double quotes.`,

		Run: renamefile,
	}
)

func init() {
	FileCmd.AddCommand(renamefileCmd)

	// Hidden / Aliased
}

func renamefile(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Invalid input. See 'pcloud-cli file rename --help'.")
		os.Exit(1)
	}

	parameters := url.Values{}
	if args[0][0] != 47 {
		parameters.Add("path", "/"+args[0])
	} else {
		parameters.Add("path", args[0])
	}

	if args[1][0] != 47 {
		parameters.Add("topath", "/"+args[1])
	} else {
		parameters.Add("topath", args[1])
	}

	pcloud := new(Pcloud)
	pcloud.Endpoint = "/renamefile"
	pcloud.Parameters = parameters
	pcloud.AccessToken = AccessToken

	resp, err := pcloud.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var response models.RenamefileResponse

	if err := json.Unmarshal(resp, &response); err != nil {
		fmt.Println("Could not decode server response.")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("File renamed successfully")

	if verbose {
		fmt.Println("Name: " + response.Metadata.Name)
		fmt.Println("Modified: " + response.Metadata.Modified)
		fmt.Println("Size: " + strconv.Itoa(response.Metadata.Size))
		fmt.Println("Content type: " + response.Metadata.ContentType)
	}
}
