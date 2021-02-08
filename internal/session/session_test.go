package session

import (
	rpc "github.com/notobject/sessiond/api/rpc"
	"google.golang.org/grpc"
)

func testNewSessionClient() rpc.SessionClient {
	conn, _ := grpc.Dial("localhost:8080", grpc.WithInsecure())
	return rpc.NewSessionClient(conn)
}
