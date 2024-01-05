package dto

// CaseRequest case request.
type CaseRequest struct {
	Path       string `json:"path"`
	FileName   string `json:"file_name"`
	StructName string `json:"struct_name"`
	FuncName   string `json:"func_name"`
}

// CaseResponse case response.
type CaseResponse struct {
	// base information.
	BaseInfo FuncBaseInfo `json:"base_info" yaml:"BASE_INFO"`
	// dependent information.
	ExtValList []Param        `json:"ext_val_list" yaml:"EXT_VAL_LIST"`
	ExtFucList []FuncBaseInfo `json:"ext_fuc_list" yaml:"EXT_FUC_LIST"`

	// cases.
	Cases []Case `json:"cases" yaml:"CASES"`
	// func-level stubs.
	VarStubList  []VarStub  `json:"var_stub_list" yaml:"VAR_STUB_LIST"`
	FuncStubList []FuncStub `json:"func_stub_list" yaml:"FUNC_STUB_LIST"`
}

// Case test case.
type Case struct {
	Desc    string   `json:"desc" yaml:"DESC"`
	Asserts []Assert `json:"asserts" yaml:"ASSERTS"`

	// case-level stubs.
	VarStubList  []VarStub  `json:"var_stub_list" yaml:"VAR_STUB_LIST"`
	FuncStubList []FuncStub `json:"func_stub_list" yaml:"FUNC_STUB_LIST"`
}

// Assert assert.
type Assert struct {
	Input Param `json:"input" yaml:"INPUT"`
	Want  Param `json:"want" yaml:"WANT"`
}

// VarStub variable stub.
type VarStub struct {
	Name  string `json:"name" yaml:"NAME"`
	Value string `json:"value" yaml:"VALUE"`
}

// FuncStub function or method stub.
type FuncStub struct {
	StructName string `json:"struct_name" yaml:"STRUCT_NAME"`
	Name       string `json:"name" yaml:"NAME"`
	Statement  string `json:"statement" yaml:"STATEMENT"`
}

// Param param.
type Param struct {
	Name  string `json:"name" yaml:"NAME"`
	Type  string `json:"type" yaml:"TYPE"`
	Value string `json:"value" yaml:"VALUE"`
}

// FuncBaseInfo function.
type FuncBaseInfo struct {
	StructName string  `json:"struct_name" yaml:"STRUCT_NAME"`
	Name       string  `json:"name" yaml:"NAME"`
	Signature  string  `json:"signature" yaml:"SIGNATURE"`
	Statement  string  `json:"statement" yaml:"STATEMENT"`
	Inputs     []Param `json:"inputs" yaml:"INPUTS"`
	OutPuts    []Param `json:"outputs" yaml:"OUTPUTS"`
}

// ------------------------------- YAML CONFIG --------------------------------------------- //

// TestCaseConfig test case config.
type TestCaseConfig struct {
	Path     string `json:"path" yaml:"PATH"`
	FileName string `json:"file" yaml:"FILE_NAME"`

	// func-level stubs.
	VarStubList  []VarStub  `json:"var_stub_list" yaml:"VAR_STUB_LIST"`
	FuncStubList []FuncStub `json:"func_stub_list" yaml:"FUNC_STUB_LIST"`

	FunctionConfigs []FuncConfig `json:"function_configs" yaml:"FUNCTION_CONFIGS"`
}

// FuncConfig function or method config.
type FuncConfig struct {
	StructName string `json:"struct_name" yaml:"STRUCT_NAME"`
	Name       string `json:"name" yaml:"NAME"`
	Cases      []Case `json:"cases" yaml:"CASES"`
}
