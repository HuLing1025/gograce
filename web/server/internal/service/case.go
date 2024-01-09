package service

import (
	"errors"
	"go/ast"
	fileprocessor "main/grace/file-processor"
	"main/web/server/internal/service/dto"
	"main/web/server/pkg/errpkg"
	"main/web/server/pkg/response"
)

// ICaseService test case service interface.
type ICaseService interface {
	GetTestCases(request dto.CaseRequest) (cases dto.CaseResponse, err errpkg.IError)
	AddTestCase(request dto.AddCaseRequestDto) (newCases dto.CaseResponse, err errpkg.IError)
}

// CaseService test case service
type CaseService struct {
	yProcessor fileprocessor.IYamlFileProcessor
	sProcessor fileprocessor.ISourceCodeProcessor
}

// NewCaseService new a test case service.
func NewCaseService() ICaseService {
	return &CaseService{
		yProcessor: fileprocessor.NewYamlFileProcessor(),
		sProcessor: fileprocessor.NewSourceCodeProcessor(),
	}
}

// testCaseYamlPath	test case yaml path.
const testCaseYamlPath = "./GRACE.yaml"

// GetTestCases get test cases.
func (s *CaseService) GetTestCases(request dto.CaseRequest) (cases dto.CaseResponse, err errpkg.IError) {
	// get function/method declaration.
	decl, _err := s.getDecl(request)
	if _err != nil {
		err = errpkg.NewMiddleErrorWithCause(_err, response.SearchFunctionError)
		return
	}

	// package configuration.
	cases.BaseInfo.Name = request.FuncName
	cases.BaseInfo.StructName = request.StructName
	// TODO other base information...
	// cases.BaseInfo.Inputs =
	// cases.BaseInfo.OutPuts =
	// cases.BaseInfo.Signature =

	statement, _err := s.sProcessor.GetStatement(decl)
	if _err != nil {
		err = errpkg.NewMiddleErrorWithCause(_err, response.SearchASTError)
		return
	}
	cases.BaseInfo.Statement = statement

	var testConfigs []dto.TestCaseConfig
	_err = s.yProcessor.ReadYaml(testCaseYamlPath, &testConfigs)
	if _err != nil {
		err = errpkg.NewMiddleErrorWithCause(_err, response.ReadYAMLConfigError)
		return
	}

	// find test case config.
	for _, file := range testConfigs {
		if file.Path == request.Path && file.FileName == request.FileName {
			cases.VarStubList = file.VarStubList
			cases.FuncStubList = file.FuncStubList
			for _, funCase := range file.FunctionConfigs {
				if funCase.StructName == request.StructName && funCase.Name == request.FuncName {
					cases.Cases = funCase.Cases
				}
			}
			break
		}
	}

	return
}

// AddTestCase add a test case.
func (s *CaseService) AddTestCase(request dto.AddCaseRequestDto) (cases dto.CaseResponse, err errpkg.IError) {
	var testConfigs []dto.TestCaseConfig
	_err := s.yProcessor.ReadYaml(testCaseYamlPath, &testConfigs)
	if _err != nil {
		err = errpkg.NewMiddleErrorWithCause(_err, response.ReadYAMLConfigError)
		return
	}

	// find test case config.
	for _, file := range testConfigs {
		if file.Path == request.Path && file.FileName == request.FileName {
			for _, funCase := range file.FunctionConfigs {
				if funCase.StructName == request.StructName && funCase.Name == request.FuncName {
					// found the function/method test cases, do add.
					funCase.Cases = append(funCase.Cases, request.Case)
					break
				}
			}
			break
		}
	}

	// sync to yaml.
	_err = s.yProcessor.CreateYamlFile(testCaseYamlPath, testConfigs)
	if _err != nil {
		err = errpkg.NewMiddleErrorWithCause(_err, response.SyncConfigError)
		return
	}

	return s.GetTestCases(dto.CaseRequest{FuncName: request.FuncName, Path: request.Path, FileName: request.FileName, StructName: request.StructName})
}

// getDecl get function/method declaration.
func (s *CaseService) getDecl(request dto.CaseRequest) (ast.Decl, error) {
	astNode, _err := s.sProcessor.GetASTTree(request.Path + "/" + request.FileName)
	if _err != nil {
		return nil, _err
	}
	if len(astNode.AST.Decls) == 0 {
		return nil, errors.New("file ast not found")
	}

	var decl ast.Decl
	var found bool
	// it is a function.
	if request.StructName == "" {
		if decl, found = s.sProcessor.SearchFuncDecl(astNode.AST, request.FuncName); found {
			return decl, nil
		}
		return nil, errors.New("not found")
	}
	// it is a method.
	if decl, found = s.sProcessor.SearchMethodDecl(astNode.AST, request.StructName, request.FuncName); found {
		return decl, nil
	}
	return nil, errors.New("not found")
}
