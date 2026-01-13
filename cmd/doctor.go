package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/princetheprogrammerbtw/gitsynq/internal/config"
	"github.com/princetheprogrammerbtw/gitsynq/internal/ssh"
	"github.com/princetheprogrammerbtw/gitsynq/internal/ui"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "🩺 Diagnose common issues",
	Long:  `Check your local and remote environment for common configuration and connection problems.`,
	Run:   runDoctor,
}

func runDoctor(cmd *cobra.Command, args []string) {
	printBanner()
	fmt.Println("\n🩺 Running GitSynq Doctor...")

	// 1. Check Git
	fmt.Print("🔍 Checking local Git... ")
	if err := exec.Command("git", "--version").Run(); err != nil {
		ui.Red.Println("❌ Not found! Please install Git.")
	} else {
		ui.Green.Println("✅ OK")
	}

	// 2. Check Config
	fmt.Print("🔍 Checking configuration file... ")
	cfg, err := config.Load()
	if err != nil {
		ui.Red.Printf("❌ Failed: %v\n", err)
		ui.Yellow.Println("   💡 Run 'gitsync init' to create one.")
		os.Exit(1)
	}
	ui.Green.Println("✅ OK")

	// 3. Check SSH Connection
	fmt.Printf("🔍 Testing SSH connection to %s... ", cfg.Server.Host)
	client, err := ssh.NewClient(cfg.Server)
	if err != nil {
		ui.Red.Printf("❌ Failed: %v\n", err)
	} else {
		ui.Green.Println("✅ OK")
		defer client.Close()

		// 4. Check Remote Git
		fmt.Print("🔍 Checking Git on remote server... ")
		output, err := client.Run("git --version")
		if err != nil {
			ui.Red.Printf("❌ Failed: %v\n", err)
			fmt.Println("   Output:", output)
		} else {
			ui.Green.Println("✅ OK")
		}

		// 5. Check Remote Path
		fmt.Printf("🔍 Checking remote path %s... ", cfg.Server.RemotePath)
		_, err = client.Run(fmt.Sprintf("mkdir -p %s", cfg.Server.RemotePath))
		if err != nil {
			ui.Red.Printf("❌ Failed: %v\n", err)
		} else {
			ui.Green.Println("✅ OK")
		}
	}

	fmt.Println("\n✅ Doctor check complete!")
}
