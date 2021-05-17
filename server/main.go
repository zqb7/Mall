package main

import (
	"context"
	"fmt"
	"github.com/zqhhh/gomall/app"
	"github.com/zqhhh/gomall/model"
	"github.com/zqhhh/gomall/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	app.InitConfig()
	model.InitDB()
}
func main() {
	engine := router.InitEngine()
	server := &http.Server{
		Addr:    ":8001",
		Handler: engine,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	fmt.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println(err)
	}

}
