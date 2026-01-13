package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/princetheprogrammerbtw/gitsynq/internal/bundle"
	"github.com/princetheprogrammerbtw/gitsynq/internal/config"
	"github.com/princetheprogrammerbtw/gitsynq/internal/ssh"
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
	green.Println("\n📥 Pulling from Remote Server\n")

	// Load config
	cfg, err := config.Load()
	if err != nil {
		red.Printf("❌ Error loading config: %v\n", err)
		os.Exit(1)
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)

	// Step 1: Connect to server
	s.Suffix = fmt.Sprintf(" Connecting to %s...", cfg.Server.Host)
	s.Start()

	client, err := ssh.NewClient(cfg.Server)
	if err != nil {
		s.Stop()
		red.Printf("❌ SSH connection failed: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	s.Stop()
	green.Println("✅ Connected to server")

	// Step 2: Create bundle on server
	s.Suffix = " Creating bundle on server..."
	s.Start()

	timestamp := time.Now().Format("20060102-150405")
	remoteBundleName := fmt.Sprintf("%s-server-%s.bundle", cfg.Project.Name, timestamp)
	remoteRepoPath := filepath.Join(cfg.Server.RemotePath, cfg.Project.Name)
	remoteBundlePath := filepath.Join(cfg.Server.RemotePath, remoteBundleName)

	createBundleScript := fmt.Sprintf(`
		cd "%s" || exit 1
		
		# Create bundle with all refs
		git bundle create "%s" --all
		
		echo "BUNDLE_CREATED"
	`, remoteRepoPath, remoteBundlePath)

	output, err := client.Run(createBundleScript)
	s.Stop()

	if err != nil || !strings.Contains(output, "BUNDLE_CREATED") {
		red.Printf("❌ Failed to create bundle on server: %v\n", err)
		if verbose {
			fmt.Println("Output:", output)
		}
		os.Exit(1)
	}

	green.Println("✅ Bundle created on server")

	// Step 3: Download bundle
	s.Suffix = " Downloading bundle..."
	s.Start()

	localBundlePath := filepath.Join(cfg.Bundle.Directory, remoteBundleName)
	if err := client.Download(remoteBundlePath, localBundlePath); err != nil {
		s.Stop()
		red.Printf("❌ Download failed: %v\n", err)
		os.Exit(1)
	}

	s.Stop()

	info, _ := os.Stat(localBundlePath)
	green.Printf("✅ Downloaded: %s (%s)\n", remoteBundleName, formatBytes(info.Size()))

	// Step 4: Merge bundle into local repo
	s.Suffix = " Merging changes..."
	s.Start()

	if err := bundle.Merge(localBundlePath, cfg.Project.Branch); err != nil {
		s.Stop()
		red.Printf("❌ Merge failed: %v\n", err)
		yellow.Println("💡 You may need to resolve conflicts manually")
		os.Exit(1)
	}

	s.Stop()
	green.Println("✅ Changes merged successfully!")

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
			yellow.Printf("⚠️  Push to origin failed: %v\n", err)
			yellow.Println("💡 Run 'git push origin " + cfg.Project.Branch + "' manually")
		} else {
			s.Stop()
			green.Println("✅ Pushed to origin!")
		}
	}

	// Success!
	printPullSuccess(cfg, autoPush)
}

func printPullSuccess(cfg *config.Config, pushed bool) {
	green.Println("\n" + strings.Repeat("═", 50))
	green.Println("          🎉 PULL SUCCESSFUL! 🎉")
	green.Println(strings.Repeat("═", 50))

	cyan.Println("\n📊 Latest commits:")
	bundle.ShowRecentCommits(5)

	if !pushed {
		yellow.Println("\n💡 Don't forget to push to GitHub:")
		fmt.Println("   git push origin", cfg.Project.Branch)
	}
}
