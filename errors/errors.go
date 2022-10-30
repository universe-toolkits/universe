package errors

import "fmt"

func ServiceNotFoundError(msg string, serviceName string) error {
	return fmt.Errorf(msg)
}

func InvalidActionNameError() error {
	return fmt.Errorf("invalid action name")
}
