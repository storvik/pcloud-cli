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
	copyfileCmd = &cobra.Command{
		Use:   "copy [source file] [destination file]",
		Short: "Copy file to another location.",
		Long: `Copy source to destination.
If destination file exists it will NOT overwrite it unless the
--overwrite flag i set. Paths containing spaces should be wrapped
in double quotes.`,

		Run: copyfile,
	}
)

var (
	overwrite bool
)

func init() {
	FileCmd.AddCommand(copyfileCmd)
	copyfileCmd.Flags().BoolVarP(&overwrite, "overwrite", "o", false, "overwrite if file exists")

	// Hidden / Aliased
}

func copyfile(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Invalid input. See 'pcloud-cli copyfile --help'.")
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

	if !overwrite {
		parameters.Add("noover", strconv.Itoa(1))
	}

	pcloud := new(Pcloud)
	pcloud.Endpoint = "/copyfile"
	pcloud.Parameters = parameters
	pcloud.AccessToken = ACCESS_TOKEN

	resp, err := pcloud.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var response models.CopyfileResponse

	if err := json.Unmarshal(resp, &response); err != nil {
		fmt.Println("Could not decode server response.")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("File copied successfully")

	if verbose {
		fmt.Println("Name: " + response.Metadata.Name)
		fmt.Println("Modified: " + response.Metadata.Modified)
		fmt.Println("Size: " + strconv.Itoa(response.Metadata.Size))
		fmt.Println("Content type: " + response.Metadata.ContentType)
	}
}
