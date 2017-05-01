package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of pCloud-cli",
	Long:  `The version number of pCloud-cli tells the user which version is currently being used`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(RootCmd.Use + " " + VERSION)
		fmt.Println("Build: " + BUILD)
		fmt.Println("Build time: " + BUILD_TIME)
	},
}
