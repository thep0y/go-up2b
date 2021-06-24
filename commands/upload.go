/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: upload.go (c) 2021
 * @Created:  2021-06-23 20:55:15
 * @Modified: 2021-06-24 08:59:57
 */

package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thep0y/go-logger/log"
	"github.com/thep0y/go-up2b/apis"
)

func init() {
	rootCmd.AddCommand(uploadCmd)
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload multiple images, the maximum is 10",
	Args:  cobra.MaximumNArgs(10),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("At least one image is passed in")
		}
		err = checkImageIsExists(args)
		if err != nil {
			log.Fatal(err)
		}

		client, err := apis.NewImageBedClient(config, configFile)
		if err != nil {
			log.Fatal(err)
		}

		downloadURL, err := client.UploadImages(args)
		if err != nil {
			log.Fatal(err)
		}
		for _, u := range downloadURL {
			fmt.Println(u)
		}
	},
}

func checkImageIsExists(images []string) error {
	for _, image := range images {
		_, err := os.Stat(image)
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf(`"%s" does not exist `, image)
			}
		}
	}
	return nil
}
