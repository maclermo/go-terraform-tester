package steps

import (
	"fmt"

	"go-terraform-tester/pkg/config"
	. "go-terraform-tester/pkg/internal/logger"
	"go-terraform-tester/pkg/internal/utils"
)

func RunTfOutputAssertLessThan(test config.Steps, output map[string]interface{}) error {
	opts := test.Options.(*config.TfOutputAssertLessThan)

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
		return fmt.Errorf("cannot use type %T for this assert less than context in value", val)
	}

	var lessThanFloat float64

	switch val := opts.LessThan.(type) {
	case int:
		lessThanFloat = float64(val)
	case float32:
		lessThanFloat = float64(val)
	case float64:
		lessThanFloat = val
	default:
		return fmt.Errorf("cannot use type %T for this assert less than context in less than", val)
	}

	valueNumber := utils.ConvertNumber(valueFloat)
	lessThanNumber := utils.ConvertNumber(lessThanFloat)

	var returnedValueNumber, returnedLessThanNumber string

	if opts.Sensitive {
		returnedValueNumber = "<sensitive>"
		returnedLessThanNumber = "<sensitive>"
	} else {
		returnedValueNumber = valueNumber
		returnedLessThanNumber = lessThanNumber
	}

	Logger.Debugf("Testing if \"%s\" is less than \"%s\" on field \"%s\"", returnedValueNumber, returnedLessThanNumber, opts.Field)

	if valueFloat > lessThanFloat {
		return fmt.Errorf("value %s is not less than %s", returnedValueNumber, returnedLessThanNumber)
	}

	return nil
}
