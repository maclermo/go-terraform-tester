package steps

import (
	"fmt"

	"go-terraform-tester/pkg/config"
	. "go-terraform-tester/pkg/internal/logger"
	"go-terraform-tester/pkg/internal/utils"
)

func RunTfOutputAssertGreaterThan(test config.Steps, output map[string]interface{}) error {
	opts := test.Options.(*config.TfOutputAssertGreaterThan)

	value, ok := output[opts.Field]
	if !ok {
		return fmt.Errorf("field %s not present in payload", opts.Field)
	}

	var valueFloat float64

	switch val := value.(type) {
	case int:
		valueFloat = float64(val)
	case float32:
		valueFloat = float64(val)
	case float64:
		valueFloat = val
	default:
		return fmt.Errorf("cannot use type %T for this assert greater than context in value", val)
	}

	var greaterThanFloat float64

	switch val := opts.GreaterThan.(type) {
	case int:
		greaterThanFloat = float64(val)
	case float32:
		greaterThanFloat = float64(val)
	case float64:
		greaterThanFloat = val
	default:
		return fmt.Errorf("cannot use type %T for this assert greater than context in greater than", val)
	}

	valueNumber := utils.ConvertNumber(valueFloat)
	greaterThanNumber := utils.ConvertNumber(greaterThanFloat)

	var returnedValueNumber, returnedGreaterThanNumber string

	if opts.Sensitive {
		returnedValueNumber = "<sensitive>"
		returnedGreaterThanNumber = "<sensitive>"
	} else {
		returnedValueNumber = valueNumber
		returnedGreaterThanNumber = greaterThanNumber
	}

	Logger.Debugf("Testing if \"%s\" is greater than \"%s\" on field \"%s\"", returnedValueNumber, returnedGreaterThanNumber, opts.Field)

	if valueFloat < greaterThanFloat {
		return fmt.Errorf("value %s is not greater than %s", returnedValueNumber, returnedGreaterThanNumber)
	}

	return nil
}
