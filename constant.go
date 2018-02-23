package api

// Method Type
const (
	MethodPOST   = "POST"
	MethodGET    = "GET"
	MethodPUT    = "PUT"
	MethodPATCH  = "PATCH"
	MethodDELETE = "DELETE"
)

// Content Type
const (
	ContentTypeJSON = "application/json"
	ContentTypeXML  = "application/xml"
	// ContentTypeXWWWFormURLEncoded = "application/x-www-form-urlencoded"
)

var (
	methodMap = map[string]bool{
		MethodPOST:   true,
		MethodGET:    true,
		MethodPATCH:  true,
		MethodPUT:    true,
		MethodDELETE: true,
	}

	contentTypeMap = map[string]bool{
		ContentTypeJSON: true,
		ContentTypeXML:  true,
		// ContentTypeXWWWFormURLEncoded: true,
	}
)

// Default Constant
const (
	defaultMethod      = MethodGET
	defaultContentType = ContentTypeJSON
)
