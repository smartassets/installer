package util

import (
	"context"
	"github.com/docker/docker/client"
)

func CheckDocker()error {
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	ctx := context.TODO()
	_, err = cli.Info(ctx)
	if err != nil {
		return err
	}

	return nil
}
