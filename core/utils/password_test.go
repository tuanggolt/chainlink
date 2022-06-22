package utils_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink/core/utils"
	"github.com/stretchr/testify/assert"
)

func TestVerifyPasswordComplexity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		password string
		errors   []error
	}{
		{"thispasswordislongenough", []error{}},
		{"exactlyrightlen1", []error{}},
		{"notlongenough", []error{utils.ErrPasswordMinLength}},
	}

	for _, test := range tests {
		test := test

		t.Run(test.password, func(t *testing.T) {
			t.Parallel()

			err := utils.VerifyPasswordComplexity(test.password)
			if len(test.errors) == 0 {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "password does not meet the requirements")
				for _, subErr := range test.errors {
					assert.ErrorContains(t, err, subErr.Error())
				}
			}
		})
	}
}
