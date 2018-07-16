package api

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/ajg/form"
	json "github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
)

var methods = make(map[string]bool)

func init() {
	for _, m := range [...]string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPatch,
		http.MethodPut,
		http.MethodConnect,
		http.MethodHead,
		http.MethodTrace,
		http.MethodDelete,
		http.MethodOptions,
	} {
		methods[m] = true
	}
}

// Option :
type Option struct {
	ContentType string
	Method      string
	Headers     map[string]string
	Body        interface{}
}

var (
	defaultOption = Option{
		ContentType: ContentTypeJSON,
		Method:      http.MethodGet,
	}
)

// Client :
type Client struct {
	method      string
	headers     *fasthttp.RequestHeader
	client      *fasthttp.Client
	contentType string
	timeTaken   time.Duration
	body        []byte
	err         error
}

// Fetch :
func Fetch(url string, option ...Option) *Client {
	startAt := time.Now()
	c := new(Client)
	c.client = new(fasthttp.Client)
	c.method = http.MethodGet
	c.contentType = ContentTypeJSON
	c.headers = new(fasthttp.RequestHeader)
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
		c.timeTaken = time.Since(startAt)
	}()

	if err := validateURL(url); err != nil {
		c.err = err
		return c
	}

	c.headers.SetRequestURI(url)
	opt := defaultOption
	if len(option) > 0 {
		opt = option[0]
	}

	if opt.Method != "" {
		c.method = opt.Method
	}
	if opt.ContentType != "" {
		c.contentType = opt.ContentType
	}

	if _, isOk := methods[c.method]; !isOk {
		c.err = fmt.Errorf("api: invalid request method %q", c.method)
		return c
	}
	c.headers.SetMethod(c.method)
	c.headers.SetContentType(c.contentType)
	for key, value := range opt.Headers {
		c.headers.Add(key, value)
	}
	c.headers.CopyTo(&request.Header)
	if opt.Body != nil {
		bb, err := c.toByte(opt.Body)
		if err != nil {
			c.err = err
			return c
		}
		request.SetBody(bb)
	}
	if err := c.client.Do(request, response); err != nil {
		c.err = err
		return c
	}

	c.body = response.Body()
	return c
}

func (c *Client) toByte(i interface{}) ([]byte, error) {
	switch string(c.contentType) {
	case ContentTypeXML:
		return xml.Marshal(i)
	case ContentTypeXWWWFormURLEncoded:
		str, err := form.EncodeToString(i)
		return []byte(str), err
	default:
		return json.Marshal(i)
	}
}

// ToString :
func (c *Client) ToString() (string, error) {
	if c.err != nil {
		return "", c.err
	}
	return string(c.body), nil
}

// ToXML :
func (c *Client) ToXML(i interface{}) error {
	if c.err != nil {
		return c.err
	}

	var (
		errUnableUnmarshalXML  = errors.New("api: unable to unmarshal the xml")
		errStructShouldPointer = errors.New("api: struct should be pointer")
	)

	if reflect.ValueOf(i).Kind() != reflect.Ptr {
		return errStructShouldPointer
	}

	if err := xml.Unmarshal(c.body, i); err != nil {
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

	if err := json.Unmarshal(c.body, i); err != nil {
		return errUnableUnmarshalJSON
	}

	return nil
}
