package tools

import "errors"

// ErrNotEnoughArguments é um erro que indica que não foram fornecidos argumentos suficientes.
var ErrNotEnoughArguments = errors.New("not enough arguments")

// ReadCmdParameters reads the second command line argument.
func ReadCmdParameters(args []string) (string, error) {
	if len(args) < 2 {
		return "", ErrNotEnoughArguments
	}
	return args[1], nil
}
