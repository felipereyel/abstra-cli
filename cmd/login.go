package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"runtime"

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func openbrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func login(cmd *cobra.Command, args []string) {
	user, _ := user.Current()
	configDir := user.HomeDir + "/.config/abstra/"
	os.MkdirAll(configDir, os.ModePerm)
	userFile := configDir + "user.json"

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade:", err)
			return
		}

		_, msg, _ := ws.ReadMessage()
		if err != nil {
			return
		}

		werr := os.WriteFile(userFile, msg, os.ModePerm)
		check(werr)
		fmt.Printf("Done!")
		os.Exit(0)
	})

	openbrowser("http://localhost:8001/cli-login")
	fmt.Printf("Enter http://localhost:8001/cli-login in your browser if it does not load automatically\n")
	http.ListenAndServe(":6553", nil)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with Abstra Cloud",
	Run:   login,
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
