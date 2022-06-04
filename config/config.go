package config

import (
	"abstra-cli/utils"
)

const (
	LoginUrl          = "http://localhost:8001/cli-login"
	Port              = ":6553"
	configRelativeDir = "/.config/abstra/"
	configFileName    = "user.json"
)

func CreateConfigFile(content []byte) {
	utils.CreateDir(configRelativeDir)
	utils.WriteFile(configRelativeDir+configFileName, content)
}

// func ReadConfigFile() []byte {
// 	configFullDir := utils.CreateRelativeDir(configRelativeDir)
// 	userFile := configFullDir + configFileName
// 	content, readErr := utils.ReadFile(userFile)
// 	utils.Check(readErr)
// 	return content
// }
