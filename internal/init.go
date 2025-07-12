package internal

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"ad-api/config"
	"ad-api/internal/app/ad/v1"
	adV1 "ad-api/internal/pkg/pb/ad/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func (a *App) initControllers(_ context.Context) error {
	// init services
	adV1.RegisterAdServiceServer(a.grpcServer, ad.NewAdService())
	return nil
}

func (a *App) initGrpcServer(ctx context.Context) error {
	// init listener
	l, err := net.Listen("tcp", net.JoinHostPort(config.Instance().GrpcServer.Host, config.Instance().GrpcServer.Port))
	if err != nil {
		return err
	}
	a.grpcListener = l
	a.publicCloser.Add(a.grpcListener.Close)

	// init server
	a.grpcServer = grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: config.Instance().GrpcServer.MaxConnectionIdle,
		MaxConnectionAge:  config.Instance().GrpcServer.MaxConnectionAge,
		Time:              config.Instance().GrpcServer.Time,
		Timeout:           config.Instance().GrpcServer.Timeout,
	}))

	reflection.Register(a.grpcServer)

	a.publicCloser.Add(func() error {
		gracefulCtx, cancel := context.WithTimeout(context.Background(), config.Instance().Graceful.Timeout)
		defer cancel()

		done := make(chan struct{})
		go func() {
			a.grpcServer.GracefulStop()
			close(done)
		}()

		select {
		case <-done:
			slog.Warn("grpc: server gracefully stopped")
		case <-gracefulCtx.Done():
			err = fmt.Errorf("grpc: error while graceful shutdown server: %w", ctx.Err())
			a.grpcServer.Stop()
			return fmt.Errorf("grpc: stopped: %w", err)
		}
		return nil
	})

	return nil
}
