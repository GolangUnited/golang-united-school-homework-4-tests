package string_sum

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSum(t *testing.T) {
	testCases := map[string]struct {
		input    string
		output   string
		expErr   error
		numError bool
	}{
		"both operands positive":    {input: "24+55", output: "79", expErr: nil},
		"first operand negative":    {input: "-24+55", output: "31", expErr: nil},
		"second operand negative":   {input: "24-55", output: "-31", expErr: nil},
		"both operands negative":    {input: "-24-55", output: "-79", expErr: nil},
		"with whitespace":           {input: " -24 - 55 ", output: "-79", expErr: nil},
		"empty input":               {input: "", output: "", expErr: errorEmptyInput},
		"three operands":            {input: "11+23+43", output: "", expErr: errorNotTwoOperands},
		"one operand":               {input: "42", output: "", expErr: errorNotTwoOperands},
		"letters in first operand":  {input: "24c+55", output: "", expErr: &strconv.NumError{Func: "Atoi", Num: "24c", Err: strconv.ErrSyntax}, numError: true},
		"letters in second operand": {input: "24+55f", output: "", expErr: &strconv.NumError{Func: "Atoi", Num: "55f", Err: strconv.ErrSyntax}, numError: true},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			output, err := StringSum(tt.input)
			if tt.expErr != nil {
				if tt.numError {
					e := errors.Unwrap(err)
					assert.Equal(t, tt.expErr, e)
					assert.ErrorAs(t, err, &tt.expErr)
				} else {
					assert.ErrorIs(t, err, tt.expErr)
					assert.NotEqual(t, tt.expErr, err)
				}
			}
			assert.Equal(t, tt.output, output)
		})
	}
}
