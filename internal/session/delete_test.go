package session

import (
	"context"
	"testing"

	rpc "github.com/notobject/sessiond/api/rpc"
)

func TestDeleteSession(t *testing.T) {
	client := testNewSessionClient()

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
	client := testNewSessionClient()

	rsp, err := client.DeleteSessionBatch(context.Background(), &rpc.DeleteSessionBatchReq{
		SessionIDs: []string{"ABCDEFG1234"},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)

	if rsp.Ret != 0 {
		t.Fail()
	}
}
