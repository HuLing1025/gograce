package response

import (
	"main/web/server/pkg/errpkg"
	"net/http"
)

// httpCode response code mapping.
var httpCode = map[string]int{
	// System error code.
	DbError:               http.StatusInternalServerError, // 500
	RedisError:            http.StatusInternalServerError, // 500
	BindingError:          http.StatusBadRequest,          // 400
	PermissionDeniedError: http.StatusUnauthorized,        // 401
	TokenError:            http.StatusUnauthorized,        // 401
	TokenExpiredError:     http.StatusUnauthorized,        // 401
	JSONError:             http.StatusInternalServerError, // 500

	// Customer error code mapping.
	UserNotExist:      10001,
	WrongPassword:     10002,
	EmailFormatError:  10003,
	UserHasRegistered: 10004,

	// TODO more customer code mapping definitions...
}

// defaultHTTPCode default response code.
var defaultHTTPCode = map[errpkg.ErrorLevel]int{
	errpkg.ErrHighLevel:   500,
	errpkg.ErrMiddleLevel: 400,
	errpkg.ErrLowLevel:    400,
}
