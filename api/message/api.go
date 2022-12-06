package message

import (
	"github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	message.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	message.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
