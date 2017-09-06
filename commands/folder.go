package commands

import "github.com/spf13/cobra"

var (
	// FolderCmd declares the sub-command group folder
	FolderCmd = &cobra.Command{
		Use:   "folder",
		Short: "Actions to manage folders.",
		Long: `All commands related to mangage folders.
This includes rename folder, create folder, delete folder, etc.`,

		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
)

func init() {
	RootCmd.AddCommand(FolderCmd)
}
