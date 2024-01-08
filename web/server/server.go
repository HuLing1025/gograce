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
	engine.Static("/grace", "./web/client/go-grace/dist")

	engine.GET("", func(c *gin.Context) {
		c.Redirect(302, "http://localhost:8080/grace")
	})

	v1 := engine.Group("/v1")
	{
		v1.GET("tree", api.GetTree)
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
