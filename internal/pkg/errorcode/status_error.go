package errorcode

// ErrorCode type for error
type ErrorCode int

const (
	// Internal internal error
	Internal ErrorCode = 500

	//NotFound resource not found error
	NotFound ErrorCode = 404
)
