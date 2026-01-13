package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/princetheprogrammerbtw/gitsynq/internal/config"
	"github.com/princetheprogrammerbtw/gitsynq/internal/ssh"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "🔍 Show what will be synced",
	Long:  `Display the commits and file changes that would be sent to the remote server on the next 'push'.`,
	Run:   runDiff,
}

func runDiff(cmd *cobra.Command, args []string) {
	printBanner()
	
	cfg, err := config.Load()
	if err != nil {
		red.Printf("❌ Error loading config: %v\n", err)
		os.Exit(1)
	}

	cyan.Println("\n🔍 Comparing local branch with remote server...\n")

	// Step 1: Connect to server to get last commit
	client, err := ssh.NewClient(cfg.Server)
	if err != nil {
		red.Printf("❌ Failed to connect to server: %v\n", err)
		yellow.Println("   💡 Showing diff against local origin tracking branch instead.")
		showLocalDiff(cfg.Project.Branch)
		return
	}
	defer client.Close()

	repoPath := fmt.Sprintf("%s/%s", cfg.Server.RemotePath, cfg.Project.Name)
	output, err := client.Run(fmt.Sprintf("cd %s && git rev-parse HEAD", repoPath))
	if err != nil {
		red.Printf("❌ Failed to get remote state: %v\n", err)
		return
	}

	remoteHead := strings.TrimSpace(output)
	
	// Step 2: Show local commits since remoteHead
	fmt.Printf("📊 Commits to push since remote HEAD (%s):\n\n", remoteHead[:7])
	
	logCmd := exec.Command("git", "log", fmt.Sprintf("%s..HEAD", remoteHead), "--oneline", "--graph", "--color")
	logCmd.Stdout = os.Stdout
	logCmd.Run()

	// Step 3: Show file summary
	fmt.Println("\n📝 Files changed:")
	statCmd := exec.Command("git", "diff", remoteHead, "HEAD", "--stat", "--color")
	statCmd.Stdout = os.Stdout
	statCmd.Run()
}

func showLocalDiff(branch string) {
	fmt.Printf("📊 Commits not in origin/%s:\n\n", branch)
	exec.Command("git", "log", fmt.Sprintf("origin/%s..%s", branch, branch), "--oneline", "--graph", "--color").Run()
}
