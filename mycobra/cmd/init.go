package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init project",
	Long:  "init project",
	Args: func(cmd *cobra.Command, args []string) error {
		fmt.Println(len(args), "-arg len-", args)

		return cobra.RangeArgs(1, 1)(cmd, args)

	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---------init cmd run begin----------")
		fmt.Println(
			cmd.Flags().Lookup("viper").Value,
			cmd.Flags().Lookup("author").Value,
			cmd.Flags().Lookup("config").Value,
			cmd.Flags().Lookup("license").Value,
			cmd.Parent().Flags().Lookup("source").Value,
		)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
