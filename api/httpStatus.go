package api

type Error struct {
	Message string `json:"message"`
}

type APIError struct {
	Status Status `json:"status"`
	Error  Error  `json:"error"`
}

func BuildAPIError(code int, message string) APIError {
	return APIError{
		Status: Status{
			Code: code,
			Ok:   false,
		},
		Error: Error{
			Message: message,
		},
	}
}

func BuildAPISuccess(code int) Status {
	return Status{
		Code: code,
		Ok:   true,
	}
}

func OK() Status {
	return BuildAPISuccess(200)
}

func Created() Status {
	return BuildAPISuccess(201)
}

func Accepted() Status {
	return BuildAPISuccess(202)
}

func NonAuthoritativeInformation() Status {
	return BuildAPISuccess(203)
}

func NoContent() Status {
	return BuildAPISuccess(204)
}

func ResetContent() Status {
	return BuildAPISuccess(205)
}

func PartialContent() Status {
	return BuildAPISuccess(206)
}

func MultiStatus() Status {
	return BuildAPISuccess(207)
}

func AlreadyReported() Status {
	return BuildAPISuccess(208)
}

func IMUsed() Status {
	return BuildAPISuccess(226)
}

func BadRequest() APIError {
	return BuildAPIError(400, "Bad request")
}

func Unauthorized() APIError {
	return BuildAPIError(401, "Unauthorized")
}

// Unlike 401 Unauthorized, the client's identity is known to the server.
func Forbidden() APIError {
	return BuildAPIError(403, "Forbidden")
}

func NotFound() APIError {
	return BuildAPIError(404, "Not found")
}

func MethodNotAllowed() APIError {
	return BuildAPIError(405, "Method Not Allowed")
}

func Conflict() APIError {
	return BuildAPIError(409, "Conflict")
}

func UnsupportedMediaType() APIError {
	return BuildAPIError(415, "Unsupported Media Type")
}

func IAmTeapot() APIError {
	return BuildAPIError(418, "I'm a teapot")
}

func Locked() APIError {
	return BuildAPIError(423, "Locked")
}

func TooManyRequests() APIError {
	return BuildAPIError(429, "Too Many Requests")
}

func Internal(msg *string) APIError {
	if msg == nil {
		return BuildAPIError(500, "Internal")
	}

	return BuildAPIError(500, *msg)
}

func BadGateway() APIError {
	return BuildAPIError(502, "Bad Gateway")
}

func ServiceUnavailable() APIError {
	return BuildAPIError(503, "Service Unavailable")
}
