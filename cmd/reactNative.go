package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var reactNativeCmd = &cobra.Command{
	Use:   "reactNative [projectname]",
	Short: "Generate a React native CLI project",
	Long:  `Creates a React Native CLI project with several optional dependencies such as nativewind, navigation packages etc`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("reactNative called")
		projectname := args[0]
		Program = tea.NewProgram(initialModel([]string{
			"Base Project",
			"RN + NativeWind only",
			"RN + NativeWind + Navigations only",
			"Fully configured with essential modules",
			"Fully configured with essential modules + Redux"}, "reactNative", projectname))
		if _, err := Program.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(reactNativeCmd)
}
