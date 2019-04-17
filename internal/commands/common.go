package commands

import (
	"errors"
	"fmt"
	"strconv"
)

func validateNumbers(a, b string) (int, int, error) {
	lhn, err := strconv.Atoi(a)
	if err != nil {
		return 0, 0, errors.New(fmt.Sprintf("invalid first number %s", a))
	}
	rhn, err := strconv.Atoi(b)
	if err != nil {
		return 0, 0, errors.New(fmt.Sprintf("invalid second number %s", b))
	}
	return lhn, rhn, nil
}
