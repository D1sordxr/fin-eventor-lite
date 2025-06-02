package errors

import "errors"

func In(err error, errs ...error) bool {
	if err == nil {
		return false
	}

	for _, e := range errs {
		if errors.Is(err, e) {
			return true
		}
	}

	return false
}
