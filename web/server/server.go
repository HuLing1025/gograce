package server

import (
	"context"
	"fmt"
	"log"
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
	engine.Static("", "./web/client/")

	// engine.GET("grace")

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
		fmt.Printf(fmt.Sprintf("捕获信号signal.Notify,sigs:%v", sig))
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := httpSrv.Shutdown(ctx); err != nil {
			fmt.Printf(fmt.Sprintf("捕获信号signal.shutdown,err::%v", err))
		}
		fmt.Printf("gograce shutdown...")
	}

	time.Sleep(3 * time.Second)
}
