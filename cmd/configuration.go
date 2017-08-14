package cmd

import (
	"fmt"
	"sort"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configurationCmd prints current configuration.
var configurationCmd = &cobra.Command{
	Use:   "configuration",
	Short: "print current configuration",
	Long:  `This command loads config file, overrides default settings and prints all config values.`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("Loaded config: %v", viper.ConfigFileUsed())
		keys := viper.AllKeys()
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Printf("%v: %v\n", color.RedString(k), viper.GetString(k))
		}
	},
}

func init() {
	RootCmd.AddCommand(configurationCmd)
}
