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
	checksumfileCmd = &cobra.Command{
		Use:   "checksum [file]",
		Short: "Calculate chacksums of file.",
		Long: `Calculate md5 and sha1 checksums of file.
Paths containing spaces should be wrapped in double quotes.`,

		Run: checksumfile,
	}
)

func init() {
	FileCmd.AddCommand(checksumfileCmd)

	// Hidden / Aliased
}

func checksumfile(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Invalid input. See 'pcloud-cli file checksum --help'.")
		os.Exit(1)
	}

	parameters := url.Values{}
	if args[0][0] != 47 {
		parameters.Add("path", "/"+args[0])
	} else {
		parameters.Add("path", args[0])
	}

	pcloud := new(Pcloud)
	pcloud.Endpoint = "/checksumfile"
	pcloud.Parameters = parameters
	pcloud.AccessToken = AccessToken

	resp, err := pcloud.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var response models.ChecksumfileResponse

	if err := json.Unmarshal(resp, &response); err != nil {
		fmt.Println("Could not decode server response.")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("MD5: " + response.Md)
	fmt.Println("SHA1: " + response.Sha)

	if verbose {
		fmt.Println("Name: " + response.Metadata.Name)
		fmt.Println("Modified: " + response.Metadata.Modified)
		fmt.Println("Size: " + strconv.Itoa(response.Metadata.Size))
		fmt.Println("Content type: " + response.Metadata.ContentType)
	}
}
