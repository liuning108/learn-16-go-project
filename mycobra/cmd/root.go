package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "short Description",
	Long:  "long Description",
	//Args:  cobra.RangeArgs(0, 2),

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("root cmd run begin")
		fmt.Println(
			cmd.Flags().Lookup("viper").Value,
			cmd.Flags().Lookup("author").Value,
			cmd.Flags().Lookup("config").Value,
			cmd.Flags().Lookup("license").Value,
			cmd.Flags().Lookup("source").Value,
		)

		fmt.Println("---------------viper------------")

		fmt.Println(
			viper.GetString("author"),
			viper.GetString("license"),
		)

		fmt.Println("root cmd end begin")
	},
	TraverseChildren: true,
}

func Execute() {
	rootCmd.Execute()
}

var cfgFile string
var userLicense string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().Bool("viper", true, "")
	rootCmd.PersistentFlags().StringP("author", "a", "YOU NAME", "")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "")

	rootCmd.Flags().StringP("source", "s", "", "")
	viper.SetDefault("author", "defaut author")
	viper.SetDefault("license", "defaut license")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("license", rootCmd.PersistentFlags().Lookup("license"))

}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Using config file :", viper.ConfigFileUsed())

}
