package ociseccompgen

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	types "github.com/opencontainers/runc/libcontainer/configs"
)

// Arguments takes a list of arguments (delimArgs)  and a pointer to a
// corresponding syscall struct. It parses and fills out the argument information
func parseArguments(syscallArg string) ([]*types.Arg, error) {
	delimArgs := strings.Split(syscallArg, ":")

	nilArgSlice := []*types.Arg{}

	if len(delimArgs) == 1 {
		return nilArgSlice, nil
	}

	if len(delimArgs) == 5 {
		syscallIndex, err := strconv.ParseUint(delimArgs[1], 10, 0)
		if err != nil {
			return nilArgSlice, err
		}

		syscallValue, err := strconv.ParseUint(delimArgs[2], 10, 64)
		if err != nil {
			return nilArgSlice, err
		}

		syscallValueTwo, err := strconv.ParseUint(delimArgs[3], 10, 64)
		if err != nil {
			return nilArgSlice, err
		}

		syscallOp, err := parseOperator(delimArgs[4])
		if err != nil {
			return nilArgSlice, err
		}

		argStruct := &types.Arg{
			Index:    uint(syscallIndex),
			Value:    syscallValue,
			ValueTwo: syscallValueTwo,
			Op:       syscallOp,
		}

		argSlice := []*types.Arg{argStruct}
		return argSlice, nil
	}

	return nilArgSlice, errors.New("Incorrect number of arguments passed with syscall")
}

func parseOperator(operator string) (types.Operator, error) {
	switch operator {
	case "NE":
		return types.NotEqualTo, nil
	case "LT":
		return types.LessThan, nil
	case "LE":
		return types.LessThanOrEqualTo, nil
	case "EQ":
		return types.EqualTo, nil
	case "GE":
		return types.GreaterThanOrEqualTo, nil
	case "GT":
		return types.GreaterThan, nil
	case "ME":
		return types.MaskEqualTo, nil
	default:
		return types.NotEqualTo, fmt.Errorf("Unrecognized operator: %s", operator)
	}
}
