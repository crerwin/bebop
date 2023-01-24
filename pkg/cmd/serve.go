package cmd

import (
	"github.com/crerwin/bebop/internal/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the Bebop API",
	Long:  "Serve the Bebop API",
	Run:   serveRun,
}

func serveRun(cmd *cobra.Command, args []string) {
	server.Start()
}
