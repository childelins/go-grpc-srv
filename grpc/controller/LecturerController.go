package controller

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/childelins/go-grpc-srv/global"
	"github.com/childelins/go-grpc-srv/grpc/model"
	"github.com/childelins/go-grpc-srv/proto"
)

type Lecturer struct{}

// 获取讲师列表
func (l *Lecturer) GetLecturerList(ctx context.Context, request *proto.LecturerListRequest) (*proto.LecturerListResponse, error) {
	page := 1
	limit := 10
	var name string
	companyId := 13
	if request.Page > 0 {
		page = int(request.Page)
	}
	if request.Limit > 0 {
		limit = int(request.Limit)
	}
	if len(request.Name) > 0 {
		name = request.Name
	}

	lecturerModel := &model.Lecturer{}
	lecturers, err := lecturerModel.GetList(companyId, name, page, limit)
	if err != nil {
		global.Logger.Errorf("[GetLecturerList] error: %v", err)
		return nil, status.Error(codes.Internal, "访问讲师列表异常")
	}

	var lecturerInfo []*proto.LecturerInfo
	for _, lecturer := range lecturers {
		lecturerInfo = append(lecturerInfo, &proto.LecturerInfo{
			LecturerId: int32(lecturer.LecturerId),
			Name:       lecturer.Name,
			Avatar:     lecturer.Avatar,
			Title:      lecturer.Title,
			CreatedAt:  lecturer.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	total := lecturerModel.GetCount(companyId, name)

	return &proto.LecturerListResponse{
		Total: int32(total),
		Data:  lecturerInfo,
	}, nil
}
