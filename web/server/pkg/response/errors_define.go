package response

// system errors.
const (
	BindingError = "binding error"
	JSONError    = "json error"

	ReadYAMLConfigError = "read yaml config error"

// TODO more system error definitions...
)

// customer errors.
const (
	BuildASTError       = "build ast fail"
	SearchASTError      = "search ast error"
	SearchFunctionError = "search function error"
	SearchMethodError   = "search method error"
	StatementError      = "statement error"
	SyncConfigError     = "sync test config error"

// TODO more customer error definitions...
)
