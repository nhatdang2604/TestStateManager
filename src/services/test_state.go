package services

import (
	"log"
	"sync"

	"github.com/nhatdang2604/TestStateManager/src/constants"
	"github.com/nhatdang2604/TestStateManager/src/daos"
	"github.com/nhatdang2604/TestStateManager/src/helpers"
)

type TestStateService struct {
	TestStateDao             daos.TestStateDao
	Mutex                    sync.Mutex
	InprogressTestAttemptMap map[int]interface{}
}

func (s *TestStateService) StartTest(userId int, testId int) (int, error) {

	userTestMap := map[int]([]int){
		userId: []int{testId},
	}
	testAttemptIds, err := s.TestStateDao.CreateTestAttempts(userTestMap)

	if nil != err {
		log.Fatalf("Service: Error while creating test attemp with user's id = %v and test's id = %v", userId, testId)
		return -1, err
	}

	var testAttemptId int = testAttemptIds[0]

	//Start the countdown timer
	s.Mutex.Lock()
	s.InprogressTestAttemptMap[testAttemptId] = helpers.NewTimer(constants.READING_TEST_DURATION)

	//Casting the value in the map into helpers.Timer
	timer := (s.InprogressTestAttemptMap[testAttemptId]).(*helpers.Timer)
	channel := make(chan helpers.RemainTime)
	go timer.CountDown(channel)
	<-channel
	defer s.Mutex.Unlock()

	return testAttemptId, err
}
