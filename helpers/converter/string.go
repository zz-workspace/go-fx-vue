package converter

import (
	"strconv"
	"errors"
)

func StringToInt(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, errors.New("Convert string to int error")
	}
	return num, nil
}