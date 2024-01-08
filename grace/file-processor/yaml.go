package fileprocessor

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// IYamlFileProcessor YAML file processor export interfaces.
type IYamlFileProcessor interface {
	CreateYamlFile(filePath string, input any) error
	ReadYaml(filePath string, output any) error
}

// YamlFileProcessor YAML file processor
type YamlFileProcessor struct {
}

// NewYamlFileProcessor new a YAML file processor.
func NewYamlFileProcessor() IYamlFileProcessor {
	return &YamlFileProcessor{}
}

// CreateYamlFile create YAML config file.
func (c *YamlFileProcessor) CreateYamlFile(filePath string, input any) (err error) {
	configs, _err := yaml.Marshal(input)
	if _err != nil {
		err = _err
		return
	}

	yamlFile, err := os.Create(filePath)
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

// ReadYaml read yaml file.
func (c *YamlFileProcessor) ReadYaml(filePath string, output any) error {
	configByte, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(configByte, output); err != nil {
		return err
	}

	return nil
}
