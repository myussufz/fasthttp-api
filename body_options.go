package api

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"reflect"

	"github.com/ajg/form"
)

func (c *Client) setBody(i interface{}) {
	var (
		errInvalidBody = errors.New("api: invalid body")
		errJSONBody    = errors.New("api: invalid json")
		errXMLBody     = errors.New("api: invalid xml")
		errFormBody    = errors.New("api: invalid form")
	)

	switch reflect.Indirect(reflect.ValueOf(i)).Kind() {
	case reflect.Struct:
	case reflect.Slice:
	case reflect.Map:
		goto routineMarshal

	default:
		c.err = errInvalidBody
		return
	}

routineMarshal:
	switch string(c.request.Header.ContentType()) {
	case ContentTypeJSON:
		data, err := json.Marshal(i)
		if err != nil {
			c.err = errJSONBody
			break
		}
		c.request.SetBody(data)

	case ContentTypeXML:
		data, err := xml.Marshal(i)
		if err != nil {
			c.err = errXMLBody
			break
		}
		c.request.SetBody(data)

	case ContentTypeXWWWFormURLEncoded:
		data, err := form.EncodeToString(i)
		if err != nil {
			c.err = errFormBody
			break
		}
		c.request.SetBody([]byte(data))
		break
	}

	return
}
