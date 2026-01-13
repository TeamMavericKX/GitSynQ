package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/princetheprogrammerbtw/gitsynq/internal/config"
	"github.com/princetheprogrammerbtw/gitsynq/pkg/utils"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "🕒 Show sync history",
	Long:  `List the local Git bundles that have been created for synchronization.`, 
	Run:   runHistory,
}

func runHistory(cmd *cobra.Command, args []string) {
	printBanner()

	cfg, err := config.Load()
	if err != nil {
		red.Printf("❌ Error loading config: %v\n", err)
		os.Exit(1)
	}

	cyan.Println("\n🕒 Synchronization History (Local Bundles)\n")

	files, err := os.ReadDir(cfg.Bundle.Directory)
	if err != nil {
		if os.IsNotExist(err) {
			yellow.Println("📭 No history found. Perform a 'push' or 'pull' first.")
			return
		}
		red.Printf("❌ Failed to read bundle directory: %v\n", err)
		os.Exit(1)
	}

	if len(files) == 0 {
		yellow.Println("📭 No bundles found in", cfg.Bundle.Directory)
		return
	}

	// Sort files by name (which includes timestamp) descending
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() > files[j].Name()
	})

	fmt.Printf("%-35s %-12s %-20s\n", "BUNDLE NAME", "SIZE", "CREATED")
	fmt.Println(strings.Repeat("-", 70))

	for _, f := range files {
		if f.IsDir() || filepath.Ext(f.Name()) != ".bundle" {
			continue
		}

		info, _ := f.Info()
		fmt.Printf("%-35s %-12s %-20s\n", 
			f.Name(), 
			utils.FormatBytes(info.Size()), 
			info.ModTime().Format("2006-01-02 15:04:05"))
	}
}
