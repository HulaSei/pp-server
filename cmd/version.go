package cmd

import (
	"fmt"

	"github.com/perfect-panel/server/internal/constant"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "PPanel version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[PPanel version] " + constant.Display())
	},
}
