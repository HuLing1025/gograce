package response

import (
	"main/web/server/pkg/errpkg"
)

// httpCode response code mapping.
var httpCode = map[string]int{
	// System error code.
	BindingError: 400,

	// Customer error code mapping.
	BuildASTError:       10001,
	SearchASTError:      10002,
	ReadYAMLConfigError: 10003,
	SearchFunctionError: 10004,
	SearchMethodError:   10005,
	StatementError:      10006,
	SyncConfigError:     10007,

	// TODO more customer code mapping definitions...
}

// defaultHTTPCode default response code.
var defaultHTTPCode = map[errpkg.ErrorLevel]int{
	errpkg.ErrHighLevel:   500,
	errpkg.ErrMiddleLevel: 400,
	errpkg.ErrLowLevel:    400,
}
