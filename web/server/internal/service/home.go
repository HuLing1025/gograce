package service

import (
	"go/ast"
	fileprocessor "main/grace/file-processor"
	"main/web/server/internal/service/dto"
	"main/web/server/pkg/errpkg"
	"main/web/server/pkg/response"
)

// IHomeService home service interfaces.
type IHomeService interface {
	GetTree() (directories []dto.Directory, err errpkg.IError)
}

// HomeService home service.
type HomeService struct {
	fProcessor fileprocessor.ISourceCodeProcessor
}

// NewHomeService create a home service.
func NewHomeService() IHomeService {
	return &HomeService{
		fProcessor: fileprocessor.NewSourceCodeProcessor(),
	}
}

// GetTree get tree.
func (s *HomeService) GetTree() (directories []dto.Directory, err errpkg.IError) {
	root, _err := s.fProcessor.BuildASTTree(".")
	if _err != nil {
		err = errpkg.NewHighErrorWithCause(_err, response.BuildASTError)
		return
	}

	directories, _err = dfsFiles(root)
	if _err != nil {
		err = errpkg.NewHighErrorWithCause(_err, response.SearchASTError)
		return
	}

	return
}

// dfsFiles deep first search files.
func dfsFiles(root *fileprocessor.ASTNode) (directories []dto.Directory, err error) {
	var directory = dto.Directory{Name: root.RelativePath}

	// ignore null directory.
	if len(root.Children) == 0 && root.FType == fileprocessor.Directory {
		return
	}

	for index := range root.Children {
		switch root.Children[index].FType {
		case fileprocessor.Directory:
			subDirectories, _err := dfsFiles(root.Children[index])
			if _err != nil {
				err = _err
				return
			}
			directories = append(directories, subDirectories...)
		case fileprocessor.SourceFile:
			// get functions and methods.
			fs, ms := dfsFuncs(root.Children[index].AST.Decls)
			// ignore source file which has no functions and methods.
			if len(fs) == 0 && len(ms) == 0 {
				continue
			}

			var file = dto.FileDto{
				Name:      root.Children[index].FileName,
				Functions: fs,
				Methods:   ms,
			}
			directory.Files = append(directory.Files, file)
		}
	}

	// ignore null directory.
	if len(directory.Files) != 0 {
		directories = append(directories, directory)
	}

	return
}

// dfsFuncs deep first search functions.
func dfsFuncs(decls []ast.Decl) (fs []string, ms []dto.MethodDto) {
	for _, decl := range decls {
		switch f := decl.(type) {
		case *ast.FuncDecl:
			// it's a function.
			if f.Recv == nil {
				fs = append(fs, f.Name.Name)
				continue
			}
			// it's a method.
			ms = append(ms, dto.MethodDto{
				Name:   f.Name.Name,
				Struct: getStructName(f),
			})
		default:
		}
	}

	return
}

// getStructName get struct name.
func getStructName(decl *ast.FuncDecl) string {
	return decl.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name
}
