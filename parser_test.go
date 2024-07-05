package main

import (
	"errors"
	"testing"
)

func TestTokens(t *testing.T) {

	// Test empty tokens.
	tokens := Tokens{}

	if len(tokens.String(" ")) != 0 {
		t.Errorf("expected empty string in decimal: %s", tokens.String(" "))
	}

	if len(tokens.DecimalsToASCIIString()) != 0 {
		t.Errorf("expected empty string in ASCII: %s", tokens.DecimalsToASCIIString())
	}

	// Test output of basic strings.
	tokens = Tokens{decimals: []int{123, 125}}

	if tokens.String(" ") != "123 125" {
		t.Errorf("expected 123 125: %s", tokens.String(" "))
	}

	if tokens.DecimalsToASCIIString() != "{}" {
		t.Errorf("expected empty string: %s", tokens.DecimalsToASCIIString())
	}
}

func TestParseDecimals(t *testing.T) {

	// test empty value edge case.
	tokens, err := ParseDecimals("", " ")

	if !errors.Is(err, ErrParseNoValue) {
		t.Errorf("expected ErrParseNoValue, but got: %s", err)
	}

	if tokens != nil {
		t.Error("expected *Token to be nil")
	}

	// test for valid decimals
	tokens, err = ParseDecimals("123 125", " ")

	if err != nil {
		t.Errorf("unexpected error ParseDecimals, got: %s", err)
	}

	if tokens.String(" ") != "123 125" {
		t.Errorf("expected 123 125: %s", tokens.String(" "))
	}

	if tokens.DecimalsToASCIIString() != "{}" {
		t.Errorf("expected empty string: %s", tokens.DecimalsToASCIIString())
	}

	// test for invalid tokens
	tokens, err = ParseDecimals("123a 125", " ")

	if errors.Is(err, ErrTokenParse{}) {
		t.Errorf("unexpected error ParseDecimals, got: %s", err)
	}

	if tokens != nil {
		t.Error("expected *Token to be nil")
	}

	// test for max out of range decimal tokens
	tokens, err = ParseDecimals("200 125", " ")

	if errors.Is(err, ErrTokenParse{}) {
		t.Errorf("unexpected error ParseDecimals, got: %s", err)
	}

	if tokens != nil {
		t.Error("expected *Token to be nil")
	}

	// test for negetive range decimal tokens
	tokens, err = ParseDecimals("-123 125", " ")

	if errors.Is(err, ErrTokenParse{}) {
		t.Errorf("unexpected error ParseDecimals, got: %s", err)
	}

	if tokens != nil {
		t.Error("expected *Token to be nil")
	}
}

func TestStringToDecimals(t *testing.T) {

	// test empty value edge case.
	tokens, err := StringToDecimals("")

	if !errors.Is(err, ErrParseNoValue) {
		t.Errorf("expected ErrParseNoValue, but got: %s", err)
	}

	if tokens != nil {
		t.Error("expected *Token to be nil")
	}

	// test for valid decimals
	tokens, err = StringToDecimals("{}")

	if err != nil {
		t.Errorf("unexpected error StringToDecimals, got: %s", err)
	}

	if tokens.String(" ") != "123 125" {
		t.Errorf("expected 123 125: %s", tokens.String(" "))
	}

	if tokens.DecimalsToASCIIString() != "{}" {
		t.Errorf("expected empty string: %s", tokens.DecimalsToASCIIString())
	}
}
