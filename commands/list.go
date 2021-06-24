/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: list.go (c) 2021
 * @Created:  2021-06-24 10:36:00
 * @Modified: 2021-06-24 10:40:13
 */

package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thep0y/go-up2b/models"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available image beds",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Here are all available picture beds:")
		for i, c := range config.AuthData {
			if c != nil {
				fmt.Printf("    - [%d] %s\n", i, models.ImageBedCode(i))
			}
		}
	},
}
