package fileprocessor

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// IYamlFileProcessor YAML file processor export interfaces.
type IYamlFileProcessor interface {
	// Init()
	CreateConfigFile(configs []byte) error
	AnalysisConfig() (config TestConfig, err error)
	// UpdateConfig() error
	// RecoverHistory() error
}

// YamlFileProcessor YAML file processor
type YamlFileProcessor struct {
}

// StubLevel stub level type
type StubLevel int

const (
	// FuncStubLevel function stub type
	FuncStubLevel StubLevel = iota + 1
	// CaseStubLevel case stub type
	CaseStubLevel
)

// TestConfig test config
type TestConfig struct {
	FileConfigs []FileConfig `json:"file_configs" yaml:"FILE_CONFIGS"`
}

// FileConfig file config
type FileConfig struct {
	Ignored       bool           `json:"ignored" yaml:"IGNORED"`
	Name          string         `json:"name" yaml:"NAME"`
	Path          string         `json:"path" yaml:"PATH"`
	Package       string         `json:"package" yaml:"PACKAGE"`
	FunConfigs    []FunConfig    `json:"fun_configs" yaml:"FUN_CONFIGS"`
	MethodConfigs []MethodConfig `json:"method_configs" yaml:"METHOD_CONFIGS"`
}

// FunConfig function config
type FunConfig struct {
	Ignored   bool       `json:"ignored" yaml:"IGNORED"`
	Name      string     `json:"name" yaml:"NAME"`
	TestCases []TestCase `json:"test_cases" yaml:"TEST_CASES"`
	Stubs     []Stub     `json:"stubs" yaml:"STUBS"`
}

// MethodConfig method config
type MethodConfig struct {
	FunConfig  `yaml:"FUN_CONFIG"`
	StructName string `json:"struct_name" yaml:"STRUCT_NAME"`
	Package    string `json:"package" yaml:"PACKAGE"`
}

// Param param
type Param struct {
	Name      string      `json:"name" yaml:"NAME"`
	Package   string      `json:"package" yaml:"PACKAGE"`
	ParamType string      `params:"param_type" yaml:"PARAM_TYPE"`
	Value     interface{} `json:"value" yaml:"VALUE"`
}

// TestCase test case
type TestCase struct {
	Desc    string  `json:"desc" yaml:"DESC"`
	Inputs  []Param `json:"inputs" yaml:"INPUTS"`
	Outputs []Param `json:"outputs" yaml:"OUTPUTS"` // wants return.
	Globals []Param `json:"globals" yaml:"GLOBALS"` // wants global.
	Stub    []Stub  `json:"stubs" yaml:"STUBS"`
}

// Stub stub
type Stub struct {
	Level      StubLevel   `json:"stub_level" yaml:"STUB_LEVEL"`
	TargetName string      `json:"target_name" yaml:"TARGET_NAME"`
	Package    string      `json:"package" yaml:"PACKAGE"`
	Value      interface{} `json:"value" yaml:"VALUE"`
}

// NewYamlFileProcessor new a YAML file processor.
func NewYamlFileProcessor() IYamlFileProcessor {
	return &YamlFileProcessor{}
}

// CreateConfigFile create YAML config file.
func (c *YamlFileProcessor) CreateConfigFile(configs []byte) (err error) {
	yamlFile, err := os.Create("GRACE.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer yamlFile.Close()

	_, err = yamlFile.Write(configs)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

// AnalysisConfig analysis YAML config file.
func (c *YamlFileProcessor) AnalysisConfig() (config TestConfig, err error) {
	configData, _err := os.ReadFile("GRACE.yaml")
	if _err != nil {
		err = _err
		return
	}

	err = yaml.Unmarshal(configData, &config)

	return
}
