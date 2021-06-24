/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: version.go (c) 2021
 * @Created:  2021-06-23 08:26:49
 * @Modified: 2021-06-23 08:30:43
 */

package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const VERSION string = "0.0.1"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the current version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("up2b version: %s\n", VERSION)
	},
}
