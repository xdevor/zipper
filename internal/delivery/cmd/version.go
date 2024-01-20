package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xdevor/zipper/internal/config"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Zipper",
	Long:  `All software has versions. This is Zipper's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Zipper current version is ", config.App.Version)
	},
}
