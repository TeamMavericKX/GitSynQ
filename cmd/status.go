package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/princetheprogrammerbtw/gitsynq/internal/config"
	"github.com/princetheprogrammerbtw/gitsynq/internal/ssh"
	"github.com/princetheprogrammerbtw/gitsynq/internal/ui"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "📊 Show sync status",
	Long:  `Show the current synchronization status between local and remote.`, 
	Run:   runStatus,
}

func runStatus(cmd *cobra.Command, args []string) {
	printBanner()
	ui.Green.Println("\n📊 Sync Status")

	cfg, err := config.Load()
	if err != nil {
		ui.Red.Printf("❌ Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Local status
	ui.Cyan.Println("═══ LOCAL REPOSITORY ═══")
	printLocalStatus()

	// Remote status
	ui.Cyan.Println("\n═══ REMOTE SERVER ═══")
	printRemoteStatus(cmd.Context(), cfg)

	// Sync recommendation
	ui.Cyan.Println("\n═══ RECOMMENDATION ═══")
	printRecommendation(cfg)
}

func printLocalStatus() {
	// Current branch
	branch, _ := exec.Command("git", "branch", "--show-current").Output()
	fmt.Printf("🌿 Branch: %s", string(branch))

	// Last commit
	commit, _ := exec.Command("git", "log", "-1", "--oneline").Output()
	fmt.Printf("📝 Last commit: %s", string(commit))

	// Uncommitted changes
	status, _ := exec.Command("git", "status", "--porcelain").Output()
	if len(status) > 0 {
		ui.Yellow.Printf("⚠️  Uncommitted changes: %d files\n", len(strings.Split(strings.TrimSpace(string(status)), "\n")))
	} else {
		ui.Green.Println("✅ Working tree clean")
	}

	// Unpushed commits
	unpushed, _ := exec.Command("git", "log", "@{u}..HEAD", "--oneline").Output()
	if len(unpushed) > 0 {
		lines := strings.Split(strings.TrimSpace(string(unpushed)), "\n")
		ui.Yellow.Printf("📤 Unpushed commits: %d\n", len(lines))
	}
}

func printRemoteStatus(ctx context.Context, cfg *config.Config) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " Checking remote server..."
	s.Start()

	client, err := ssh.NewClient(cfg.Server)
	if err != nil {
		s.Stop()
		ui.Red.Printf("❌ Cannot connect: %v\n", err)
		return
	}
	defer client.Close()

	repoPath := fmt.Sprintf("%s/%s", cfg.Server.RemotePath, cfg.Project.Name)

	checkScript := fmt.Sprintf(`
		if [ -d "%s/.git" ]; then
			cd "%s"
			echo "EXISTS:true"
			echo "BRANCH:$(git branch --show-current)"
			echo "COMMIT:$(git log -1 --oneline)"
			
			# Check for uncommitted changes
			if git diff --quiet HEAD 2>/dev/null; then
				echo "CLEAN:true"
			else
				echo "CLEAN:false"
				echo "CHANGES:$(git status --porcelain | wc -l)"
			fi
		else
			echo "EXISTS:false"
		fi
	`, repoPath, repoPath)

	output, err := client.Run(ctx, checkScript)
	s.Stop()

	if err != nil {
		ui.Red.Printf("❌ Error checking remote: %v\n", err)
		return
	}

	lines := strings.Split(output, "\n")
	info := make(map[string]string)
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			info[parts[0]] = strings.TrimSpace(parts[1])
		}
	}

	if info["EXISTS"] == "true" {
		fmt.Printf("🖥️  Server: %s@%s\n", cfg.Server.User, cfg.Server.Host)
		fmt.Printf("📂 Path: %s\n", repoPath)
		fmt.Printf("🌿 Branch: %s\n", info["BRANCH"])
		fmt.Printf("📝 Last commit: %s\n", info["COMMIT"])

		if info["CLEAN"] == "true" {
			ui.Green.Println("✅ Working tree clean")
		} else {
			ui.Yellow.Printf("⚠️  Uncommitted changes: %s files\n", info["CHANGES"])
		}
	} else {
		ui.Yellow.Println("📭 Repository not found on server")
		fmt.Println("💡 Run 'gitsync push --full' to initialize")
	}
}

func printRecommendation(cfg *config.Config) {
	ui.Yellow.Println("💡 Suggested actions:")
	fmt.Println("   • Run 'gitsync push' if you have local changes to sync")
	fmt.Println("   • Run 'gitsync pull' if you worked on the server")
	fmt.Println("   • Run 'gitsync pull --push' to sync and push to GitHub")
}
