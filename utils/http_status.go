package utils

import "errors"

func GetHttpStatus(err error) int {
	switch {
	case errors.Is(err, ErrAuthFailed):
		return 401
	default:
		return 500
	}
}
