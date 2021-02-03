package session

import (
	"context"
	"log"

	pb "github.com/notobject/sessiond/api/rpc"
)

// DeleteSession 通过sessionID删除会话
func (*DefaultSessionServer) DeleteSession(
	ctx context.Context,
	req *pb.DeleteSessionReq) (*pb.DeleteSessionRsp, error) {

	log.Println("DeleteSession Called...")
	return &pb.DeleteSessionRsp{Ret: 0}, nil
}

// DeleteSessionBatch 批量删除会话
func (*DefaultSessionServer) DeleteSessionBatch(ctx context.Context, req *pb.DeleteSessionBatchReq) (*pb.DeleteSessionBatchRsp, error) {

	log.Println("DeleteSessionBatch Called...")
	retMap := make(map[string]int32, len(req.SessionIDs))

	for _, sessionID := range req.SessionIDs {
		retMap[sessionID] = 0
	}
	return &pb.DeleteSessionBatchRsp{Ret: 0, RetMap: retMap}, nil
}
