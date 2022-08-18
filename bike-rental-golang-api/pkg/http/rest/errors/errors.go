package errors

// Error strings to be used for http response
const (
	ErrEmptyRequestBody = "request body is empty"
	ErrParsingJson      = "internal error while parsing json"
	ErrSessionID        = "invalid SessionID in header"
	ErrGetRentQuery     = "getRents needs userID and status as query parameter"
)
