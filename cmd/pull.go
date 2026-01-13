package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/schollz/progressbar/v3"
	"github.com/princetheprogrammerbtw/gitsynq/internal/bundle"
	"github.com/princetheprogrammerbtw/gitsynq/internal/config"
	"github.com/princetheprogrammerbtw/gitsynq/internal/ssh"
	"github.com/princetheprogrammerbtw/gitsynq/internal/ui"
	"github.com/princetheprogrammerbtw/gitsynq/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	autoPush bool
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "📥 Pull changes from remote server",
	Long: `Fetch the latest changes from the remote server and merge them.
	
Examples:
  gitsync pull           # Pull changes from server
  gitsync pull --push    # Pull and automatically push to GitHub`,
	Run: runPull,
}

func init() {
	pullCmd.Flags().BoolVarP(&autoPush, "push", "p", false, "Automatically push to origin after pulling")
}

func runPull(cmd *cobra.Command, args []string) {
	printBanner()
	ui.Green.Println("\n📥 Pulling from Remote Server")

	// Load config
	cfg, err := config.Load()
	if err != nil {
		ui.Red.Printf("❌ Error loading config: %v\n", err)
		os.Exit(1)
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)

	// Step 1: Connect to server
	s.Suffix = fmt.Sprintf(" Connecting to %s...", cfg.Server.Host)
	s.Start()

	client, err := ssh.NewClient(cfg.Server)
	if err != nil {
		s.Stop()
		ui.Red.Printf("❌ SSH connection failed: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	s.Stop()
	ui.Green.Println("✅ Connected to server")

	// Step 2: Create bundle on server
	s.Suffix = " Creating bundle on server..."
	s.Start()

	timestamp := time.Now().Format("20060102-150405")
	remoteBundleName := fmt.Sprintf("%s-server-%s.bundle", cfg.Project.Name, timestamp)
	remoteRepoPath := filepath.Join(cfg.Server.RemotePath, cfg.Project.Name)
	remoteBundlePath := filepath.Join(cfg.Server.RemotePath, remoteBundleName)

	createBundleScript := fmt.Sprintf(`
		cd "%s" || exit 1
		
		# Check for uncommitted changes
		if ! git diff --quiet HEAD 2>/dev/null; then
			echo "UNCOMMITTED_CHANGES"
		fi
		
		# Create bundle with all refs
		git bundle create "%s" --all
		
		echo "BUNDLE_CREATED"
	`, remoteRepoPath, remoteBundlePath)

	output, err := client.Run(createBundleScript)
	s.Stop()

	if strings.Contains(output, "UNCOMMITTED_CHANGES") {
		ui.Yellow.Println("⚠️  Warning: Uncommitted changes exist on the remote server.")
		ui.Yellow.Println("   These changes will NOT be included in the sync until you commit them on the server.")
	}

	if err != nil || !strings.Contains(output, "BUNDLE_CREATED") {
		ui.Red.Printf("❌ Failed to create bundle on server: %v\n", err)
		if verbose {
			fmt.Println("Output:", output)
		}
		os.Exit(1)
	}

	ui.Green.Println("✅ Bundle created on server")

	// Step 3: Download bundle
	s.Suffix = " Downloading bundle..."
	s.Start()

	localBundlePath := filepath.Join(cfg.Bundle.Directory, remoteBundleName)

	bar := progressbar.DefaultBytes(
		-1, // We'll set the total once the transfer starts and we have the size
		"🚀 Downloading bundle",
	)

	err = client.Download(remoteBundlePath, localBundlePath, func(current, total int64) {
		if bar.GetMax() == -1 && total > 0 {
			bar.ChangeMax64(total)
		}
		bar.Set64(current)
	})

	if err != nil {
		ui.Red.Printf("❌ Download failed: %v\n", err)
		os.Exit(1)
	}

	info, _ := os.Stat(localBundlePath)
	ui.Green.Printf("\n✅ Downloaded: %s (%s)\n", remoteBundleName, utils.FormatBytes(info.Size()))

	// Step 4: Merge bundle into local repo
	s.Suffix = " Merging changes..."
	s.Start()

	if err := bundle.Merge(localBundlePath, cfg.Project.Branch); err != nil {
		s.Stop()
		ui.Red.Printf("❌ Merge failed: %v\n", err)
		ui.Yellow.Println("💡 You may need to resolve conflicts manually")
		os.Exit(1)
	}

	s.Stop()
	ui.Green.Println("✅ Changes merged successfully!")

	// Step 5: Cleanup remote bundle
	s.Suffix = " Cleaning up..."
	s.Start()
	client.Run(fmt.Sprintf("rm -f '%s'", remoteBundlePath))
	s.Stop()

	// Step 6: Auto-push to origin (if requested)
	if autoPush {
		s.Suffix = " Pushing to origin..."
		s.Start()

		if err := bundle.PushToOrigin(cfg.Project.Branch); err != nil {
			s.Stop()
			ui.Yellow.Printf("⚠️  Push to origin failed: %v\n", err)
			ui.Yellow.Println("💡 Run 'git push origin " + cfg.Project.Branch + "' manually")
		} else {
			s.Stop()
			ui.Green.Println("✅ Pushed to origin!")
		}
	}

	// Success!
	printPullSuccess(cfg, autoPush)
}

func printPullSuccess(cfg *config.Config, pushed bool) {
	ui.Green.Println("\n" + strings.Repeat("═", 50))
	ui.Green.Println("          🎉 PULL SUCCESSFUL! 🎉")
	ui.Green.Println(strings.Repeat("═", 50))

	ui.Cyan.Println("\n📊 Latest commits:")
	bundle.ShowRecentCommits(5)

	if !pushed {
		ui.Yellow.Println("\n💡 Don't forget to push to GitHub:")
		fmt.Println("   git push origin", cfg.Project.Branch)
	}
}
