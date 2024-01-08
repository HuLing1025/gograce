package service

import (
	fileprocessor "main/grace/file-processor"
	"main/web/server/internal/service/dto"
	"os"
	"testing"
)

func TestGetCases(t *testing.T) {
	var configs = []dto.TestCaseConfig{
		{
			Path:     ".",
			FileName: "main.go",
			VarStubList: []dto.VarStub{
				{
					Name:  "systemConfig",
					Value: "systemConfig stub value",
				},
			},
			FuncStubList: []dto.FuncStub{
				{
					StructName: "Processor",
					Name:       "CreateFile",
					Statement:  "{\n\treturn true\n}",
				},
			},
			FunctionConfigs: []dto.FuncConfig{
				{
					StructName: "Processor",
					Name:       "CreateFile",
					Cases: []dto.Case{
						{
							Desc: "Test case 001",
							Asserts: []dto.Assert{
								{
									Input: dto.Param{Name: "path", Type: "string", Value: "."},
									Want:  dto.Param{Name: "path", Type: "string", Value: "."},
								},
								{
									Input: dto.Param{Name: "path", Type: "string", Value: "test"},
									Want:  dto.Param{Name: "path", Type: "string", Value: "test"},
								},
							},
						},
					},
				},
			},
		},
	}

	_ = fileprocessor.NewYamlFileProcessor().CreateYamlFile("./GRACE.yaml", configs)

	_, err := NewCaseService().GetTestCases()
	if err != nil {
		t.Error(err)
	}

	_ = os.Remove("./GRACE.yaml")

}
