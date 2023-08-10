package main

import (
	"context"
	"fmt"
	"os"

	"celestia-node/api/rpc/client"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	url := "127.0.0.1:26658"

	token := os.Getenv("CELESTIA_NODE_AUTH_TOKEN")

	var (
		rpcClient *client.Client
		err       error
	)

	if err != nil {
		fmt.Println(err)
	}

	rpcClient, err = client.NewClient(ctx, "http://"+url, token)
	if err != nil {
		fmt.Println(err)
	}

	got, err := rpcClient.State.Balance(ctx)

	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
	}

	fmt.Println(got)

}

type address struct {
	name           string
	addressString  string
	addressFromStr func(string) (interface{}, error)
	marshalJSON    func(interface{}) ([]byte, error)
	unmarshalJSON  func([]byte) (interface{}, error)
}
