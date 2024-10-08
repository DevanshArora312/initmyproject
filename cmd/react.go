/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var reactCmd = &cobra.Command{
	Use:   "react",
	Short: "Set up a React Frontend Project",
	Long: `Creates a React Native CLI project with several optional 
	dependencies such as nativewind, navigation packages etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("react called")
	},
}

func init() {
	rootCmd.AddCommand(reactCmd)

}
