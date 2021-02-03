package session

import (
	"context"
	"testing"

	rpc "github.com/notobject/sessiond/api/rpc"
	"google.golang.org/grpc"
)

func testNewSessionClient(t *testing.T) rpc.SessionClient {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	return rpc.NewSessionClient(conn)
}

func TestCreateSession(t *testing.T) {

	client := testNewSessionClient(t)

	// 二进制会话数据
	userInfo := "asasadsda/=="
	rsp, err := client.CreateSession(context.Background(),
		&rpc.CreateSessionReq{
			UserName: "ldp",
			UserInfo: userInfo})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)

	if rsp.SessionID != "ABCDEFG1234" {
		t.Fail()
	}
}

func TestGetSession(t *testing.T) {
	client := testNewSessionClient(t)
	rsp1, err := client.GetSessionInfoByID(context.Background(), &rpc.GetSessionInfoByIDReq{
		SessionID: "ABCDEFG1234",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp1)
	if rsp1.SessionID != "S" {
		t.Fail()
	}
	if rsp1.UserName != "U" {
		t.Fail()
	}
	if rsp1.UserInfo != "info" {
		t.Fail()
	}

	rsp2, err := client.GetSessionInfoByUserName(context.Background(), &rpc.GetSessionInfoByUserNameReq{
		UserName: "ldp1",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp2)
	if rsp2.Ret != rsp1.Ret {
		t.Fail()
	}
	if rsp2.SessionID != rsp1.SessionID {
		t.Fail()
	}
	if rsp2.UserName != rsp1.UserName {
		t.Fail()
	}
	if rsp2.UserInfo != rsp1.UserInfo {
		t.Fail()
	}
}

func TestDeleteSession(t *testing.T) {
	client := testNewSessionClient(t)

	rsp, err := client.DeleteSession(context.Background(), &rpc.DeleteSessionReq{
		SessionID: "ABCDEFG1234",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)

	if rsp.Ret != 0 {
		t.Fail()
	}
}

func TestDeleteSessionBatch(t *testing.T) {
	client := testNewSessionClient(t)

	rsp, err := client.DeleteSessionBatch(context.Background(), &rpc.DeleteSessionBatchReq{
		SessionIDs: []string{"ABCDEFG1234"},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)

	if rsp.Ret != 1 {
		t.Fail()
	}
}

func TestCheckSession(t *testing.T) {
	client := testNewSessionClient(t)

	rsp, err := client.CheckSession(context.Background(),
		&rpc.CheckSessionReq{SessionID: "ABCDEFG1234"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)
	if rsp.IsValid != true {
		t.Fail()
	}
}
