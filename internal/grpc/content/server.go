package grpccontent

import (
	"context"

	contentv1 "github.com/PolyAbit/protos/gen/go/content"
	"google.golang.org/grpc"
)

type serverAPI struct {
	contentv1.UnimplementedContentServer
	content Content
}

type Content interface {
}

func Register(gRPCServer *grpc.Server, content Content) {
	contentv1.RegisterContentServer(gRPCServer, &serverAPI{content: content})
}

func (s *serverAPI) CreateDirection(_ context.Context, in *contentv1.CreateDirectionRequest) (*contentv1.Direction, error) {
	return &contentv1.Direction{Id: 1, Name: "матобес", Code: in.GetCode(), Exams: in.GetExams(), Description: in.GetDescription()}, nil
}
