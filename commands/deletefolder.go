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
	deletefolderCmd = &cobra.Command{
		Use:   "delete [path to folder]",
		Short: "Delete folder.",
		Long: `Delete folder
The given folder must be empty, if not the -r flag should be set.
Paths containing spaces should be wrapped in double quotes.`,

		Run: deletefolder,
	}
)

var (
	deleterecursive bool
)

func init() {
	FolderCmd.AddCommand(deletefolderCmd)
	deletefolderCmd.Flags().BoolVarP(&deleterecursive, "recursive", "r", false, "perform recursive delete")

	// Hidden / Aliased
}

func deletefolder(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Invalid input. See 'pcloud-cli deletefolder --help'.")
		os.Exit(1)
	}

	parameters := url.Values{}
	if args[0][0] != 47 {
		parameters.Add("path", "/"+args[0])
	} else {
		parameters.Add("path", args[0])
	}

	pcloud := new(Pcloud)
	pcloud.Parameters = parameters
	pcloud.AccessToken = AccessToken

	switch {
	case deleterecursive:
		pcloud.Endpoint = "/deletefolderrecursive"
		resp, err := pcloud.Query()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var response m.DeletefolderRecursiveResponse
		if err := json.Unmarshal(resp, &response); err != nil {
			fmt.Println("Could not decode server response.")
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Successfully deleted folder.")
		if verbose {
			fmt.Println("Deleted files: " + strconv.Itoa(response.DeletedFiles))
			fmt.Println("Deleted folders: " + strconv.Itoa(response.DeletedFolders))
		}
	default:
		pcloud.Endpoint = "/deletefolder"
		resp, err := pcloud.Query()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var response m.DeletefolderResponse
		if err := json.Unmarshal(resp, &response); err != nil {
			fmt.Println("Could not decode server response.")
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Successfully deleted folder.")
		if verbose {
			fmt.Println("Name: " + response.Metadata.Name)
			fmt.Println("Path: " + response.Metadata.Path)
			fmt.Println("Modified: " + response.Metadata.Modified)
			fmt.Println("Folder ID: " + strconv.Itoa(response.Metadata.FolderId))
		}
	}
}
