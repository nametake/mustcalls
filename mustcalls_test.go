package mustcalls_test

import (
	"path/filepath"
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/nametake/mustcalls"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	tests := []struct {
		configFile string
		patterns   []string
	}{
		{
			configFile: "testdata/src/primitive/config.yaml",
			patterns:   []string{"primitive"},
		},
	}

	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	for _, tt := range tests {
		configFile := tt.configFile
		defaultPath, err := filepath.Abs(configFile)
		if err != nil {
			t.Error(err)
			return
		}
		if err := mustcalls.Analyzer.Flags.Set("config", defaultPath); err != nil {
			t.Error(err)
			return
		}
		analysistest.Run(t, testdata, mustcalls.Analyzer, tt.patterns...)
	}
}