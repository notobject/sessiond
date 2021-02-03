package session

import (
	"context"
	"log"

	pb "github.com/notobject/sessiond/api/rpc"
	"github.com/notobject/sessiond/pkg/dao"
)

// GetSessionInfoByID 通过sessionID获取会话信息
func (*DefaultSessionServer) GetSessionInfoByID(ctx context.Context, req *pb.GetSessionInfoByIDReq) (*pb.GetSessionInfoRsp, error) {

	log.Println("GetSessionInfoByID Called...")
	client, err := dao.NewDao().GetRedisClient()
	if err != nil {
		log.Fatal(err)
	}
	userInfo, err := client.HGet("vdc:services:session:session_info_mapping", req.GetSessionID()).Result()
	if err != nil {
		log.Fatal(err)
	}
	userName, err := client.HGet("vdc:services:session:session_user_mapping", req.GetSessionID()).Result()
	if err != nil {
		log.Fatal(err)
	}
	return &pb.GetSessionInfoRsp{
		Ret:       0,
		SessionID: req.GetSessionID(),
		UserName:  userName,
		UserInfo:  userInfo}, nil
}

// GetSessionInfoByUserName 通过用户名获取会话信息
func (*DefaultSessionServer) GetSessionInfoByUserName(ctx context.Context, req *pb.GetSessionInfoByUserNameReq) (*pb.GetSessionInfoRsp, error) {

	log.Println("GetSessionInfoByUserName Called...")
	client, err := dao.NewDao().GetRedisClient()
	if err != nil {
		log.Fatal(err)
	}
	sessionID, err := client.HGet("vdc:services:session:user_session_mapping", req.GetUserName()).Result()
	if err != nil {
		log.Fatal(err)
	}
	userInfo, err := client.HGet("vdc:services:session:session_info_mapping", sessionID).Result()
	if err != nil {
		log.Fatal(err)
	}
	return &pb.GetSessionInfoRsp{
		Ret:       0,
		SessionID: sessionID,
		UserName:  req.GetUserName(),
		UserInfo:  userInfo}, nil
}
