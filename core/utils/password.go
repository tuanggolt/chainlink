package utils

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

const PasswordComplexityRequirements = `
Must have a minimum length of 16 characters
Must not comprise:
	A user's API email
`

const RequiredLen = 16

var (
	ErrPasswordMinLength = errors.Errorf("must be longer than %d characters", RequiredLen)
)

func VerifyPasswordComplexity(password string, disallowedStrings ...string) (merr error) {
	if len(password) < RequiredLen {
		merr = multierr.Append(merr, ErrPasswordMinLength)
	}

	for _, s := range disallowedStrings {
		if strings.Contains(password, s) {
			merr = multierr.Append(merr, errors.Errorf("password may not contain: %s", s))
		}
	}

	if merr != nil {
		merr = fmt.Errorf("password does not meet the requirements: %s", merr.Error())
	}

	return
}
