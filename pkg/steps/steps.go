package steps

import (
	"fmt"
	"go-terraform-tester/pkg/config"
	. "go-terraform-tester/pkg/internal/logger"
)

func DispatchTests(test config.Tests, output map[string]interface{}) error {
	for _, step := range test.Steps {
		Logger.Debugf("Launching step %s on %s", step.Test, test.Config.Type)

		switch step.Test {
		case "ShellRunCommand":
			if err := RunShellRunCommand(step, output); err != nil {
				return fmt.Errorf("error while running step ShellRunCommand(): %w", err)
			}
		case "TfOutputAssertEquals":
			if err := RunTfOutputAssertEquals(step, output); err != nil {
				return fmt.Errorf("error while running step TfOutputAssertEquals(): %w", err)
			}
		case "TfOutputAssertGreaterThan":
			if err := RunTfOutputAssertGreaterThan(step, output); err != nil {
				return fmt.Errorf("error while running step TfOutputAssertGreaterThan(): %w", err)
			}
		case "TfOutputAssertLessThan":
			if err := RunTfOutputAssertLessThan(step, output); err != nil {
				return fmt.Errorf("error while running step TfOutputAssertLessThan(): %w", err)
			}
		case "TfOutputAssertType":
			if err := RunTfOutputAssertType(step, output); err != nil {
				return fmt.Errorf("error while running step TfOutputAssertType(): %w", err)
			}
		}
	}

	return nil
}
