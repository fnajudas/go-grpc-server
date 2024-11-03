package service

import (
	"context"
	"fmt"
	student "grpc-server/proto/students"
)

// StudentService implementasi dari student.StudentServiceServer
type StudentService struct {
	student.UnimplementedStudentServiceServer
}

// Student adalah handler untuk request StudentReq
func (s *StudentService) Student(ctx context.Context, req *student.StudentReq) (*student.StudentResp, error) {
	// Simulasi pencarian data berdasarkan NIK
	// Sebaiknya ini dihubungkan dengan usecase atau service lainnya
	if req.Nik == "12345" {
		return &student.StudentResp{
			Name: "John Doe",
			Age:  21,
		}, nil
	}

	return nil, fmt.Errorf("student not found")
}
