package grpccontent

import (
	"context"
	"errors"

	"github.com/PolyAbit/content/internal/lib/converter"
	"github.com/PolyAbit/content/internal/lib/validators"
	"github.com/PolyAbit/content/internal/models"
	"github.com/PolyAbit/content/internal/services/content"
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
	CreateDirection(ctx context.Context, code string, name string, exams string, description string) error
	GetDirections(ctx context.Context) ([]models.Direction, error)
	DeleteDirection(ctx context.Context, directionId int64) error
}

func Register(gRPCServer *grpc.Server, content Content) {
	contentv1.RegisterContentServer(gRPCServer, &serverAPI{content: content})
}

func (s *serverAPI) CreateDirection(ctx context.Context, in *contentv1.CreateDirectionRequest) (*contentv1.Empty, error) {
	if err := validators.ValidateCreateDirection(in); err != nil {
		return &contentv1.Empty{}, err
	}

	err := s.content.CreateDirection(ctx, in.GetCode(), in.GetName(), in.GetExams(), in.GetDescription())

	if errors.Is(err, content.ErrCodeAlreadyUsed) {
		return &contentv1.Empty{}, status.Error(codes.AlreadyExists, "direction with same code already exists")
	}
	if err != nil {
		return &contentv1.Empty{}, status.Error(codes.Internal, "internal error")
	}

	return &contentv1.Empty{}, nil
}

func (s *serverAPI) GetDirections(ctx context.Context, in *contentv1.Empty) (*contentv1.Directions, error) {
	directions, err := s.content.GetDirections(ctx)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get directions")
	}

	grpcDirections := make([]*contentv1.Direction, len(directions))

	for _, dir := range directions {
		grpcDirections = append(grpcDirections, converter.FromDirectionModelToResponse(dir))
	}

	return &contentv1.Directions{
		Directions: grpcDirections,
	}, nil
}

func (s *serverAPI) DeleteDirection(ctx context.Context, in *contentv1.DeleteDirectionRequest) (*contentv1.Empty, error) {
	err := s.content.DeleteDirection(ctx, in.GetDirectionId())

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete direction")
	}

	return &contentv1.Empty{}, nil
}
