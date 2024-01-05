package dto

// Directory directory.
type Directory struct {
	Name  string    `json:"directory"`
	Files []FileDto `json:"files"`
}

// FileDto file dto.
type FileDto struct {
	Name      string      `json:"name"`
	Functions []string    `json:"functions"`
	Methods   []MethodDto `json:"methods"`
}

// MethodDto method dto.
type MethodDto struct {
	Struct string `json:"struct"`
	Name   string `json:"name"`
}
