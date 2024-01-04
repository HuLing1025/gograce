package response

import (
	"fmt"
	"main/web/server/pkg/errpkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SLSFlag enum WithSLSLog & WithNoSLSLog.
type SLSFlag = int

const (
	// WithSLSLog response with sls log.
	WithSLSLog SLSFlag = iota + 1
	// WithNoSLSLog response without sls log.
	WithNoSLSLog
)

// Body response body.
type Body struct {
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	Reason   string      `json:"reason,omitempty"`
	Data     interface{} `json:"data"`
	MetaData interface{} `json:"metadata"`
}

// Response response to client.
func Response(c *gin.Context, responseData interface{}, err errpkg.IError) {
	// success
	if err == nil {
		c.JSON(http.StatusOK, &Body{Code: http.StatusOK, Msg: "success", Data: responseData})
		return
	}
	// handle error response.
	responseErrHandle(c, err)
}

// responseErrHandle error response handle.
func responseErrHandle(c *gin.Context, err errpkg.IError) {
	code, exist := httpCode[err.GetErrorMsg()]
	if !exist {
		code = defaultHTTPCode[err.GetErrLevel()]
	}

	var responseData = Body{
		Code: code,
		Msg:  err.GetErrorMsg(),
	}
	// debug-mode will show the reason of error.
	if gin.Mode() == gin.DebugMode {
		responseData.Reason = fmt.Sprintf("%v", err.GetErrCause())
	}

	switch err.GetErrLevel() {
	case errpkg.ErrHighLevel:
		c.JSON(http.StatusInternalServerError, &responseData)
	case errpkg.ErrMiddleLevel, errpkg.ErrLowLevel:
		c.JSON(http.StatusBadRequest, &responseData)
	default:
		c.JSON(http.StatusInternalServerError, &responseData)
	}
}

// Pong response ping.
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, &Body{
		Code: 0,
		Msg:  "Pong",
	})
}
