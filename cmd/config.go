package cmd

import (
	"fmt"
	"os"

	"github.com/princetheprogrammerbtw/gitsynq/internal/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "⚙️  Show or edit configuration",
	Long:  `Display or modify the GitSync configuration.`,
	Run:   runConfig,
}

var (
	configShow bool
	configEdit bool
)

func init() {
	configCmd.Flags().BoolVarP(&configShow, "show", "s", true, "Show current configuration")
	configCmd.Flags().BoolVarP(&configEdit, "edit", "e", false, "Edit configuration interactively")
}

func runConfig(cmd *cobra.Command, args []string) {
	printBanner()

	if configEdit {
		runInit(cmd, args)
		return
	}

	cfg, err := config.Load()
	if err != nil {
		red.Printf("❌ Error loading config: %v\n", err)
		os.Exit(1)
	}

	cyan.Println("\n⚙️  Current Configuration\n")

	data, _ := yaml.Marshal(cfg)
	fmt.Println(string(data))
}
