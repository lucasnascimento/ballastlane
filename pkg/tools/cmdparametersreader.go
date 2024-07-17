package tools

import "errors"

// ErrNotEnoughArguments is an error returned when there are not enough command line arguments.
var ErrNotEnoughArguments = errors.New("not enough arguments")

// ReadCmdParameters reads the second command line argument.
func ReadCmdParameters(args []string) (string, error) {
	if len(args) < 2 {
		return "", ErrNotEnoughArguments
	}
	return args[1], nil
}
