package api

import (
	"errors"
	"strings"
)

func (c *Client) setOptions(options *Options) {
	var (
		errMethodNotSupported = errors.New("api: Method not supported")
	)

	for key, value := range options.Headers {
		c.request.Header.Add(key, value)
	}

	if len(options.ContentType) > 0 {
		c.request.Header.SetContentType(options.ContentType)
	} else {
		c.request.Header.SetContentType(defaultContentType)
	}

	if len(string(options.Method)) > 0 {
		method := strings.ToUpper(strings.TrimSpace(options.Method))

		if _, isExist := methodMap[method]; !isExist {
			c.err = errMethodNotSupported
		} else {
			c.request.Header.SetMethod(method)
		}
	} else {
		c.request.Header.SetMethod(defaultMethod)
	}

	// set the body if not equal to nil
	if options.Body != nil {
		c.setBody(options.Body)
	}
}
