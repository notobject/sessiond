package session

import (
	"context"
	"testing"

	rpc "github.com/notobject/sessiond/api/rpc"
)

func TestGetSession(t *testing.T) {
	client := testNewSessionClient()
	rsp1, err := client.GetSessionInfoByID(context.Background(), &rpc.GetSessionInfoByIDReq{
		SessionID: "ABCDEFG1234",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp1)
	if rsp1.SessionID != "ABCDEFG1234" {
		t.Fail()
	}
	if rsp1.UserName != "ldp" {
		t.Fail()
	}
	if rsp1.UserInfo != "asasadsda/==" {
		t.Fail()
	}
}

func BenchmarkGetSession(b *testing.B) {
	client := testNewSessionClient()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		client.GetSessionInfoByID(context.Background(), &rpc.GetSessionInfoByIDReq{
			SessionID: "ABCDEFG1234"})
	}

}
