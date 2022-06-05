package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"goHomework4/api"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {

	app, err := InitApp() // 使用wire生成的injector方法获取app对象
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	api.Routers(router)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return srv.ListenAndServe()
	})

	g.Go(func() error {

		select {
		case <-ctx.Done():

			timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			fmt.Println("shutting down server...")

			return srv.Shutdown(timeoutCtx)
		}
	})

	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		select {
		case sig := <-quit:
			return errors.Errorf("get os signal: %v", sig)
		}
	})

	fmt.Println(g.Wait())
	app.Close()
}
