package cmd

import (
	"fmt"

	"github.com/crerwin/bebop/pkg/bebop"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the Bebop version",
	Long:  "Displays the Bebop version",
	Run:   versionRun,
}

func versionRun(cmd *cobra.Command, args []string) {
	fmt.Printf("Bebop version %v\n", bebop.Version)
}
