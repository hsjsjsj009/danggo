package danggo

import (
	"errors"
	"fmt"
)

func WrongTypeError(expected string) error {
	return errors.New(fmt.Sprintf("Wrong path input type, expected type : %s",expected))
}

func NotSupportedType(unsupportedType string) error {
	return errors.New(fmt.Sprintf("Unsupported type %s",unsupportedType))
}


