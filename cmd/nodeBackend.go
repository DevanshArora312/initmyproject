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

// nodeBackendCmd represents the nodeBackend command
var nodeBackendCmd = &cobra.Command{
	Use:   "nodeBackend",
	Short: "Set up a node.js + express backend Project",
	Long: `Creates a Node.js + express Backend project with several optional 
	dependencies such as mongodb, axios packages etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		mod := initialModel([]string{
			"Base Project",
			"Base Project + Packages",
			"Base Project + Mongoose + Packages",
			"Base Project + Mongoose + Packages + Files & Folder"}, "nodeBackend")
		Program = tea.NewProgram(mod)
		if _, err := Program.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
		// mod.listenForUpdates()
	},
}

func init() {
	rootCmd.AddCommand(nodeBackendCmd)

}
