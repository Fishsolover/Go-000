package Service

import (
	"Goworks/Week02/Dao"
	"github.com/pkg/errors"
)

func ContinueWrapTheError() (string, error) {
	value, err := Dao.ErrorWithMessage()

	if err != nil {
		return "", errors.Wrap(err, "Wrap errors in service")
	}
	return value, nil
}
