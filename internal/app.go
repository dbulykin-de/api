package internal

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"syscall"

	"ad-api/internal/pkg/closer"

	"google.golang.org/grpc"
)

// App application
type App struct {
	grpcServer *grpc.Server

	grpcListener net.Listener

	publicCloser *closer.Closer
}

// New конструктор
func New(ctx context.Context) *App {
	app := &App{publicCloser: closer.New(syscall.SIGTERM, syscall.SIGINT)}
	err := app.init(ctx)
	if err != nil {
		log.Fatalf("[APP] Не удалось инициализировать приложение: %s", err.Error())
	}

	return app
}

// Run запуск приложения
func (a *App) Run(_ context.Context) {
	if a.grpcServer != nil {
		go func() {
			if err := a.grpcServer.Serve(a.grpcListener); err != nil {
				slog.Error(fmt.Sprintf("grpc: %s", err.Error()))
				a.publicCloser.CloseAll()
			}
		}()
	}

	a.publicCloser.Wait()

	closer.CloseAll()
}

func (a *App) init(ctx context.Context) error {
	initFuncs := []func(context.Context) error{
		a.initGrpcServer,
		a.initControllers,
	}

	for _, f := range initFuncs {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
