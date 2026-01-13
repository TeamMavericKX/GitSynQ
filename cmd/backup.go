package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/princetheprogrammerbtw/gitsynq/internal/config"
	"github.com/princetheprogrammerbtw/gitsynq/internal/ssh"
	"github.com/princetheprogrammerbtw/gitsynq/pkg/utils"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "🛡️  Backup remote repository",
	Long:  `Download a full Git bundle of the remote repository for backup purposes.`, 
	Run:   runBackup,
}

func runBackup(cmd *cobra.Command, args []string) {
	printBanner()
	green.Println("\n🛡️  Backing up Remote Repository\n")

	cfg, err := config.Load()
	if err != nil {
		red.Printf("❌ Error loading config: %v\n", err)
		os.Exit(1)
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " Connecting to server..."
	s.Start()

	client, err := ssh.NewClient(cfg.Server)
	if err != nil {
		s.Stop()
		red.Printf("❌ Connection failed: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	s.Stop()
	green.Println("✅ Connected")

	s.Suffix = " Creating full backup bundle on server..."
	s.Start()

	timestamp := time.Now().Format("20060102-150405")
	backupName := fmt.Sprintf("%s-backup-%s.bundle", cfg.Project.Name, timestamp)
	remoteRepoPath := filepath.Join(cfg.Server.RemotePath, cfg.Project.Name)
	remoteBackupPath := filepath.Join(cfg.Server.RemotePath, backupName)

	_, err = client.Run(fmt.Sprintf("cd %s && git bundle create %s --all", remoteRepoPath, remoteBackupPath))
	s.Stop()

	if err != nil {
		red.Printf("❌ Remote backup failed: %v\n", err)
		os.Exit(1)
	}

	green.Println("✅ Backup bundle created on server")

	s.Suffix = " Downloading backup..."
	s.Start()

	backupDir := "backups"
	os.MkdirAll(backupDir, 0755)
	localBackupPath := filepath.Join(backupDir, backupName)

	if err := client.Download(remoteBackupPath, localBackupPath); err != nil {
		s.Stop()
		red.Printf("❌ Download failed: %v\n", err)
		os.Exit(1)
	}

	s.Stop()
	
	info, _ := os.Stat(localBackupPath)
	green.Printf("✅ Backup saved: %s (%s)\n", localBackupPath, utils.FormatBytes(info.Size()))

	// Cleanup remote
	client.Run(fmt.Sprintf("rm -f %s", remoteBackupPath))
}
