package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say Hello to my little friend",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello there")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
