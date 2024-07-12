package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var Aapplication = NewApp()

// APP:  应用程序APP类型
type APP struct {
	engine *gin.Engine
	server http.Server
}

// NewApp:  构建APP
func NewApp() *APP {
	a := &APP{
		engine: gin.Default(),
	}
	return a
}

// RegisterGroup: 创建RouterGroup
func (a *APP) CreateGroup(relativePath string) *gin.RouterGroup {
	return a.engine.Group(relativePath)
}

// Start: 启动APP
func (a *APP) Start(addr string) {
	a.server = http.Server{
		Addr:    addr,
		Handler: a.engine,
	}
	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server listen error: ", err)
		}
	}()
	<-make(chan int) // 阻塞代码
}

// Stop: 关闭APP方法, 默认处于阻塞状态，但是当收到os.Signal 时就停止阻塞，此时就会在指定时间内关闭。
func (a *APP) StartWithSupportStop(addr string, timeout time.Duration) {
	go a.Start(addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞，当执行

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown error")
	} else {
		log.Println("\nserver shutdown!")
	}
}
