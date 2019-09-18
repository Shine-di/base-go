package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(info)
}

var info = &cobra.Command{
	Use:   "info",
	Short: "i",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome To me service!")
	},
}
