package cmd

import (
	"fmt"
	"os"

	"github.com/CatchZeng/localize/configs"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "localize",
	Short: "Localize is a localization tool.",
	Long:  "Localize is a localization tool. Localize supports the mutual conversion of localized files on multiple platforms, including: iOS strings and Excel, android strings.xml and Excel, iOS strings and android strings.xml, etc.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(configs.InitConfig)
}
