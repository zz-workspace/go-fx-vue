package service

import (
	"context"

	"github.com/docker/docker/client"
)

type DockerService struct {
	CLI *client.Client
}

func InitDockerCLIService() *DockerService {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)
	return &DockerService{
		CLI: cli,
	}
}
