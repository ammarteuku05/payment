package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"payment/controller"
	"payment/di"
	"syscall"
	"time"

	"github.com/labstack/echo"

	"payment/shared"
)

func main() {
	err := di.Container.Invoke(func(deps shared.Deps, ch controller.Holder) error {
		var (
			sig    = make(chan os.Signal, 1)
			app    = echo.New()
			parent = context.Background()
		)

		ch.SetupRoutes(app)
		go func() {
			if err := app.Start(":3000"); err != nil {
				deps.Logger.Errorf("failed to start server %s", err)
				sig <- syscall.SIGTERM
			}
		}()

		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		deps.Logger.Info("shutdown app and closing resources ...")
		ctx, cancel := context.WithTimeout(parent, 30*time.Second)
		_ = app.Shutdown(ctx)
		cancel()
		deps.Logger.Info("application terminated ...")
		return nil
	})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
