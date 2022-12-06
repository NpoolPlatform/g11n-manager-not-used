package applang

import (
	"github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	applang.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	applang.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
