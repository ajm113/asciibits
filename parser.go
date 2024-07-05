package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	MaxASCIIRange int = 127
	MinASCIIRange int = 0
)

type (
	Tokens struct {
		decimals []int
	}

	ErrTokenParse struct {
		err error
		str string
		col int
	}
)

var (
	ErrParseNoValue       = errors.New("no value")
	ErrDecimalOutOfRange  = errors.New("decimal out of range")
	ErrDecimalInvalidChar = errors.New("decimal invalid character")
)

func StringToDecimals(input string) (*Tokens, error) {
	if len(input) == 0 {
		return nil, ErrParseNoValue
	}

	decimals := make([]int, len(input))

	for i, c := range input {
		decimals[i] = int(c)
	}

	t := &Tokens{
		decimals: decimals,
	}

	return t, nil
}

func ParseDecimals(input, separator string) (*Tokens, error) {
	input = strings.Trim(input, separator)

	if len(input) == 0 {
		return nil, ErrParseNoValue
	}

	tokens := strings.Split(input, separator)

	decimals := make([]int, len(tokens))

	for i, d := range tokens {
		decimalInt, err := strconv.Atoi(d)

		if err != nil {
			return nil, ErrTokenParse{
				err: ErrDecimalInvalidChar,
				str: d,
				col: i + 1,
			}
		}

		if decimalInt > MaxASCIIRange || decimalInt < MinASCIIRange {
			return nil, ErrTokenParse{
				err: ErrDecimalOutOfRange,
				str: d,
				col: i + 1,
			}
		}

		decimals[i] = int(decimalInt)
	}

	t := &Tokens{
		decimals: decimals,
	}

	return t, nil
}

func (t *Tokens) DecimalsToASCIIString() string {
	var result strings.Builder
	for _, v := range t.decimals {
		result.WriteByte(byte(v))
	}

	return result.String()
}

func (t *Tokens) String(sep string) string {
	tokens := make([]string, len(t.decimals))
	for i, v := range t.decimals {
		tokens[i] = strconv.FormatInt(int64(v), 10)
	}

	return strings.Join(tokens, sep)
}

func (e ErrTokenParse) Error() string {
	return fmt.Sprintf("%s (%s): col: %d", e.err, e.str, e.col)
}
