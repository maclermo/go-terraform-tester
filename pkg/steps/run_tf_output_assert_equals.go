package steps

import (
	"fmt"

	"go-terraform-tester/pkg/config"
	. "go-terraform-tester/pkg/internal/logger"
	"go-terraform-tester/pkg/internal/utils"
)

func RunTfOutputAssertEquals(test config.Steps, output map[string]interface{}) error {
	opts := test.Options.(*config.TfOutputAssertEquals)

	value, ok := output[opts.Field]
	if !ok {
		return fmt.Errorf("field %s not present in payload", opts.Field)
	}

	var valueString string

	switch val := value.(type) {
	case string:
		valueString = val
	case float32:
		valueString = utils.ConvertNumber(float64(val))
	case float64:
		valueString = utils.ConvertNumber(val)
	case bool:
		valueString = utils.ConvertBoolean(val)
	default:
		return fmt.Errorf("cannot use type %T for this assert equals context in value", val)
	}

	var equalsString string

	switch val := opts.Equals.(type) {
	case string:
		equalsString = val
	case float32:
		equalsString = utils.ConvertNumber(float64(val))
	case float64:
		equalsString = utils.ConvertNumber(val)
	case bool:
		equalsString = utils.ConvertBoolean(val)
	default:
		return fmt.Errorf("cannot use type %T for this assert equals context in equals", val)
	}

	var returnedValueString, returnedEqualsString string

	if opts.Sensitive {
		returnedValueString = "<sensitive>"
		returnedEqualsString = "<sensitive>"
	} else {
		returnedValueString = valueString
		returnedEqualsString = equalsString
	}

	Logger.Debugf("Testing if \"%s\" equals \"%s\" on field \"%s\"", returnedValueString, returnedEqualsString, opts.Field)

	if valueString != equalsString {
		return fmt.Errorf("value %s does not equal %s", returnedValueString, returnedEqualsString)
	}

	return nil
}
