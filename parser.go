package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type (
	Tokens struct {
		decimals []int64
	}
)

var (
	ErrParseNoValue = errors.New("no value")
)

func StringToDecimals(input string) (*Tokens, error) {
	if len(input) == 0 {
		return nil, ErrParseNoValue
	}

	decimals := make([]int64, len(input))

	for i, c := range input {
		decimals[i] = int64(c)
	}

	t := &Tokens{
		decimals: decimals,
	}

	return t, nil
}

func ParseDecimals(input, seperator string) (*Tokens, error) {
	input = strings.Trim(input, seperator)

	if len(input) == 0 {
		return nil, ErrParseNoValue
	}

	tokens := strings.Split(input, seperator)

	decimals := make([]int64, len(tokens))

	for i, d := range tokens {
		decimalInt, err := strconv.Atoi(d)

		if err != nil {
			return nil, fmt.Errorf("failed parsing token '%s'", d)
		}

		decimals[i] = int64(decimalInt)
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
		tokens[i] = strconv.FormatInt(v, 10)
	}

	return strings.Join(tokens, sep)
}
