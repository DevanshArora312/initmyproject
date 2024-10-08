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
	Use:   "mern",
	Short: "Set up a MERN FullStack Project",
	Long: `Creates a MERN Stack project with several optional 
	dependencies and various commonly needed frontend and backend packages like mongoose, react-router-dom etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mern called")
		p := tea.NewProgram(initialModel([]string{"Option 1", "Option 2", "Option 3"}))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(mernCmd)

}
