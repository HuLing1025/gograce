package server

import (
	"context"
	"fmt"
	"log"
	"main/web/server/internal/api"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// NewServer create server.
func NewServer() *gin.Engine {
	engine := gin.New()
	engine.Use(cors())

	engine.Static("/grace", "./web/client/go-grace/dist")

	engine.GET("", func(c *gin.Context) {
		c.Redirect(302, "http://localhost:8080/grace")
	})

	v1 := engine.Group("/v1")
	{
		v1.GET("tree", api.GetTree)
		v1.GET("case", api.GetCases)
		v1.POST("case", api.AddTestCase)
	}

	return engine
}

// Run run server
func Run(routerHandler *gin.Engine) {
	httpSrv := &http.Server{
		Addr:         ":8080",
		Handler:      routerHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("ListenAndServe error:", err)
		}
	}()

	shutdown(httpSrv)
}

// shutdown shutdown server.
func shutdown(httpSrv *http.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case sig := <-sigs:
		fmt.Println(fmt.Sprintf("catch signal.Notify,sigs:%v", sig))
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		if err := httpSrv.Shutdown(ctx); err != nil {
			fmt.Println(fmt.Sprintf("catch signal.shutdown,err::%v", err))
		}
		fmt.Println("gograce shutdown...")
	}

	time.Sleep(2 * time.Second)
}

// cors 跨域配置
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		//c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Headers", "*,content-type,x-token")
		c.Header("Access-Control-Expose-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "86400")
		//放行索引options
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		//处理请求
		c.Next()
	}
}
