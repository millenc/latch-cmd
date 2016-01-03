package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

//About command. Outputs version and authorship information.
var AboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Version and authorship information",
	Long:  `Version and authorship information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("latch-cmd - version 1.0 - Mikel Pintor (millen@gmail.com) - more info at: https://github.com/millenc/latch-cmd")
	},
}
