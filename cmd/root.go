/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "takenotesctl",
	Short: "takenotesctl is a tool to manage notes from cli",
	Long: `takenotesctl is command line application to quickly take and
		view notes from cli. Notes can be filtered by date and time,
		and can also be imported/exported in a csv file.`,

	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.takenotesctl.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
