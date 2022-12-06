package country

import (
	"github.com/NpoolPlatform/message/npool/g11n/mgr/v1/country"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	country.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	country.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
