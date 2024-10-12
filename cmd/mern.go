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

// mernCmd represents the mern command
var mernCmd = &cobra.Command{
	Use:   "mern [projectname]",
	Short: "Set up a MERN FullStack Project",
	Long: `Creates a MERN Stack project with several optional 
	dependencies and various commonly needed frontend and backend packages like mongoose, react-router-dom etc.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectname := args[0]
		Program = tea.NewProgram(initialModel([]string{
			"Base Project",
			"Base Project + Packages + Tailwind",
			"Base Project + Redux + Packages + Tailwind",
			"Base Project + Redux + Packages + Tailwind + MUI",
			"Base Project + Redux + Packages + Tailwind + Antd",
			"Base Project + Redux + Packages + Tailwind + Shadcn"}, "mern", projectname))
		if _, err := Program.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(mernCmd)

}
