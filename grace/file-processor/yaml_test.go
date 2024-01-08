package fileprocessor

import (
	"os"
	"testing"
)

func TestCreateConfigFile(t *testing.T) {
	var input any = "test yaml input"
	err := NewYamlFileProcessor().CreateYamlFile("GRACE.yaml", &input)
	if err != nil {
		t.Error(err)
	}
}

func TestAnalysisConfig(t *testing.T) {
	var output any
	err := NewYamlFileProcessor().ReadYaml("GRACE.yaml", &output)
	if err != nil {
		t.Error(err)
	}

	_ = os.Remove("GRACE.yaml")
}
