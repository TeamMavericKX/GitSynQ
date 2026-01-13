package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/princetheprogrammerbtw/gitsynq/internal/config"
	"github.com/princetheprogrammerbtw/gitsynq/internal/ui"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "🎯 Initialize gitsync for current repository",
	Long:  `Initialize gitsync configuration for the current Git repository.`,
	Run:   runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	printBanner()
	green.Println("\n🎯 Initializing GitSync Configuration")

	reader := bufio.NewReader(os.Stdin)
	var projectName, serverIP, username, remotePath string

	// Check if we're in a git repo
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		red.Println("❌ Error: Not a git repository!")
		yellow.Println("💡 Run 'git init' first or navigate to a git repository")
		os.Exit(1)
	}

	// Get project name
	for {
		cyan.Print("📁 Project name: ")
		projectName, _ = reader.ReadString('\n')
		projectName = strings.TrimSpace(projectName)
		if projectName != "" {
			break
		}
		red.Println("❌ Project name cannot be empty")
	}

	// Get server details
	for {
		cyan.Print("🖥️  Server IP/Hostname (e.g., 192.168.12.4): ")
		serverIP, _ = reader.ReadString('\n')
		serverIP = strings.TrimSpace(serverIP)
		if serverIP != "" {
			break
		}
		red.Println("❌ Server IP/Hostname cannot be empty")
	}

	for {
		cyan.Print("👤 Server username (e.g., prince): ")
		username, _ = reader.ReadString('\n')
		username = strings.TrimSpace(username)
		if username != "" {
			break
		}
		red.Println("❌ Username cannot be empty")
	}

	for {
		cyan.Print("📂 Remote project path (e.g., ~/projects): ")
		remotePath, _ = reader.ReadString('\n')
		remotePath = strings.TrimSpace(remotePath)
		if remotePath != "" {
			break
		}
		red.Println("❌ Remote path cannot be empty")
	}

	cyan.Print("🔑 SSH key path (leave empty for default): ")
	sshKeyPath, _ := reader.ReadString('\n')
	sshKeyPath = strings.TrimSpace(sshKeyPath)

	cyan.Print("🌿 Main branch name (default: main): ")
	mainBranch, _ := reader.ReadString('\n')
	mainBranch = strings.TrimSpace(mainBranch)
	if mainBranch == "" {
		mainBranch = "main"
	}

	// Create config
	cfg := config.Config{
		Project: config.ProjectConfig{
			Name:   projectName,
			Branch: mainBranch,
		},
		Server: config.ServerConfig{
			Host:       serverIP,
			User:       username,
			Port:       22,
			RemotePath: remotePath,
			SSHKeyPath: sshKeyPath,
		},
		Bundle: config.BundleConfig{
			Directory:  ".gitsync-bundles",
			Compress:   true,
			MaxHistory: 10,
		},
	}

	// Save config
	if err := config.Save(cfg); err != nil {
		red.Printf("❌ Error saving config: %v\n", err)
		os.Exit(1)
	}

	// Create bundle directory
	os.MkdirAll(".gitsync-bundles", 0755)

	// Add to .gitignore
	addToGitignore()

	green.Println("\n✅ GitSync initialized successfully!")
	cyan.Println("\n📋 Configuration saved to .gitsync.yaml")
	yellow.Println("\n🚀 Next steps:")
	fmt.Println("   1. Run 'gitsync push' to sync repo to server")
	fmt.Println("   2. Work on the server")
	fmt.Println("   3. Run 'gitsync pull' to get changes back")
}

func addToGitignore() {
	entries := []string{
		".gitsync-bundles/",
		"*.bundle",
	}

	f, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	// Read existing content to avoid duplicates
	content, _ := os.ReadFile(".gitignore")
	existing := string(content)

	for _, entry := range entries {
		if !strings.Contains(existing, entry) {
			f.WriteString("\n" + entry)
		}
	}
}
