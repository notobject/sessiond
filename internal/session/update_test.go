package session

import (
	"context"
	"testing"

	"github.com/notobject/sessiond/api/rpc"
)

func TestUpdateSession(t *testing.T) {
	client := testNewSessionClient()

	rsp, err := client.UpdateSession(context.Background(),
		&rpc.UpdateSessionReq{SessionID: "ABCDEFG1234", UserInfo: "aaa"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)
	if rsp.Ret != 0 {
		t.Fail()
	}
}

func BenchmarkUpdateSession(b *testing.B) {
	client := testNewSessionClient()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.UpdateSession(context.Background(),
			&rpc.UpdateSessionReq{SessionID: "ABCDEFG1234", UserInfo: "aaa"})
	}
}
