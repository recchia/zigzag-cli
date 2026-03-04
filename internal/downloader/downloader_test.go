package downloader_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/recchia/zigzag-cli/internal/downloader"
)

func TestDownload(t *testing.T) {
	// Create a temporary directory for test output
	tempDir, err := os.MkdirTemp("", "zigzag-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a mock server
	content := "test build content"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, content)
	}))
	defer server.Close()

	fileName := "testfile.apk"

	finalPath, err := downloader.Download(server.URL, tempDir, fileName)
	if err != nil {
		t.Fatalf("Download failed: %v", err)
	}

	// Verify file existence
	if _, err := os.Stat(finalPath); os.IsNotExist(err) {
		t.Errorf("expected file %s to exist", finalPath)
	}

	// Verify content
	data, err := os.ReadFile(finalPath)
	if err != nil {
		t.Fatalf("failed to read downloaded file: %v", err)
	}
	if string(data) != content {
		t.Errorf("expected content %q, got %q", content, string(data))
	}

	// Verify it's inside a timestamp folder
	parentDir := filepath.Base(filepath.Dir(finalPath))
	if len(parentDir) != 19 { // YYYY-MM-DD_HH-MM-SS is 19 chars
		t.Errorf("expected parent directory name to be a timestamp, got %q", parentDir)
	}
}
