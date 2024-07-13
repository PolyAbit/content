package grpccontent

import (
	"context"
	"fmt"

	"github.com/PolyAbit/content/internal/models"
	contentv1 "github.com/PolyAbit/protos/gen/go/content"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	contentv1.UnimplementedContentServer
	content Content
}

type Content interface {
	CreateDirection(ctx context.Context, code string, name string, exams string, description string) (models.Direction, error)
}

func Register(gRPCServer *grpc.Server, content Content) {
	contentv1.RegisterContentServer(gRPCServer, &serverAPI{content: content})
}

func (s *serverAPI) CreateDirection(ctx context.Context, in *contentv1.CreateDirectionRequest) (*contentv1.Direction, error) {
	_, err := s.content.CreateDirection(ctx, in.GetCode(), in.GetName(), in.GetExams(), in.GetDescription())

	fmt.Println(err)

	if err != nil {
		return &contentv1.Direction{}, status.Error(codes.InvalidArgument, "invalid email or password")
	}

	return &contentv1.Direction{Id: 1, Name: "матобес", Code: in.GetCode(), Exams: in.GetExams(), Description: in.GetDescription()}, nil
}
