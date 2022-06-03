package cmd

import (
	"github.com/spf13/cobra"
)

func sync(cmd *cobra.Command, args []string) {
	// read and user config from $HOME/.config/abstra/user.json
	// update user credentials and persist them
	// get urls from api and download files
	// get upload urls from api and upload files
	// watch for changes in files and upload them
}

var syncCmd = &cobra.Command{
	Use:   "sync [localDir] [workspaceId]",
	Short: "Sync Directory with remote Workspace",
	Run:   sync,
	Args:  cobra.ExactArgs(2),
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
