package session

import (
	"context"
	"log"
	"time"

	pb "github.com/notobject/sessiond/api/rpc"
	"github.com/notobject/sessiond/pkg/dao"
)

// CreateSession 创建Session
func (*DefaultSessionServer) CreateSession(ctx context.Context, req *pb.CreateSessionReq) (*pb.CreateSessionRsp, error) {

	client, err := dao.NewDao().GetRedisClient()
	if err != nil {
		log.Fatal(err)
	}

	// 生成sessionid
	sessionID := "ABCDEFG1234"

	pipe := client.Pipeline()

	// 保存user -> sessionID的映射
	pipe.HSet("vdc:services:session:user_session_mapping", req.UserName, sessionID)
	// 保存sessionID -> user的映射
	pipe.HSet("vdc:services:session:session_user_mapping", sessionID, req.UserName)
	// 保存会话数据
	pipe.HSet("vdc:services:session:session_info_mapping", sessionID, req.UserInfo)
	// 用户数+1
	pipe.Incr("vdc:services:session:total")

	_, err = pipe.Exec()
	if err != nil {
		return nil, err
	}
	return &pb.CreateSessionRsp{SessionID: sessionID, Timestamp: time.Now().UTC().Unix()}, nil
}
