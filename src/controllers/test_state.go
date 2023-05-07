package controllers

import (
	"context"

	"github.com/nhatdang2604/TestStateManager/src/protos"
	"github.com/nhatdang2604/TestStateManager/src/services"
)

type TestStateManagementController struct {
	protos.UnimplementedTestStateManagementServer
	TestStateService *services.TestStateService
}

func NewTestStateManagementController() *TestStateManagementController {
	var controller *TestStateManagementController = &TestStateManagementController{
		TestStateService: services.NewTestStateService(),
	}
	return controller
}

func (c *TestStateManagementController) GetInprogressTest(ctx context.Context, req *protos.GetInprogressTestRequest) (*protos.GetInprogressTestResponse, error) {
	return nil, nil
}

func (c *TestStateManagementController) StartTest(ctx context.Context, req *protos.StartTestRequest) (*protos.StartTestResponse, error) {

	var userId int32 = req.GetUserId()
	var testId int32 = req.GetTestId()

	testAttemptId, err := c.TestStateService.StartTest(int(userId), int(testId))

	if nil != err {
		return nil, err
	}

	var response *protos.StartTestResponse = &protos.StartTestResponse{TestAttemptId: int32(testAttemptId)}

	return response, nil
}

func (c *TestStateManagementController) EndTest(ctx context.Context, req *protos.EndTestRequest) (*protos.EndTestResponse, error) {
	return nil, nil
}
