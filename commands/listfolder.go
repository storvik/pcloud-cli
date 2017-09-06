package commands

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	m "github.com/storvik/pcloud-cli/models"
)

var (
	listfolderCmd = &cobra.Command{
		Use:   "list [path to folder to list]",
		Short: "List folders in pCloud directory",
		Long: `List all folders in given pCloud directory.
If no path is given, the top level directory is listed.
Paths containing spaces should be wrapped in double quotes.`,

		Run: listfolder,
	}

	lsCmd = &cobra.Command{
		Hidden: true,

		Use:   "ls [path to folder to list]",
		Short: "List folders in pCloud directory",
		Long: `List all folders in given pCloud directory.
If no path is given, the top level directory is listed.
Paths containing spaces should be wrapped in double quotes.`,

		Run: listfolder,
	}
)

var (
	showdeleted bool
	nofiles     bool
)

func init() {
	FolderCmd.AddCommand(listfolderCmd)
	listfolderCmd.Flags().BoolVarP(&showdeleted, "showdeleted", "", false, "show deleted files")
	listfolderCmd.Flags().BoolVarP(&nofiles, "nofiles", "", false, "list directories only")

	// Hidden / aliased
	FolderCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolVarP(&showdeleted, "showdeleted", "", false, "show deleted files")
	lsCmd.Flags().BoolVarP(&nofiles, "nofiles", "", false, "list directories only")
}

func listfolder(cmd *cobra.Command, args []string) {
	parameters := url.Values{}
	if len(args) > 0 {
		if args[0][0] != 47 {
			parameters.Add("path", "/"+args[0])
		} else {
			parameters.Add("path", args[0])
		}
	} else {
		parameters.Add("path", "/")
	}
	if nofiles {
		parameters.Add("nofiles", "1")
	}
	if showdeleted {
		parameters.Add("showdeleted", "1")
	}

	pcloud := new(Pcloud)
	pcloud.Endpoint = "/listfolder"
	pcloud.Parameters = parameters
	pcloud.AccessToken = AccessToken

	resp, err := pcloud.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var response m.ListfolderResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		fmt.Println("Could not decode server response.")
		fmt.Println(err)
		os.Exit(1)
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 20, 4, 4, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\n", "Name", "Modified", "ID")
	for i := range response.Metadata.Contents {
		fmt.Fprintf(w, "%s\t%s\t%d\n", response.Metadata.Contents[i].Name, response.Metadata.Contents[i].Modified, response.Metadata.Contents[i].FolderID)
	}
	w.Flush()
}
