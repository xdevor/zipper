package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	_ "github.com/xdevor/zipper/onetimeops"
	"github.com/xdevor/zipper/pkg/zipper"
)

var rootCmd = &cobra.Command{
	Use:   "zipper",
	Short: "Zipper is a one time operation tool allow you run task after deployment, just like migrations",
	Long:  `Zipper is a one time operation tool allow you run task after deployment, just like migrations in Go. Complete documentation is available at https://github.com/xdevor/zipper`,
	Run: func(cmd *cobra.Command, args []string) {
		zipper.Execute()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
