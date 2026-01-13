package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version = "1.0.0" // default version, can be overridden by ldflags
	cfgFile string
	verbose bool

	// Colors
	cyan    = color.New(color.FgCyan, color.Bold)
	green   = color.New(color.FgGreen, color.Bold)
	red     = color.New(color.FgRed, color.Bold)
	yellow  = color.New(color.FgYellow, color.Bold)
	magenta = color.New(color.FgMagenta, color.Bold)
)

var rootCmd = &cobra.Command{
	Use:   "gitsync",
	Short: "🔄 Sync Git repos between laptop and air-gapped servers",
	Long: `
   ██████╗ ██╗████████╗███████╗██╗   ██╗███╗   ██╗ ██████╗
  ██╔════╝ ██║╚══██╔══╝██╔════╝╚██╗ ██╔╝████╗  ██║██╔════╝
  ██║  ███╗██║   ██║   ███████╗ ╚████╔╝ ██╔██╗ ██║██║     
  ██║   ██║██║   ██║   ╚════██║  ╚██╔╝  ██║╚██╗██║██║     
  ╚██████╔╝██║   ██║   ███████║   ██║   ██║ ╚████║╚██████╗
   ╚═════╝ ╚═╝   ╚═╝   ╚══════╝   ╚═╝   ╚═╝  ╚═══╝ ╚═════╝
                                                          
  Sync your Git repositories with air-gapped servers!
  No internet on server? No problem! 🚀
  
  Created by: PrinceTheProgrammer`,
	Version: version,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default: .gitsync.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// Add subcommands
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(pushCmd)
	rootCmd.AddCommand(pullCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(doctorCmd)
	rootCmd.AddCommand(watchCmd)
	rootCmd.AddCommand(diffCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".gitsync")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if verbose {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}

func printBanner() {
	magenta.Println(`
  ╔═══════════════════════════════════════╗
  ║         🔄 GITSYNC v1.0.0 🔄          ║
  ║   Air-Gapped Git Synchronization      ║
  ╚═══════════════════════════════════════╝`)
}
