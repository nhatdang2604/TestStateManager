package controllers

import (
	"context"
	"fmt"

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

	var testAttemptId int32 = req.GetTestAttemptId()
	var timeRemainInSecond = c.TestStateService.GetRemainTimeOfInprogressTest(testAttemptId)
	var response *protos.GetInprogressTestResponse = &protos.GetInprogressTestResponse{
		TimeRemainInSecond: int32(timeRemainInSecond),
	}

	return response, nil
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

	var testAttemptId int32 = req.GetTestAttemptId()

	var message string = "The test has been submit"
	errorCode, err := c.TestStateService.EndTest(testAttemptId)
	if nil != err {
		message = fmt.Sprintf("%v", err)
	}

	var response *protos.EndTestResponse = &protos.EndTestResponse{
		ErrorCode: int32(errorCode),
		Message:   message,
	}

	return response, err
}
