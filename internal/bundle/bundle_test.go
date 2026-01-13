package bundle

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func setupTestRepo(t *testing.T) string {
	dir := t.TempDir()
	
	runCmd := func(name string, args ...string) {
		cmd := exec.Command(name, args...)
		cmd.Dir = dir
		if err := cmd.Run(); err != nil {
			t.Fatalf("%s %v failed: %v", name, args, err)
		}
	}

	runCmd("git", "init")
	runCmd("git", "config", "user.email", "test@example.com")
	runCmd("git", "config", "user.name", "Test User")
	
	err := os.WriteFile(filepath.Join(dir, "file.txt"), []byte("hello"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	
	runCmd("git", "add", "file.txt")
	runCmd("git", "commit", "-m", "initial commit")
	
	return dir
}

func TestCreateFull(t *testing.T) {
	repoDir := setupTestRepo(t)
	bundlePath := filepath.Join(t.TempDir(), "test.bundle")
	
	// Change to repo dir
	oldWd, _ := os.Getwd()
	defer os.Chdir(oldWd)
	os.Chdir(repoDir)
	
	err := CreateFull(bundlePath)
	if err != nil {
		t.Fatalf("CreateFull failed: %v", err)
	}
	
	if _, err := os.Stat(bundlePath); os.IsNotExist(err) {
		t.Error("Bundle file not created")
	}
}

func TestMerge(t *testing.T) {
	repoDir := setupTestRepo(t)
	bundlePath := filepath.Join(t.TempDir(), "test.bundle")
	
	oldWd, _ := os.Getwd()
	defer os.Chdir(oldWd)
	os.Chdir(repoDir)
	
	// Create bundle
	CreateFull(bundlePath)
	
	// Create a new repo and merge the bundle
	newRepoDir := t.TempDir()
	os.Chdir(newRepoDir)
	exec.Command("git", "init").Run()
	exec.Command("git", "config", "user.email", "test@example.com").Run()
	exec.Command("git", "config", "user.name", "Test User").Run()
	
	// We need a commit to merge into usually, or it's a clone
	// But Merge expects an existing repo
	err := Merge(bundlePath, "master") // setupTestRepo uses master/main depending on git version
	if err != nil {
		// Try main
		err = Merge(bundlePath, "main")
	}
	
	if err != nil {
		t.Fatalf("Merge failed: %v", err)
	}
	
	if _, err := os.Stat("file.txt"); os.IsNotExist(err) {
		t.Error("file.txt not found after merge")
	}
}
