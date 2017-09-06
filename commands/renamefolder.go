package commands

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	m "github.com/storvik/pcloud-cli/models"
)

var (
	renamefolderCmd = &cobra.Command{
		Use:   "rename [source folder] [destination folder]",
		Short: "Rename / Move folder.",
		Long: `Rename or move source folder.
If source is to be moved to destination without changing name,
destination must end with /. This to avoid name change.
Paths containing spaces should be wrapped in double quotes.`,

		Run: renamefolder,
	}

	mvCmd = &cobra.Command{
		Hidden: true,

		Use:   "mv [source folder] [destination folder]",
		Short: "Move / Rename folder.",
		Long: `Rename or move source folder.
If source is to be moved to destination without changing name,
destination must end with /. This to avoid name change.
Paths containing spaces should be wrapped in double quotes.`,

		Run: renamefolder,
	}
)

func init() {
	FolderCmd.AddCommand(renamefolderCmd)

	// Hidden / Aliased
	FolderCmd.AddCommand(mvCmd)
}

func renamefolder(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Invalid input. See 'pcloud-cli renamefolder --help'.")
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
	pcloud.Endpoint = "/renamefolder"
	pcloud.Parameters = parameters
	pcloud.AccessToken = AccessToken

	resp, err := pcloud.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var response m.RenamefolderResponse

	if err := json.Unmarshal(resp, &response); err != nil {
		fmt.Println("Could not decode server response.")
		fmt.Println(err)
		os.Exit(1)
	}

	if args[1][len(args[1])-1] != 47 {
		fmt.Println("Folder renamed successfully")
	} else {
		fmt.Println("Folder moved successfully")
	}

	if verbose {
		fmt.Println("Name: " + response.Metadata.Name)
		fmt.Println("Path: " + response.Metadata.Path)
		fmt.Println("Modified: " + response.Metadata.Modified)
		fmt.Println("Folder ID: " + strconv.Itoa(response.Metadata.FolderID))
	}
}
