package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(workerCmd)
}

var goWorkerTemplate = template.Must(template.New("zipper.go-worker").Parse(`package main

import (
	_ "{your/module/path}/onetimeops"

	"github.com/xdevor/zipper/pkg/zipper"
)

func main() {
	zipper.Execute()
}
`))

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Create a one-time-operation worker file",
	Long:  `Create a one-time-operation worker file`,
	Run: func(cmd *cobra.Command, args []string) {
		// assign operation filename
		filename := "zworker.go"
		path := filepath.Join("cmd/zworker", filename)
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			err = errors.New("failed to create, the file '" + path + " is duplicated.")
			fmt.Printf("%s\n", err.Error())
			return
		}

		// create file with operation name
		folder := "cmd/zworker"
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			os.Mkdir(folder, 0755)
		}
		f, err := os.Create(path)
		if err != nil {
			fmt.Printf("failed to create worker file: %s\n", err.Error())
			return
		}
		defer f.Close()

		// write operation template to file
		vars := map[string]string{
			"ModulePath": "{your/module/path}",
		}
		if err := goWorkerTemplate.Execute(f, vars); err != nil {
			fmt.Printf("failed to execute template: %s\n", err.Error())
		}

		fmt.Printf("make a new worker file '%s' successfully\n", filename)
	},
}
