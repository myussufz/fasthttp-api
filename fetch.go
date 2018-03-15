package api

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"reflect"

	"github.com/valyala/fasthttp"
)

// Options :
type Options struct {
	ContentType string
	Method      string
	Headers     map[string]string
	Body        interface{}
	// Credentials string
}

// Client :
type Client struct {
	client   *fasthttp.Client
	request  *fasthttp.Request
	response *fasthttp.Response
	err      error
}

// Fetch :
func Fetch(url string, options *Options) *Client {
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()

	c := new(Client)
	c.request = request
	c.response = response

	if err := validateURL(url); err != nil {
		return c
	}

	c.request.SetRequestURI(url)

	if options != nil {
		c.setOptions(options)
	} else {
		c.request.Header.SetMethod(defaultMethod)
		c.request.Header.SetContentType(defaultContentType)
	}

	c.client = new(fasthttp.Client)

	if err := c.client.Do(c.request, c.response); err != nil {
		c.err = err
		return c
	}

	return c
}

// ToString :
func (c *Client) ToString() (string, error) {
	if c.err != nil {
		return "", c.err
	}

	return string(c.response.Body()), nil
}

// ToXML :
func (c *Client) ToXML(i interface{}) error {
	var (
		errUnableUnmarshalXML  = errors.New("api: unable to unmarshal the xml")
		errStructShouldPointer = errors.New("api: struct should be pointer")
	)

	if c.err != nil {
		return c.err
	}

	if reflect.ValueOf(i).Kind() != reflect.Ptr {
		return errStructShouldPointer
	}

	if err := xml.Unmarshal(c.response.Body(), i); err != nil {
		return errUnableUnmarshalXML
	}

	return nil
}

// ToJSON :
func (c *Client) ToJSON(i interface{}) error {
	var (
		errUnableUnmarshalJSON = errors.New("api: unable to unmarshal the json")
		errStructShouldPointer = errors.New("api: struct should be pointer")
	)

	if c.err != nil {
		return c.err
	}

	if reflect.ValueOf(i).Kind() != reflect.Ptr {
		return errStructShouldPointer
	}

	if err := json.Unmarshal(c.response.Body(), i); err != nil {
		return errUnableUnmarshalJSON
	}

	return nil
}
