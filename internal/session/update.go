package session

import (
	"context"
	"log"

	pb "github.com/notobject/sessiond/api/rpc"
)

// UpdateSession 检查Session是否有效
func (*DefaultSessionServer) UpdateSession(ctx context.Context, req *pb.UpdateSessionReq) (*pb.UpdateSessionRsp, error) {

	log.Println("CheckSession Called...")
	return &pb.UpdateSessionRsp{Ret: 0}, nil
}
