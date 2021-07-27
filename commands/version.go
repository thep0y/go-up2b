/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: version.go
 * @Created:  2021-06-23 08:26:49
 * @Modified: 2021-07-27 16:02:21
 */

package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const VERSION string = "0.0.3"

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
