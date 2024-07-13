package grpccontent

import (
	"context"
	"errors"
	"regexp"

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
}

func Register(gRPCServer *grpc.Server, content Content) {
	contentv1.RegisterContentServer(gRPCServer, &serverAPI{content: content})
}

func (s *serverAPI) CreateDirection(ctx context.Context, in *contentv1.CreateDirectionRequest) (*contentv1.Empty, error) {
	if err := validateCreateDirection(in); err != nil {
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

const codeRegexp = `^\d{2}\.\d{2}\.\d{2}$`

func validateCreateDirection(req *contentv1.CreateDirectionRequest) error {
	if err := validateCode(req.GetCode()); err != nil {
		return err
	}
	if req.GetName() == "" {
		return status.Error(codes.InvalidArgument, "name is required")
	}
	if req.GetExams() == "" {
		return status.Error(codes.InvalidArgument, "exams is required")
	}

	return nil
}

func validateCode(code string) error {
	if code == "" {
		return status.Error(codes.InvalidArgument, "code is required")
	}
	if !regexp.MustCompile(codeRegexp).MatchString(code) {
		return status.Error(codes.InvalidArgument, "code must be like 12.34.56")
	}
	return nil
}
