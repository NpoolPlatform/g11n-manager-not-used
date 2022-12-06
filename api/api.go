package api

import (
	"context"

	g11n "github.com/NpoolPlatform/message/npool/g11n/mgr/v1"

	"github.com/NpoolPlatform/g11n-manager/api/country"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	g11n.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	g11n.RegisterManagerServer(server, &Server{})
	country.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := g11n.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := country.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
