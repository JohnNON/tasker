package model

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

var (
	rule = regexp.MustCompile("^[a-z0-9]+$")
)

func required() validation.RuleFunc {
	return func(value interface{}) error {
		val, ok := value.(string)
		if !ok {
			return errors.New("must be a string")
		}

		if !rule.MatchString(val) {
			return errors.New("must contain only a-z and 0-9 symbols")
		}

		return nil
	}
}
