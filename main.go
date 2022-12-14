package main

import (
	"github.com/mittacy/go-toy-layout/cmd/start"
	"github.com/spf13/cobra"
)

const version = "v0.1.0"

var rootCmd = &cobra.Command{
	Use:     "server",
	Short:   "server: An elegant toolkit for start server.",
	Long:    "server: An elegant toolkit for start server.",
	Version: version,
}

func init() {
	rootCmd.AddCommand(start.Cmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
