/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: current_bed.go (c) 2021
 * @Created: 2021-06-23 07:43:44
 * @Modified: 2021-06-23 20:45:14
 */

package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(currentBedCmd)
}

var currentBedCmd = &cobra.Command{
	Use:   "current",
	Short: "Show image bed in use",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Image bed in use ==> [ %s ]\n", config.ImageBed)
	},
}
