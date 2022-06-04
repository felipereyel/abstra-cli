package utils

import (
	"os/exec"
	"runtime"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func OpenBrowser(url string) bool {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		return false
	}
	return err == nil
}
