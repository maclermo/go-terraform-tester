package config

import (
	"fmt"
	"io/ioutil"

	. "go-terraform-tester/pkg/internal/logger"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

var validate *validator.Validate

type TesterConfig struct {
	Tests []Tests `yaml:"tests"`
}

type Tests struct {
	Name        string  `yaml:"name" validate:"required"`
	Description string  `yaml:"description" validate:"required"`
	Config      Config  `yaml:"config" validate:"required"`
	Steps       []Steps `yaml:"steps" validate:"dive,min=1"`
}

type Config struct {
	Type       string      `yaml:"type" validate:"required"`
	Parameters interface{} `yaml:"parameters" validate:"required"`
}

type TerraformParams struct {
	Location       string `yaml:"location" validate:"required"`
	BackendConfig  string `yaml:"backendConfig" validate:"required"`
	InputVariables string `yaml:"inputVariables" validate:"required"`
}

type TerragruntParams struct {
	Location       string `yaml:"location" validate:"required"`
	InputVariables string `yaml:"inputVariables" validate:"required"`
}

type Steps struct {
	Test    string      `yaml:"test" validate:"required"`
	Options interface{} `yaml:"options" validate:"required"`
}

type ShellRunCommand struct {
	Command string `yaml:"command" validate:"required"`
	Inline  bool   `yaml:"inline"`
	Shell   string `yaml:"shell" validate:"required_if=Inline false"`
	Sensitive bool `yaml:"sensitive"`
}

type TfOutputAssertEquals struct {
	Field  string      `yaml:"field" validate:"required"`
	Equals interface{} `yaml:"equals" validate:"required"`
	Sensitive bool `yaml:"sensitive"`
}

type TfOutputAssertGreaterThan struct {
	Field       string      `yaml:"field" validate:"required"`
	GreaterThan interface{} `yaml:"greaterThan" validate:"required"`
	Sensitive bool `yaml:"sensitive"`
}

type TfOutputAssertLessThan struct {
	Field    string      `yaml:"field" validate:"required"`
	LessThan interface{} `yaml:"lessThan" validate:"required"`
	Sensitive bool `yaml:"sensitive"`
}

type TfOutputAssertType struct {
	Field string `yaml:"field" validate:"required"`
	Type  string `yaml:"type" validate:"required,oneof=list map string number bool"`
	Sensitive bool `yaml:"sensitive"`
}

func (c *Config) UnmarshalYAML(n *yaml.Node) error {
	var z struct {
		Type       string    `yaml:"type"`
		Parameters yaml.Node `yaml:"parameters"`
	}

	if err := n.Decode(&z); err != nil {
		return err
	}

	c.Type = z.Type

	switch z.Type {
	case "terraform":
		c.Parameters = new(TerraformParams)
	case "terragrunt":
		c.Parameters = new(TerragruntParams)
	default:
		return fmt.Errorf("invalid config type %s", z.Type)
	}

	return z.Parameters.Decode(c.Parameters)
}

func (s *Steps) UnmarshalYAML(n *yaml.Node) error {
	var z struct {
		Test    string    `yaml:"test"`
		Options yaml.Node `yaml:"options"`
	}

	if err := n.Decode(&z); err != nil {
		return err
	}

	s.Test = z.Test

	switch z.Test {
	case "ShellRunCommand":
		s.Options = new(ShellRunCommand)
	case "TfOutputAssertEquals":
		s.Options = new(TfOutputAssertEquals)
	case "TfOutputAssertGreaterThan":
		s.Options = new(TfOutputAssertGreaterThan)
	case "TfOutputAssertLessThan":
		s.Options = new(TfOutputAssertLessThan)
	case "TfOutputAssertType":
		s.Options = new(TfOutputAssertType)
	default:
		return fmt.Errorf("invalid test type %s", z.Test)
	}

	return z.Options.Decode(s.Options)
}

func (tc *TesterConfig) ParseConfig() error {
	filename := "config.yaml"

	Logger.Trace(fmt.Sprintf("Opening file %s", filename))

	yamlConfig, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error opening configuration file: %w", err)
	}

	Logger.Trace(fmt.Sprintf("Unmarshalling file %s to yaml", filename))

	if err := yaml.Unmarshal(yamlConfig, tc); err != nil {
		return fmt.Errorf("cannot unmarshal yaml file: %w", err)
	}

	return nil
}

func (tc *TesterConfig) ValidateConfig() error {
	validate = validator.New()

	for _, test := range tc.Tests {
		if err := validate.Struct(test); err != nil {
			return err
		}
	}

	return nil
}
