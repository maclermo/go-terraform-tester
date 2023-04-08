package cmd

import (
	"fmt"
	"os"

	"go-terraform-tester/pkg/config"
	. "go-terraform-tester/pkg/internal/logger"
	"go-terraform-tester/pkg/tests"
)

func Exec() error {
	Logger.Debug("Starting go-terraform-tester...")

	var tc config.TesterConfig

	Logger.Debug("Parsing configuration...")

	if err := tc.ParseConfig(); err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	Logger.Debug("Validating configuration...")

	if err := tc.ValidateConfig(); err != nil {
		return fmt.Errorf("error validating config: %w", err)
	}

	Logger.Debug("Dispatching tests...")

	if err := tests.DispatchTests(tc); err != nil {
		return err
	}

	return nil
}

func Run() {
	if err := Exec(); err != nil {
		Logger.Error(err)
		os.Exit(1)
	}
}
