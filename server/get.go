package main

import (
	"context"
	"fmt"

	"celestia-node/api/rpc/client"

	"github.com/celestiaorg/celestia-node/share"
)

func Get(name string, blockheight uint64) []byte {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	url := "127.0.0.1:26658"

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJwdWJsaWMiLCJyZWFkIiwid3JpdGUiLCJhZG1pbiJdfQ.rQUXHwVHoTT0rZL9ONaZSF6VBzJtoVR1MLMe5Z6grr8"

	var (
		rpcClient *client.Client
		err       error
	)

	namespace := sha256Bytes_to10bytes(name)

	namespace, err = share.NewBlobNamespaceV0(namespace)
	if err != nil {
		panic(err)
	}

	rpcClient, err = client.NewClient(ctx, "http://"+url, token)
	if err != nil {
		fmt.Println(err)
	}

	na := []share.Namespace{}
	na = append(na, namespace)

	got, err := rpcClient.Blob.GetAll(ctx, blockheight, na)
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
	}

	return got[0].Blob.Data
}

type get_res struct {
	Blob string
}
