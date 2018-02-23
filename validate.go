package api

import (
	"errors"
	"net/url"
)

func validateURL(u string) error {
	var (
		errURLRequired = errors.New("api: url is required")
		errInvalidURL  = errors.New("api: invalid url")
	)

	if len(u) == 0 {
		return errURLRequired
	}

	_, err := url.Parse(u)
	if err != nil {
		return errInvalidURL
	}

	return nil

}
