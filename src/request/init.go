package request

// Req holds the main object
type Req struct {
	HTTPUserAgent string
}

// Init does exactly what it says
func Init(HTTPUserAgent string) (req Req) {
	return Req{
		HTTPUserAgent: HTTPUserAgent,
	}
}
