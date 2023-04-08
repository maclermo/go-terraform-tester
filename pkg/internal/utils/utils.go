package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func ConvertNumber(number float64) string {
	numberString := fmt.Sprintf("%f", number)
	numberStringTrimmed := strings.TrimRight(numberString, "0")

	return strings.TrimRight(numberStringTrimmed, ".")
}

func ConvertBoolean(boolean bool) string {
	if boolean {
		return "true"
	}

	return "false"
}

func RunCommand(command string, shell string, inline bool) ([]byte, error) {
	var cmd *exec.Cmd

	if inline {
		cmd = exec.Command(command)
	} else {
		cmd = exec.Command(shell, "-c", command)
	}

	stdout, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return stdout, nil
}
