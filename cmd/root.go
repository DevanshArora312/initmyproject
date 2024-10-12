/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "initmyproject",
	Short: "A CLI tool to aid developers in initialising their projects",
	Long: `A CLI tool built in Golang, Cobra and Bubble tea to aid developers
in initialising and bootstrapping their projects with commonly used 
packages and files.
	
This CLI also removes the need to write the boilerplate code such as 
redux setup to allow the user to focus on creating awesome applications.
	`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.initmyproject.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
