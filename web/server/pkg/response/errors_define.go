package response

// system errors.
const (
	DbError      = "database error"
	RedisError   = "redis error"
	BindingError = "binding error"
	JSONError    = "json serialization error"

	// JWT
	PermissionDeniedError = "permission denied"
	TokenExpiredError     = "token expired"
	TokenError            = "token error"

	// copy error
	CopyError = "copy error"

	// email error
	EmailFormatError = "The email address is invalid"

	// TODO more system error definitions...
)

// customer errors.
const (
	UserNotExist      = "user not exist"
	UserHasRegistered = "user has registered"
	WrongPassword     = "wrong password"

	// TODO more customer error definitions...
)
