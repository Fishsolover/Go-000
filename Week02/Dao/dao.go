package Dao

import (
	"Goworks/Week02/DB"
	"github.com/pkg/errors"
)

func ErrorWithMessage() (string, error) {
	value, err := DB.ThrowError()

	if err != nil {
		return "", errors.WithMessage(err, "Wrap errors in DAO")
	}
	return value, nil
}
