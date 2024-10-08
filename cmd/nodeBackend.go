/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// nodeBackendCmd represents the nodeBackend command
var nodeBackendCmd = &cobra.Command{
	Use:   "nodeBackend",
	Short: "Set up a node.js backend Project",
	Long: `Creates a Node.js Backend project with several optional 
	dependencies such as mongodb, axios packages etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nodeBackend called")
	},
}

func init() {
	rootCmd.AddCommand(nodeBackendCmd)

}
