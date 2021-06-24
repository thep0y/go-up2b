/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: root.go (c) 2021
 * @Created: 2021-06-23 07:31:31
 * @Modified: 2021-06-23 20:50:12
 */

package commands

import (
	"os"
	"path"
	"runtime"

	"github.com/mailru/easyjson"
	"github.com/spf13/cobra"
	"github.com/thep0y/go-logger/log"
	"github.com/thep0y/go-up2b/models"
)

var (
	configFile string
	config     models.Config
	err        error
)

func findConfigFile() (string, error) {
	var configFolder string

	platform := runtime.GOOS
	switch platform {
	case "darwin":
		configFolder = path.Join(
			os.Getenv("HOME"),
			".config",
		)
	case "windows":
		configFolder = path.Join(
			os.Getenv("APPDATA"),
		)
	case "linux":
		configFolder = path.Join(
			os.Getenv("HOME"),
			".config",
		)
	default:
		panic("unknown operating system")
	}

	configFolder = path.Join(configFolder, "up2b", "conf")

	err := checkFolder(configFolder)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return path.Join(
		configFolder,
		"conf.up2b.json",
	), nil
}

func checkFolder(p string) error {
	_, err := os.Stat(p)
	if err != nil {
		if os.IsNotExist(err) {
			log.Warnf("[ %s ] not exists", p)
			err = os.MkdirAll(p, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func readConfigFile(configFile string) error {
	content, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Warn("the configuration file does not exist")
		} else {
			return err
		}
	}

	if len(content) == 0 {
		log.Warn("the configuration file is empty, you need to save a image bed configuration information first")
	} else {
		if err := easyjson.Unmarshal(content, &config); err != nil {
			return err
		}
	}
	return nil
}

var rootCmd = &cobra.Command{
	Use:   "up2b",
	Short: "up2b can upload images to the specified image bed",
}

func init() {
	configFile, err = findConfigFile()
	if err != nil {
		panic(err)
	}

	err = readConfigFile(configFile)
	if err != nil {
		panic(err)
	}
}

// Execute xeecutes the commands and parameters
// passed in the terminal
func Execute() {
	rootCmd.Execute()
}
