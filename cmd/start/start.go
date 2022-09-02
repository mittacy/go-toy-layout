package start

import (
	"github.com/mittacy/go-toy-layout/cmd/start/http"
	"github.com/spf13/cobra"
)

// Cmd api server
var Cmd = &cobra.Command{
	Use:   "start",
	Short: "start server command",
	Long:  "start server command",
	Run:   run,
}

func init() {
	Cmd.AddCommand(http.StartCmd)
}

func run(cmd *cobra.Command, args []string) {
}
