package fileprocessor

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"strings"
)

// ISourceCodeProcessor the processor of source code.
type ISourceCodeProcessor interface {
	BuildASTTree(rootPath string) (root *ASTNode, err error)
}

// SourceCodeProcessor source code processor.
type SourceCodeProcessor struct {
}

// FileType file type
type FileType int

const (
	// FileUnknown unknown file type
	FileUnknown FileType = iota
	// Directory directory
	Directory
	// SourceFile go file
	SourceFile
)

// ASTNode AST node
type ASTNode struct {
	FType        FileType
	AST          *ast.File
	FileName     string
	RelativePath string
	Children     []*ASTNode
}

// NewSourceCodeProcessor new source code processor.
func NewSourceCodeProcessor() ISourceCodeProcessor {
	return &SourceCodeProcessor{}
}

// BuildASTTree build AST tree.
func (s *SourceCodeProcessor) BuildASTTree(rootPath string) (root *ASTNode, err error) {
	root = &ASTNode{FType: Directory, RelativePath: rootPath}

	// create a file set.
	fSet := token.NewFileSet()

	err = dfs(rootPath, root, fSet)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

// FileInfo file info
type FileInfo struct {
	fs.DirEntry
	Path string
}

// dfs search all go source files.
func dfs(rootPath string, astTree *ASTNode, fSet *token.FileSet) (err error) {
	currents, _err := os.ReadDir(rootPath)
	if _err != nil {
		err = _err
		return
	}

	for _, file := range currents {
		if file.IsDir() && !strings.HasPrefix(file.Name(), ".") {
			var subDir = ASTNode{FType: Directory, RelativePath: rootPath + "/" + file.Name()}

			err = dfs(rootPath+"/"+file.Name(), &subDir, fSet)
			if err != nil {
				return
			}

			// not null dir.
			if len(subDir.Children) != 0 {
				astTree.Children = append(astTree.Children, &subDir)
			}

			continue
		}

		if !strings.HasPrefix(file.Name(), ".") && !strings.HasPrefix(file.Name(), "_") &&
			!strings.HasSuffix(file.Name(), "_test.go") && strings.HasSuffix(file.Name(), ".go") {

			f, _err := parser.ParseFile(fSet, rootPath+"/"+file.Name(), nil, 0)
			if _err != nil {
				fmt.Println("解析文件失败：", _err)
				return
			}

			astTree.Children = append(astTree.Children, &ASTNode{FType: SourceFile, FileName: file.Name(), RelativePath: rootPath, AST: f})
		}
	}

	return
}
