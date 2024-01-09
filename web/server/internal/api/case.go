package api

import (
	"main/web/server/internal/service"
	"main/web/server/internal/service/dto"
	"main/web/server/pkg/errpkg"
	"main/web/server/pkg/response"

	"github.com/gin-gonic/gin"
)

// GetCases get test cases.
func GetCases(c *gin.Context) {
	var request dto.CaseRequest
	if _err := c.ShouldBindQuery(&request); _err != nil {
		err := errpkg.NewMiddleErrorWithCause(_err, response.BindingError)
		response.Response(c, nil, err)
		return
	}

	data, err := service.NewCaseService().GetTestCases(request)
	response.Response(c, data, err)
}

// AddTestCase add a test case.
func AddTestCase(c *gin.Context) {
	var request dto.AddCaseRequestDto
	if _err := c.ShouldBindJSON(&request); _err != nil {
		err := errpkg.NewMiddleErrorWithCause(_err, response.BindingError)
		response.Response(c, nil, err)
		return
	}

	data, err := service.NewCaseService().AddTestCase(request)
	response.Response(c, data, err)
}
