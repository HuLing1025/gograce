package fileprocessor

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestCreateConfigFile(t *testing.T) {
	var config TestConfig

	root, err := NewSourceCodeProcessor().BuildASTTree(".")
	if err != nil {
		t.Error(err)
	}
	if len(root.Children) == 0 {
		t.Error("no files.")
	}

	for _, file := range root.Children {
		var fileConfig = FileConfig{
			Ignored:       false,
			Name:          file.FileName,
			Path:          file.RelativePath,
			Package:       file.AST.Name.Name,
			FunConfigs:    []FunConfig{},
			MethodConfigs: []MethodConfig{},
		}

		for _, decl := range file.AST.Decls {
			// function configs.
			if funcDecl, ok := decl.(*ast.FuncDecl); ok {
				var funcConfig = FunConfig{
					Ignored: false,
					Name:    funcDecl.Name.Name,
				}

				fileConfig.FunConfigs = append(fileConfig.FunConfigs, funcConfig)
				continue
			}

			// method configs.
			if genDecl, ok := decl.(*ast.GenDecl); ok {
				var structName string
				var methodName string
				for _, spec := range genDecl.Specs {
					switch s := spec.(type) {
					case *ast.TypeSpec: // 处理类型定义
						if t, ok := s.Type.(*ast.StructType); ok {
							t.Fields.NumFields()
							if len(t.Fields.List) == 0 || len(t.Fields.List[0].Names) == 0 {
								continue
							}
							v := t.Fields.List[0].Names[0]
							structName = v.String() + "." + s.Name.String()
							methodName = v.String() + "." + s.Name.String()
							break
						}
					default:
						continue
					}
				}
				var methConfig = MethodConfig{
					FunConfig: FunConfig{
						Ignored: false,
						Name:    methodName,
					},
					StructName: structName,
					Package:    file.AST.Name.Name,
				}
				fileConfig.MethodConfigs = append(fileConfig.MethodConfigs, methConfig)
				continue
			}
		}

		config.FileConfigs = append(config.FileConfigs, fileConfig)
	}

	yamlContent, err := yaml.Marshal(config)
	if err != nil {
		t.Error(err)
		return
	}

	err = NewYamlFileProcessor().CreateConfigFile(yamlContent)
	if err != nil {
		t.Error(err)
	}
}

func TestAnalysisConfig(t *testing.T) {
	config, err := NewYamlFileProcessor().AnalysisConfig()
	data, _ := json.Marshal(config)
	fmt.Println(string(data))
	if err != nil {
		t.Error(err)
	}
}
