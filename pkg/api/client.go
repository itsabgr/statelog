package api

import (
	"context"
	"github.com/itsabgr/statelog/internal/pb"
	"google.golang.org/grpc"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) GetEvent(server string, ctx context.Context, in *pb.GetEventInput) (*pb.GetEventOutput, error) {
	//TODO optimize with pool
	conn, err := grpc.Dial(server)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	cli := pb.NewAPIClient(conn)
	return cli.GetEvent(ctx, in)
}

func (c *Client) PushEvent(server string, ctx context.Context, in *pb.PushEventInput) (*pb.PushEventOutput, error) {
	//TODO optimize with pool
	conn, err := grpc.Dial(server)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	cli := pb.NewAPIClient(conn)
	return cli.PushEvent(ctx, in)
}
