package grpc

import (
	"context"

	pkggrpc "github.com/Maxcillius/GoChat/pkg/grpc"
	db "github.com/Maxcillius/GoChat/platforms/db/db"

	"github.com/Maxcillius/GoChat/platforms/db/proto"
	"github.com/go-logr/logr"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {

	connection, err := db.New(ctx, logger)
	if err != nil {
		logger.Error(err, "Unable to connect to the database")
		return err
	}

	svc := &server{
		db: *connection,
	}

	return pkggrpc.NewServer(port, logger, func(s *grpc.Server) {
		proto.RegisterDBServiceServer(s, svc)
	}).Start(ctx)
}
