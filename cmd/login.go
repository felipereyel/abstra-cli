package cmd

import (
	"abstra-cli/config"
	"abstra-cli/utils"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func login(cmd *cobra.Command, args []string) {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		ws, upgradeErr := upgrader.Upgrade(w, r, nil)
		utils.Check(upgradeErr)

		_, msg, readErr := ws.ReadMessage()
		utils.Check(readErr)

		config.CreateConfigFile(msg)
		fmt.Println("Done!")
		os.Exit(0)
	})

	opened := utils.OpenBrowser(config.LoginUrl)
	if !opened {
		fmt.Printf("Could not open browser automatically.\nPlease open %v in your browser.\n", config.LoginUrl)
	}
	http.ListenAndServe(config.Port, nil)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with Abstra Cloud",
	Run:   login,
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
