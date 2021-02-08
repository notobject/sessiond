package session

import (
	"context"
	"testing"

	rpc "github.com/notobject/sessiond/api/rpc"
)

func createSession(client rpc.SessionClient) (*rpc.CreateSessionRsp, error) {
	// 二进制会话数据
	userInfo := "asasadsda/=="
	rsp, err := client.CreateSession(context.Background(),
		&rpc.CreateSessionReq{
			UserName: "ldp",
			UserInfo: userInfo})
	return rsp, err
}

func TestCreateSession(t *testing.T) {

	client := testNewSessionClient()
	rsp, err := createSession(client)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)

	if rsp.SessionID != "ABCDEFG1234" {
		t.Fail()
	}
}

func BenchmarkCreateSession(b *testing.B) {
	client := testNewSessionClient()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createSession(client)
	}
}
