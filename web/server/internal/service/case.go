package service

import (
	"encoding/json"
	fileprocessor "main/grace/file-processor"
	"main/web/server/internal/service/dto"
	"main/web/server/pkg/errpkg"
	"main/web/server/pkg/response"
)

// ICaseService test case service interface.
type ICaseService interface {
	GetTestCases() (response dto.CaseResponse, err errpkg.IError)
}

// CaseService test case service
type CaseService struct {
	yProcessor fileprocessor.IYamlFileProcessor
}

// NewCaseService new a test case service.
func NewCaseService() ICaseService {
	return &CaseService{
		yProcessor: fileprocessor.NewYamlFileProcessor(),
	}
}

// testCaseYamlPath	test case yaml path.
const testCaseYamlPath = "./GRACE.yaml"

// GetTestCases get test cases.
func (s *CaseService) GetTestCases() (cases dto.CaseResponse, err errpkg.IError) {
	configs, _err := s.yProcessor.ReadYaml(testCaseYamlPath)
	if _err != nil {
		err = errpkg.NewMiddleErrorWithCause(_err, response.ReadYAMLConfigError)
		return
	}

	var testConfigs []dto.TestCaseConfig
	if _err = json.Unmarshal(configs, &testConfigs); _err != nil {
		err = errpkg.NewMiddleErrorWithCause(_err, response.JSONError)
		return
	}

	return

}