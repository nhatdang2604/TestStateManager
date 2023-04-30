package controllers

import (
	"context"

	"github.com/nhatdang2604/TestStateManager/src/protos"
)

type TestStateManagementController struct {
	protos.UnimplementedTestStateManagementServer
}

func (controller *TestStateManagementController) GetInprogressTest(ctx context.Context, req *protos.GetInprogressTestRequest) (*protos.GetInprogressTestResponse, error) {
	return nil, nil
}

func (controller *TestStateManagementController) StartTest(ctx context.Context, req *protos.StartTestRequest) (*protos.StartTestResponse, error) {
	return nil, nil
}

func (controller *TestStateManagementController) EndTest(ctx context.Context, req *protos.EndTestRequest) (*protos.EndTestResponse, error) {
	return nil, nil
}
