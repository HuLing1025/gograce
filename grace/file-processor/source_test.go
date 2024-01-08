package fileprocessor

import (
	"os"
	"testing"
)

func TestBuildASTTree(t *testing.T) {
	_, err := NewSourceCodeProcessor().GetASTTree(".")
	if err != nil {
		t.Error("GetASTTree error")
	}
}

func TestGenerateSourceFile(t *testing.T) {
	root, err := NewSourceCodeProcessor().GetASTTree(".")
	if err != nil {
		t.Error("GetASTTree error")
	}

	newFileName := "generate_demo.go"
	root.Children[0].RelativePath = "."
	root.Children[0].FileName = newFileName
	err = NewSourceCodeProcessor().GenerateSourceFile(root.Children[0])
	if err != nil {
		t.Error("GenerateSourceFile error")
	}

	os.Remove(newFileName)
}
