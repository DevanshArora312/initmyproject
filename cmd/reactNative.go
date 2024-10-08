package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var reactNativeCmd = &cobra.Command{
	Use:   "reactNative",
	Short: "Generate a basic configuration React native CLI project",
	Long:  `Creates a React Native CLI project with several optional dependencies such as nativewind, navigation packages etc`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reactNative called")
		p := tea.NewProgram(initialModel([]string{"Bare Project", "RN + NativeWind only", "RN + Navigatons only", "RN + NativeWind + Navigations only", "Fully configured with essential modules", "Fully configured with essential modules + Redux"}))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(reactNativeCmd)
}
