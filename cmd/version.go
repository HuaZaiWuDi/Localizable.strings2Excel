package cmd

import (
	"log"

	"github.com/CatchZeng/gutils/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "localize version",
	Long:  `localize version`,
	Run:   runVersionCmd,
}

func runVersionCmd(_ *cobra.Command, _ []string) {
	value := version.Stringify("2.0.0", "2020/08/17")
	log.Println(value)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
