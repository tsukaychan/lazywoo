package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tsukaychan/lazywoo/internal/algo"
)

func init() {
	algo.Register(rootCmd)
}

var rootCmd = &cobra.Command{
	Use:   "lazywoo",
	Short: "lazywoo people simple life",
	Long:  "lazywoo people simple life",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this is lazywoo tools")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
