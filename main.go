package main

import (
	ledgerPB "MySuperLedger/client/src/autogen/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/ini.v1"
	"os"
)

func main() {
	cfg, err := ini.Load("config/dev.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	serverAddr := cfg.Section("server").Key("cluster").String()
	fmt.Println("Cluster:", serverAddr)

	request := &ledgerPB.ConfigureAccountMetadata_Request{
		Metadata: &ledgerPB.AccountMetadata{
			Version:       1,
			Type:          ledgerPB.AccountType_Asset,
			LowInclusive:  10,
			HighInclusive: 2000,
		},
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Fail create connection to server: %v", err)
		os.Exit(1)
	}
	client := ledgerPB.NewLedgerServiceClient(conn)
	response, err := client.ConfigureAccountMetadata(context.Background(), request)
	if err != nil {
		fmt.Printf("Fail to send grpc request: %v", err)
		os.Exit(1)
	}
	fmt.Println("response:", response.Message)
	defer conn.Close()
}
