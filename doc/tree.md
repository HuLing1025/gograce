
- `GET` `localhost:8080/api/v1/tree`
- Structure Defines
  ```go
  // Directory directory.
  type Directory struct {
  	Name  string    `json:"directory"`
  	Files []FileDto `json:"files"`
  }q
  
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
  ```
- Response Demo
  ```json
    {
        "code": 200,
        "msg": "success",
        "data": [
            {
                "directory": "./cmd",
                "files": [
                    {
                        "name": "clean.go",
                        "functions": [
                            "init"
                        ],
                        "methods": null
                    },
                    {
                        "name": "init.go",
                        "functions": [
                            "init"
                        ],
                        "methods": null
                    },
                    {
                        "name": "report.go",
                        "functions": [
                            "init"
                        ],
                        "methods": null
                    },
                    {
                        "name": "root.go",
                        "functions": [
                            "Execute",
                            "init"
                        ],
                        "methods": null
                    },
                    {
                        "name": "start.go",
                        "functions": [
                            "init"
                        ],
                        "methods": null
                    }
                ]
            },
            {
                "directory": "./grace/executer/examples",
                "files": [
                    {
                        "name": "compare.go",
                        "functions": [
                            "Compare"
                        ],
                        "methods": null
                    }
                ]
            },
            {
                "directory": "./grace/executer",
                "files": [
                    {
                        "name": "executer.go",
                        "functions": [
                            "NewExecuter"
                        ],
                        "methods": [
                            {
                                "struct": "Executer",
                                "name": "Execute"
                            }
                        ]
                    }
                ]
            },
            {
                "directory": "./grace/file-processor",
                "files": [
                    {
                        "name": "source.go",
                        "functions": [
                            "NewSourceCodeProcessor",
                            "dfs"
                        ],
                        "methods": [
                            {
                                "struct": "SourceCodeProcessor",
                                "name": "BuildASTTree"
                            }
                        ]
                    },
                    {
                        "name": "yaml.go",
                        "functions": [
                            "NewYamlFileProcessor"
                        ],
                        "methods": [
                            {
                                "struct": "YamlFileProcessor",
                                "name": "CreateConfigFile"
                            },
                            {
                                "struct": "YamlFileProcessor",
                                "name": "AnalysisConfig"
                            }
                        ]
                    }
                ]
            },
            {
                "directory": "./web/server/internal/api",
                "files": [
                    {
                        "name": "home.go",
                        "functions": [
                            "GetTree"
                        ],
                        "methods": null
                    }
                ]
            },
            {
                "directory": "./web/server/internal/service",
                "files": [
                    {
                        "name": "home.go",
                        "functions": [
                            "NewHomeService",
                            "dfsFiles",
                            "dfsFuncs",
                            "getStructName"
                        ],
                        "methods": [
                            {
                                "struct": "HomeService",
                                "name": "GetTree"
                            }
                        ]
                    }
                ]
            },
            {
                "directory": "./web/server/pkg/errpkg",
                "files": [
                    {
                        "name": "error.go",
                        "functions": [
                            "newError",
                            "NewHighErrorWithCause",
                            "NewHighError",
                            "NewMiddleErrorWithCause",
                            "NewMiddleError",
                            "NewLowErrorWithCause",
                            "NewLowError"
                        ],
                        "methods": [
                            {
                                "struct": "Error",
                                "name": "Error"
                            },
                            {
                                "struct": "Error",
                                "name": "GetErrCause"
                            },
                            {
                                "struct": "Error",
                                "name": "GetErrLevel"
                            },
                            {
                                "struct": "Error",
                                "name": "GetErrorMsg"
                            }
                        ]
                    }
                ]
            },
            {
                "directory": "./web/server/pkg/response",
                "files": [
                    {
                        "name": "response.go",
                        "functions": [
                            "Response",
                            "responseErrHandle",
                            "Pong"
                        ],
                        "methods": null
                    }
                ]
            },
            {
                "directory": "./web/server",
                "files": [
                    {
                        "name": "server.go",
                        "functions": [
                            "NewServer",
                            "Run",
                            "shutdown"
                        ],
                        "methods": null
                    }
                ]
            },
            {
                "directory": ".",
                "files": [
                    {
                        "name": "gograce.go",
                        "functions": [
                            "main"
                        ],
                        "methods": null
                    }
                ]
            }
        ],
        "metadata": null
    }
  ```
---

[APIs](./API.md)