package session

import (
	"context"
	"testing"

	rpc "github.com/notobject/sessiond/api/rpc"
)

func TestCheckSession(t *testing.T) {
	client := testNewSessionClient()

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

func BenchmarkCheckSession(b *testing.B) {
	client := testNewSessionClient()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		client.CheckSession(context.Background(),
			&rpc.CheckSessionReq{SessionID: "ABCDEFG1234"})
	}
}
