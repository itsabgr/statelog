package api

import (
	"context"
	"github.com/itsabgr/statelog/internal/pb"
	"github.com/itsabgr/statelog/pkg/event"
	"github.com/itsabgr/statelog/pkg/node"
	"google.golang.org/grpc"
	_ "modernc.org/sqlite"
	"net"
)

//go:generate bash ../../scripts/generate-go.sh

func NewServer() *Server {
	server := grpc.NewServer()
	s := &Server{server: server}
	pb.RegisterAPIServer(server, s)
	return s
}

func (s *Server) Serve(ln net.Listener) error {
	return s.server.Serve(ln)
}

func (s *Server) Close() error {
	s.server.Stop()
	return nil
}

type Server struct {
	pb.UnimplementedAPIServer
	node   *node.Node
	server *grpc.Server
}

func (s *Server) GetEvent(ctx context.Context, input *pb.GetEventInput) (output *pb.GetEventOutput, err error) {
	ev, confirm, err := s.node.Get(ctx, input.Signer, input.Index)
	return &pb.GetEventOutput{Event: (*pb.Event)(ev), Confirm: confirm}, err
}

func (s *Server) PushEvent(ctx context.Context, input *pb.PushEventInput) (output *pb.PushEventOutput, err error) {
	ev := (*event.Event)(input.Event)
	err = s.node.Push(ctx, ev)
	if err != nil {
		return nil, err
	}
	return &pb.PushEventOutput{Sig: s.node.Sign(ev.Encode()), Signer: s.node.PK()}, nil
}
