package git

import (
	"os"
	"os/exec"
	"testing"
)

func TestGetDiff(t *testing.T) {
	// Geçici bir test ortamı oluştur
	tempDir, err := os.MkdirTemp("", "test-repo")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Geçici dizinde bir Git deposu oluştur
	cmd := exec.Command("git", "init")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to initialize git repo: %v", err)
	}

	// Geçici dosya oluştur ve stage et
	filePath := tempDir + "/test.txt"
	if err := os.WriteFile(filePath, []byte("Test content"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	cmd = exec.Command("git", "add", "test.txt")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to stage file: %v", err)
	}

	// Dosyada değişiklik yap ve tekrar stage et
	if err := os.WriteFile(filePath, []byte("Updated content"), 0644); err != nil {
		t.Fatalf("Failed to update file: %v", err)
	}

	cmd = exec.Command("git", "add", "test.txt")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to stage updated file: %v", err)
	}

	// Git diff çıktısını al ve kontrol et
	cmd = exec.Command("git", "diff", "--cached")
	cmd.Dir = tempDir
	output, err := cmd.Output()

	if err != nil {
		t.Fatalf("Failed to get git diff: %v", err)
	}

	diff := string(output)
	if diff == "" {
		t.Error("Expected diff output, but got empty string")
	} else {
		t.Logf("Git diff output:\n%s", diff)
	}
}

func TestGetDiff_NoGitRepo(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "no-git-repo")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	cmd := exec.Command("git", "diff", "--cached")
	cmd.Dir = tempDir
	output, err := cmd.Output()

	if err == nil {
		t.Error("Expected error when running git diff outside a repo, but got none")
	}

	if len(output) > 0 {
		t.Errorf("Expected no output, but got: %s", string(output))
	}
}
