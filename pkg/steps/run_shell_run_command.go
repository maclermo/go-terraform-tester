package steps

import (
	"go-terraform-tester/pkg/config"
	. "go-terraform-tester/pkg/internal/logger"
	"go-terraform-tester/pkg/internal/utils"
)

func RunShellRunCommand(test config.Steps, output map[string]interface{}) error {
	opts := test.Options.(*config.ShellRunCommand)

	if opts.Inline {
		Logger.Debugf("Running command: %s", opts.Command)
	} else {
		Logger.Debugf("Running inline command: %s -c %s", opts.Shell, opts.Command)
	}

	stdout, err := utils.RunCommand(opts.Command, opts.Shell, opts.Inline)

	if err != nil {
		return err
	}

	out := string(stdout)

	if out != "" {
		Logger.Debugf("Results: %s", out)
	}

	return nil
}
