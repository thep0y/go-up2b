/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: choose.go
 * @Created:  2021-06-23 08:10:16
 * @Modified: 2021-07-27 15:24:42
 */

package commands

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/mailru/easyjson"
	"github.com/spf13/cobra"
	"github.com/thep0y/go-logger/log"
	"github.com/thep0y/go-up2b/apis"
	"github.com/thep0y/go-up2b/models"
)

// ChooseImageBed uses the specified image bed code
// to select the image bed to be used
func ChooseImageBed(code models.ImageBedCode) error {
	if len(config.AuthData) == 0 {
		err := readConfigFile(configFile)
		if err != nil {
			return err
		}
	}

	if code != apis.IPFS {
		if config.AuthData[code] == nil {
			return fmt.Errorf("the configuration of [ %s ] is empty", code)
		}
	}

	config.ImageBed = code

	newData, err := easyjson.Marshal(config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configFile, newData, 0o644)
	if err != nil {
		return err
	}

	fmt.Printf("switched to [ %s ]\n", code)

	return nil
}

var chooseImgBedCmd = &cobra.Command{
	Use:   "choose",
	Short: "Switch the image bed to be used",
	Long: `
the iamge codes: 
	{0: 'sm.ms', 1: 'imgtu.com', 2: 'gitee.com', 3: 'github.com', 4: 'ipfs'}
	`,
	Example: `You need to pass in the parameter "3" when you want to use "github".
Such as:
	up2b choose 3
	`,
	ValidArgs: []string{"0", "1", "2", "3", "4"},
	Args:      cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("you need to pass in a parameter")
		}
		if len(args) > 1 {
			log.Warn("you passed multiple parameters, but only the first parameter is valid")
		}

		code, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		err = ChooseImageBed(models.ImageBedCode(code))
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(chooseImgBedCmd)
}
