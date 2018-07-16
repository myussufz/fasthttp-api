package api

// Content Type
const (
	ContentTypeJSON               = "application/json"
	ContentTypeXML                = "application/xml"
	ContentTypeXWWWFormURLEncoded = "application/x-www-form-urlencoded"
)

var (
	contentTypeMap = map[string]bool{
		ContentTypeJSON:               true,
		ContentTypeXML:                true,
		ContentTypeXWWWFormURLEncoded: true,
	}
)
