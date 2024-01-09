package fileprocessor

import (
	"bytes"
	"errors"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

// ISourceCodeProcessor the processor of source code.
type ISourceCodeProcessor interface {
	GetASTTree(rootPath string) (root *ASTNode, err error)
	GenerateSourceFile(node *ASTNode) (err error)
	SearchFuncDecl(node *ast.File, fName string) (ast.Decl, bool)
	SearchMethodDecl(node *ast.File, sName string, mName string) (ast.Decl, bool)
	GetStatement(decl ast.Decl) (statement string, err error)
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

// GetASTTree build AST tree.
func (s *SourceCodeProcessor) GetASTTree(rootPath string) (root *ASTNode, err error) {
	root = &ASTNode{FType: Directory, RelativePath: rootPath}

	// create a file set.
	fSet := token.NewFileSet()

	err = dfs(rootPath, &root, fSet)

	return
}

// GenerateSourceFile generate source file.
func (s *SourceCodeProcessor) GenerateSourceFile(node *ASTNode) (err error) {
	file, _err := os.Create(node.RelativePath + "/" + node.FileName)
	if _err != nil {
		return _err
	}
	defer file.Close()

	fset := token.NewFileSet()
	fset.AddFile(node.FileName, fset.Base(), 0)

	err = printer.Fprint(file, fset, node.AST)

	return
}

// SearchFuncDecl search func declaration.
func (s *SourceCodeProcessor) SearchFuncDecl(node *ast.File, fName string) (ast.Decl, bool) {
	for _, decl := range node.Decls {
		// function.
		if fn, ok := decl.(*ast.FuncDecl); ok && fn.Recv == nil {
			return decl, true
		}
	}

	return nil, false
}

// SearchMethodDecl search method declaration.
func (s *SourceCodeProcessor) SearchMethodDecl(node *ast.File, sName string, mName string) (ast.Decl, bool) {
	for _, decl := range node.Decls {
		// method
		if fn, ok := decl.(*ast.FuncDecl); ok && fn.Recv != nil {
			return decl, true
		}
	}

	return nil, false
}

// GetStatement get statement.
func (s *SourceCodeProcessor) GetStatement(decl ast.Decl) (statement string, err error) {
	printerConfig := &printer.Config{Mode: printer.TabIndent | printer.UseSpaces}
	buf := bytes.Buffer{}

	fSet := token.NewFileSet()
	err = printerConfig.Fprint(&buf, fSet, decl)
	statement = buf.String()

	return
}

// dfs search all go source files.
func dfs(rootPath string, astTree **ASTNode, fSet *token.FileSet) (err error) {
	if strings.HasSuffix(rootPath, ".go") {
		if _, _err := os.Stat(rootPath); _err != nil {
			return errors.New("file not found")
		}

		pathSlice := strings.Split(rootPath, "/")
		f, _err := parser.ParseFile(fSet, rootPath, nil, 0)
		if _err != nil {
			return _err
		}
		*astTree = &ASTNode{FType: SourceFile, FileName: pathSlice[len(pathSlice)-1], RelativePath: strings.Join(pathSlice[:len(pathSlice)-1], "/"), AST: f}
		return
	}

	currents, _err := os.ReadDir(rootPath)
	if _err != nil {
		return _err
	}

	for _, file := range currents {
		if file.IsDir() && !strings.HasPrefix(file.Name(), ".") {
			var subDir = &ASTNode{FType: Directory, RelativePath: rootPath + "/" + file.Name()}

			err = dfs(rootPath+"/"+file.Name(), &subDir, fSet)
			if err != nil {
				return
			}

			// not null dir.
			if len(subDir.Children) != 0 {
				(*astTree).Children = append((*astTree).Children, subDir)
			}

			continue
		}

		if !strings.HasPrefix(file.Name(), ".") && !strings.HasPrefix(file.Name(), "_") &&
			!strings.HasSuffix(file.Name(), "_test.go") && strings.HasSuffix(file.Name(), ".go") {

			f, _err := parser.ParseFile(fSet, rootPath+"/"+file.Name(), nil, 0)
			if _err != nil {
				return _err
			}

			(*astTree).Children = append((*astTree).Children, &ASTNode{FType: SourceFile, FileName: file.Name(), RelativePath: rootPath, AST: f})
		}
	}

	return
}
