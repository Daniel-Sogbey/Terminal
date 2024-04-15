package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

// custom json encoding
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

// custom json decoding
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {

	unquotedJsonValue, err := strconv.Unquote(string(jsonValue))

	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJsonValue, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 64)

	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)

	fmt.Println("RUNTIME - ", *r)

	return nil
}
