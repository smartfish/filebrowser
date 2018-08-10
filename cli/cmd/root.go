package cmd

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	v "github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "filebrowser",
	Version: "(untracked)",
	Aliases: []string{"serve"},
	Short: "A stylish web-based file manager",
	Long: `File Browser is static binary composed of a golang backend
and a Vue.js frontend to create, edit, copy, move, download your files
easily, everywhere, every time.`,
//	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	checkRootAlias()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.SetVersionTemplate("File Browser {{printf \"version %s\" .Version}}\n")

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (defaults are './.filebrowser[ext]', '$HOME/.filebrowser[ext]' or '/etc/filebrowser/.filebrowser[ext]')")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile == "" {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		v.AddConfigPath(".")
		v.AddConfigPath(home)
		v.AddConfigPath("/etc/filebrowser/")
		v.SetConfigName(".filebrowser")
	} else {
		// Use config file from the flag.
		v.SetConfigFile(cfgFile)
	}

	v.SetEnvPrefix("FB")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(v.ConfigParseError); ok {
			panic(err)
		}
	} else {
		fmt.Println("Using config file:", v.ConfigFileUsed())
	}
}
