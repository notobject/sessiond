package session

import (
	"log"
	"net"

	pb "github.com/notobject/sessiond/api/rpc"
	"github.com/notobject/sessiond/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// DefaultSessionServer 会话服务
type DefaultSessionServer struct {
	pb.UnimplementedSessionServer
	logger log.Logger
}

// Start session服务
func Start(cfg *config.Config) {
	listener, err := net.Listen("tcp", cfg.Addr.RPC)
	if err != nil {
		return
	}
	server := grpc.NewServer()

	pb.RegisterSessionServer(server, &DefaultSessionServer{})

	reflection.Register(server)
	go server.Serve(listener)
	log.Println("start grpc service at ", cfg.Addr.RPC)
}
