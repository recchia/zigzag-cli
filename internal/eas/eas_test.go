package eas_test

import (
	"encoding/json"
	"testing"

	"github.com/recchia/zigzag-cli/internal/eas"
)

type MockCommandRunner struct {
	Output []byte
	Err    error
}

func (m MockCommandRunner) Run(name string, args ...string) ([]byte, error) {
	return m.Output, m.Err
}

func TestListBuilds(t *testing.T) {
	expectedBuilds := []eas.Build{
		{ID: "1", Status: "finished", Platform: "android"},
		{ID: "2", Status: "finished", Platform: "ios"},
	}
	output, _ := json.Marshal(expectedBuilds)

	runner := MockCommandRunner{Output: output}
	builds, err := eas.ListBuilds(runner)

	if err != nil {
		t.Fatalf("ListBuilds failed: %v", err)
	}

	if len(builds) != 2 {
		t.Errorf("expected 2 builds, got %d", len(builds))
	}

	if builds[0].ID != "1" {
		t.Errorf("expected ID '1', got %s", builds[0].ID)
	}
}
