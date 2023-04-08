package steps

import (
	"fmt"
	"reflect"

	"go-terraform-tester/pkg/config"
	. "go-terraform-tester/pkg/internal/logger"
)

func RunTfOutputAssertType(test config.Steps, output map[string]interface{}) error {
	opts := test.Options.(*config.TfOutputAssertType)

	value, ok := output[opts.Field]
	if !ok {
		return fmt.Errorf("field %s not present in payload", opts.Field)
	}

	valueType := reflect.TypeOf(value)

	var realValue string

	switch valueType {
	case reflect.TypeOf([]interface{}{}):
		realValue = "list"
	case reflect.TypeOf(map[string]interface{}{}):
		realValue = "map"
	case reflect.TypeOf(""):
		realValue = "string"
	case reflect.TypeOf(0), reflect.TypeOf(0.0):
		realValue = "number"
	case reflect.TypeOf(false):
		realValue = "bool"
	default:
		return fmt.Errorf("type %v is not known", valueType)
	}

	var returnedField, returnedRealValue string

	if opts.Sensitive {
		returnedField = "<sensitive>"
		returnedRealValue = "<sensitive>"
	} else {
		returnedField = opts.Field
		returnedRealValue = realValue
	}

	Logger.Debugf("Testing if field \"%s\" (%s) is of type \"%s\"", returnedField, returnedRealValue, opts.Type)

	if opts.Type != realValue {
		return fmt.Errorf("value %s is not of type %s", returnedField, opts.Type)
	}

	return nil
}
