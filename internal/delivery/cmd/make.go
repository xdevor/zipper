package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(makeCmd)
}

var goOperationTemplate = template.Must(template.New("zipper.go-onetimeops").Parse(`package onetimeops

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(up{{.CamelName}}, down{{.CamelName}})
}

func up{{.CamelName}}(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func down{{.CamelName}}(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
`))

var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Make new one time operation file",
	Long:  `Make new one time operation file`,
	Run: func(cmd *cobra.Command, args []string) {
		// assign operation filename
		version := time.Now().Format("20060102150405")
		name := args[0]
		filename := version + "_" + name + ".go"
		path := filepath.Join("onetimeops", filename)
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			err = errors.New("failed to create, the file '" + path + " is duplicated.")
			fmt.Printf("%s\n", err.Error())
			return
		}

		// create file with operation name
		folder := "onetimeops"
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			os.Mkdir(folder, 0755)
		}
		f, err := os.Create(path)
		if err != nil {
			fmt.Printf("failed to create operation file: %s\n", err.Error())
			return
		}
		defer f.Close()

		// write operation template to file
		vars := map[string]string{
			"Version": version,
			"Name":    name,
		}
		if err := goOperationTemplate.Execute(f, vars); err != nil {
			fmt.Printf("failed to execute template: %s\n", err.Error())
		}

		fmt.Printf("make a new operation file '%s' successfully\n", filename)
	},
}
