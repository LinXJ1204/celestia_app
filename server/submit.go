package main

import (
	"context"
	"crypto/sha256"
	"fmt"

	"celestia-node/api/rpc/client"

	"github.com/celestiaorg/celestia-node/blob"
	"github.com/celestiaorg/celestia-node/share"

	"cosmossdk.io/math"
)

func Submit(arr []byte, name string) interface{} {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	url := "127.0.0.1:26658"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJwdWJsaWMiLCJyZWFkIiwid3JpdGUiLCJhZG1pbiJdfQ.luZsE4xYLt1nSzpWZvlrdLdVc53cM-ub41NcVuNlXWU"
	fmt.Println(token)

	var (
		rpcClient *client.Client
		err       error
		blob_arr  []*blob.Blob
		res       Submit_res
	)

	fee := math.NewIntFromUint64(500000)

	namespace := sha256Bytes_to10bytes(name)
	fmt.Println("namespace")
	fmt.Println((namespace))

	namespace, err = share.NewBlobNamespaceV0(namespace)
	if err != nil {
		panic(err)
	}

	generatedBlob, err := blob.NewBlobV0(namespace, arr)
	if err != nil {
		panic(err)
	}

	rpcClient, err = client.NewClient(ctx, "http://"+url, token)
	if err != nil {
		fmt.Println(err)
	}

	blob_arr = append(blob_arr, generatedBlob)

	got, err := rpcClient.State.SubmitPayForBlob(ctx, fee, 1000000, blob_arr)
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
	}

	fmt.Println(got)
	res.Blockheight = int(got.Height)
	res.Tx_hash = got.TxHash
	return res
}

func sha256Bytes_to10bytes(data string) []byte {
	hash := sha256.Sum256([]byte(data))
	re := make([]byte, 10)
	copy(re, hash[:])
	return re
}

type Submit_res struct {
	Blockheight int    `json:"blockheight"`
	Tx_hash     string `json:"tx_hash"`
}
