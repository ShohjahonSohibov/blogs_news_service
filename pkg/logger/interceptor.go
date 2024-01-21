package logger

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func LogUnaryInterceptor(l LoggerI) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()
		l.Info("request received",
			zap.String("method", info.FullMethod),
			zap.Any("request", req),
		)
		res, err := handler(ctx, req)
		elapsed := time.Since(start)
		l.Info("response sent",
			zap.Any("response", res),
			zap.Duration("elapsed", elapsed),
			zap.Error(err),
		)
		return res, err
	}
}
