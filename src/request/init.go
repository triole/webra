package request

// Req holds the main object
type Req struct {
	HTTPUserAgent string
	Timeout       int
}

// Init does exactly what it says
func Init(HTTPUserAgent string, timeout int) (req Req) {
	return Req{
		HTTPUserAgent: HTTPUserAgent,
		Timeout:       timeout,
	}
}
