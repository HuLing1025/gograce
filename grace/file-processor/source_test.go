package fileprocessor

import "testing"

func TestBuildASTTree(t *testing.T) {
	_, err := NewSourceCodeProcessor().BuildASTTree("C:/Users/Xqq/Desktop/projectsOnGit/gotesty")
	if err == nil {
		t.Error("root is nil")
	}
}
