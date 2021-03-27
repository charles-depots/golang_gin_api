package status

const (
	StatusOK         = 200111
	BadRequest       = 404102
	RequestRateLimit = 201102
)

var statusMessage = map[int]string{
	StatusOK:         "OK",
	BadRequest:       "Request was missing the 'redirect_uri' parameter. ",
	RequestRateLimit: "The current flow is too big, please try again later",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func GetStatusMessage(code int) string {
	return statusMessage[code]
}
