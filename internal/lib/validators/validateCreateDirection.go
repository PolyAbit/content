package validators

import (
	"regexp"

	contentv1 "github.com/PolyAbit/protos/gen/go/content"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const codeRegexp = `^\d{2}\.\d{2}\.\d{2}$`

func ValidateCreateDirection(req *contentv1.CreateDirectionRequest) error {
	if err := ValidateCode(req.GetCode()); err != nil {
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

func ValidateCode(code string) error {
	if code == "" {
		return status.Error(codes.InvalidArgument, "code is required")
	}
	if !regexp.MustCompile(codeRegexp).MatchString(code) {
		return status.Error(codes.InvalidArgument, "code must be like 12.34.56")
	}
	return nil
}
