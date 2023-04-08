package tests

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"go-terraform-tester/pkg/config"
	. "go-terraform-tester/pkg/internal/logger"
	"go-terraform-tester/pkg/steps"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

var Clean bool

func destroy(t *testing.T, a *terraform.Options) {
	if !Clean {
		Logger.Error("An error occured, terraform destroy has been called.")
		Logger.Error("The reason behind this failure will be available soon...")
	}

	_, _ = terraform.DestroyE(t, a)
}

func RunTerraform(tc config.Tests) error {
	params := tc.Config.Parameters.(*config.TerraformParams)

	Logger.Debug("Reading terraform files...")

	location, err := filepath.Abs(params.Location)
	if err != nil {
		return fmt.Errorf("error reading terraform files: %w", err)
	}

	Logger.Debug("Reading terraform input variables files...")

	inputVariables, err := filepath.Abs(params.InputVariables)
	if err != nil {
		return fmt.Errorf("error reading terraform input variables files: %w", err)
	}

	t := &testing.T{}

	a := &terraform.Options{
		TerraformDir:       location,
		VarFiles:           []string{inputVariables},
		Logger:             logger.New(tLogger{}),
		MigrateState:       true,
		NoColor:            true,
		NoStderr:           true,
		MaxRetries:         0,
		TimeBetweenRetries: time.Second * 0,
	}

	defer destroy(t, a)

	Logger.Debugf("Launching terraform.InitAndApplyE() on test %s (%s)", tc.Name, tc.Description)

	_, err = terraform.InitAndApplyE(t, a)
	if err != nil {
		return fmt.Errorf("error applying terraform plan: %w", err)
	}

	outputs, err := terraform.OutputAllE(t, a)
	if err != nil {
		return fmt.Errorf("error getting terraform outputs: %w", err)
	}

	err = steps.DispatchTests(tc, outputs)
	if err != nil {
		return err
	}

	Clean = true

	return nil
}
