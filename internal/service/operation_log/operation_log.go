package operation_log

import (
	"blog_r/internal/model"
	operationLogRepo "blog_r/internal/repository/operation_log"
	"blog_r/internal/request"
	"errors"
)

type OperationLogService struct{}

func NewOperationLogService() *OperationLogService {
	return &OperationLogService{}
}

func (s *OperationLogService) OperationLogList(req *request.LogListReq) ([]model.OperationLog, int64, error) {
	if req == nil {
		return nil, 0, errors.New("invalid request")
	}

	var filterUserID *int
	if req.UserID > 0 {
		filterUserID = &req.UserID
	}

	list, total, err := operationLogRepo.ListOperationLog(req.Page, req.PageSize, filterUserID)
	if err != nil {
		return nil, total, err
	}
	return list, total, nil
}

