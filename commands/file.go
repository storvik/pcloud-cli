package commands

import "github.com/spf13/cobra"

var (
	// FileCmd declares the sub-command group file
	FileCmd = &cobra.Command{
		Use:   "file",
		Short: "Actions to manage files.",
		Long: `All commands related to mangage files.
This includes rename file, upload file, download file, etc.`,

		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
)

func init() {
	RootCmd.AddCommand(FileCmd)
}
