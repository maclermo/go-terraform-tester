package tests

import (
	"fmt"
	"go-terraform-tester/pkg/config"
	. "go-terraform-tester/pkg/internal/logger"

	"github.com/gruntwork-io/terratest/modules/testing"
)

type tLogger struct{}

func (tLogger) Logf(t testing.TestingT, format string, args ...interface{}) {
	msg := fmt.Sprintf("[TERRATEST]  %s", fmt.Sprintf(format, args...))
	if msg != "[TERRATEST]  " {
		Logger.Debug(msg)
	}
}

func DispatchTests(tc config.TesterConfig) error {
	for _, test := range tc.Tests {
		Logger.Debugf("Launching test %s (%s) on %s", test.Name, test.Description, test.Config.Type)

		switch test.Config.Type {
		case "terraform":
			if err := RunTerraform(test); err != nil {
				return fmt.Errorf("error while running test \"%s\" using RunTerraform(): %w", test.Name, err)
			}
		case "terragrunt":
			if err := RunTerragrunt(test); err != nil {
				return fmt.Errorf("error running test \"%s\" using RunTerragrunt(): %w", test.Name, err)
			}
		}
	}

	return nil
}
