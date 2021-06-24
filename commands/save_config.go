/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: save_config.go (c) 2021
 * @Created:  2021-06-23 08:07:09
 * @Modified: 2021-06-24 13:35:11
 */

package commands

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/mailru/easyjson"
	"github.com/spf13/cobra"
	"github.com/thep0y/go-logger/log"
	"github.com/thep0y/go-up2b/apis"
	"github.com/thep0y/go-up2b/models"
)

// SaveConfig saves the configuration information
// of one or more websites to the local configuration file.
//
// If the configuration information of multiple image
// beds is passed in, the first one will be used as the
// default image bed
func SaveConfig(newConfig models.Config) error {
	if len(config.AuthData) == 0 {
		return errors.New("you must pass in at least one image bed configuration information")
	} else if len(config.AuthData) < int(config.ImageBed)+1 {
		return errors.New("the index of the image bed information is incorrect")
	}

	if config.AuthData != [4]*models.LoginInfo{nil, nil, nil, nil} {
		for i, v := range config.AuthData {
			if v != nil {
				config.AuthData[i] = v
			}
		}
		newConfig.AuthData = config.AuthData
	}

	data, err := easyjson.Marshal(config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configFile, data, 0o644)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(saveConfigCmd)
}

var saveConfigCmd = &cobra.Command{
	Use: "save",
	Short: `Save the configuration of one or more image beds,
	      and select the first one as the default image bed`,
	Example: `  One configuration or multiple configurations can be saved, and the image bed corresponding to the first configuration passed in is used as the used image bed by default.
  The format of the configuration is %d %s, such as:

	up2b save 1 "username password" 0 "token" ...
	
  This command will use [ imgtu.com ] as the default image bed.
  The configuration information must be enclosed in double quotation marks, and each field is separated by a space.
  The configuration information format of each image bed is as follows:
	- 0:sm.ms      => "token"
	- 1:imgtu.com  => "username password"
	- 2:gitee.com  => "token username repo folder"
	- 3:github.com => "token username repo folder"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)%2 > 0 {
			log.Fatal("the number of parameters is odd")
		}

		for i := 0; i < len(args); i += 2 {
			code, err := strconv.Atoi(args[i])
			if err != nil {
				log.Fatalf("the image code is invalid [ %s ]", args[i])
			}
			imageCode := models.ImageBedCode(code)
			loginInfo, err := parseArgs(imageCode, args[i+1])
			if err != nil {
				log.Fatal(err)
			}
			if i == 0 {
				config.ImageBed = imageCode
			}
			config.AuthData[imageCode] = loginInfo
		}

		data, err := easyjson.Marshal(config)
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile(configFile, data, 0o644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(strings.Repeat("-", 60))
		fmt.Printf("[ %s ] has been set as the default image bed, \nand the configuration is saved in \"%s\".\n", config.ImageBed, configFile)
	},
}

func parseArgs(code models.ImageBedCode, info string) (*models.LoginInfo, error) {
	args := strings.Fields(info)

	switch code {
	case apis.SMMS:
		if len(args) != 1 {
			return nil, fmt.Errorf("the configuration of `%s` should have 1 parameter1:\ntoken\nbut you passed in %d parameters", code, len(args))
		}
		return &models.LoginInfo{
			Token: args[0],
		}, nil
	case apis.IMGTU:
		if len(args) != 2 {
			return nil, fmt.Errorf("the configuration of `%s` should have 2 parameters:\nusername, password\nbut you passed in %d parameters", code, len(args))
		}
		imgtu := apis.Imgtu{
			Config: &models.LoginInfo{
				Username: args[0],
				Password: args[1],
			},
		}

		imgtu.MakeHeaders(nil)
		imgtu.NewRequest()
		return imgtu.Login(*imgtu.Config)
	case apis.GITEE:
		if len(args) != 4 {
			return nil, fmt.Errorf("the configuration of `%s` should have 4 parameters:\ntoken, username, repo, folder,\\nbut you passed in %d parameters", code, len(args))
		}
		return &models.LoginInfo{
			Token:    args[0],
			Username: args[1],
			Repo:     args[2],
			Folder:   args[3],
		}, nil
	case apis.GITHUB:
		if len(args) != 4 {
			return nil, fmt.Errorf("the configuration of `%s` should have 4 parameters:\ntoken, username, repo, folder\nbut you passed in %d parameters", code, len(args))
		}
		return &models.LoginInfo{
			Token:    args[0],
			Username: args[1],
			Repo:     args[2],
			Folder:   args[3],
		}, nil
	default:
		return nil, fmt.Errorf("unkown image bed code [ %d ]", code)
	}
}
