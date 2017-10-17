package main

import "testing"
import "strings"

func TestSign(t *testing.T) {
	//TODO: write unit test cases for sign()
	//use strings.NewReader() to get an io.Reader
	//interface over a simple string
	//https://golang.org/pkg/strings/#NewReader
	cases := []struct {
		input          string
		signingKey     string
		expectedOutput string
	}{
		{
			input:          "a value to sign",
			signingKey:     "testkey",
			expectedOutput: "mDuAdcBb01c_3z64qX4yc6g1D6dXnzuGjYmOImUd-h0=",
		},
		{
			input:          "a value to sign 2",
			signingKey:     "testkey",
			expectedOutput: "lA0D8_YIZzTiM0VWQxPBp_jhvljTZLs0QPEFUycdkIw=",
		},
	}

	for _, c := range cases {
		output, err := sign(c.signingKey, strings.NewReader(c.input))
		if err != nil {
			t.Errorf("error signing: %v", err)
		}
		if output != c.expectedOutput {
			t.Errorf("expected %s but got %s", c.expectedOutput, output)
		}
	}
}

func TestVerify(t *testing.T) {
	//TODO: write until test cases for verify()
	//use strings.NewReader() to get an io.Reader
	//interface over a simple string
	//https://golang.org/pkg/strings/#NewReader
}
