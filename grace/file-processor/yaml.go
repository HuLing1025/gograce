package fileprocessor

import (
	"fmt"
	"os"
)

// IYamlFileProcessor YAML file processor export interfaces.
type IYamlFileProcessor interface {
	CreateYamlFile(filePath string, configs []byte) error
	ReadYaml(filePath string) ([]byte, error)
}

// YamlFileProcessor YAML file processor
type YamlFileProcessor struct {
}

// NewYamlFileProcessor new a YAML file processor.
func NewYamlFileProcessor() IYamlFileProcessor {
	return &YamlFileProcessor{}
}

// CreateYamlFile create YAML config file.
func (c *YamlFileProcessor) CreateYamlFile(filePath string, configs []byte) (err error) {
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
func (c *YamlFileProcessor) ReadYaml(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}
