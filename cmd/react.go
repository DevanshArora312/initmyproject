/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var reactCmd = &cobra.Command{
	Use:   "react",
	Short: "Set up a React Frontend Project",
	Long: `Creates a React Native CLI project with several optional 
	dependencies such as nativewind, navigation packages etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		Program = tea.NewProgram(initialModel([]string{
			"Base Project",
			"Base Project + Packages",
			"Base Project + Packages + Tailwind",
			"Base Project + Redux + Packages",
			"Base Project + Redux + Packages + Tailwind",
			"Base Project + Redux + Packages + Tailwind + MUI",
			"Base Project + Redux + Packages + Tailwind + Andt",
			"Base Project + Redux + Packages + Tailwind + Shadcn"}, "react"))
		if _, err := Program.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(reactCmd)

}
