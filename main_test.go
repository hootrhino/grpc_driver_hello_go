package main

import (
	"context"
	"grpc_driver_hello_go/sidecar"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestRpc(t *testing.T) {
	go StartServer()
	conn, err := grpc.Dial("127.0.0.1:8899",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	client := sidecar.NewSidecarClient(conn)

	resp1, err1 := client.Init(context.Background(), &sidecar.Config{})
	if err1 != nil {
		t.Fatal(err1)
	}
	t.Log(resp1)
	resp2, err2 := client.Start(context.Background(), &sidecar.Request{})
	if err2 != nil {
		t.Fatal(err2)
	}
	t.Log(resp2)

	resp3, err3 := client.Status(context.Background(), &sidecar.Request{})
	if err3 != nil {
		t.Fatal(err3)
	}
	t.Log(resp3)

	resp4, err4 := client.Read(context.Background(), &sidecar.ReadRequest{})
	if err4 != nil {
		t.Fatal(err4)
	}
	t.Log(resp4)

	resp5, err5 := client.Write(context.Background(), &sidecar.WriteRequest{})
	if err5 != nil {
		t.Fatal(err5)
	}
	t.Log(resp5)

	resp6, err6 := client.Stop(context.Background(), &sidecar.Request{})
	if err6 != nil {
		t.Fatal(err6)
	}
	t.Log(resp6)

}
