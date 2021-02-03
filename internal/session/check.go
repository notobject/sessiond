package session

import (
	"context"
	"log"

	pb "github.com/notobject/sessiond/api/rpc"
)

// CheckSession 检查Session是否有效
func (*DefaultSessionServer) CheckSession(ctx context.Context, req *pb.CheckSessionReq) (*pb.CheckSessionRsp, error) {

	log.Println("CheckSession Called...")
	return &pb.CheckSessionRsp{Ret: 0, IsValid: true}, nil
}
