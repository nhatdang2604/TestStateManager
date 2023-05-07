package services

import (
	"log"
	"sync"

	"github.com/nhatdang2604/TestStateManager/src/constants"
	"github.com/nhatdang2604/TestStateManager/src/daos"
	"github.com/nhatdang2604/TestStateManager/src/datatypes"
	"github.com/nhatdang2604/TestStateManager/src/helpers"
)

type TestStateService struct {
	TestStateDao             daos.TestStateDao
	Mutex                    sync.Mutex
	InprogressTestAttemptMap map[int]interface{}
}

func NewTestStateService() *TestStateService {
	var service *TestStateService = &TestStateService{
		InprogressTestAttemptMap: map[int]interface{}{},
		TestStateDao:             daos.NewTestStateDAO(),
	}

	return service
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
	go timer.CountDown()
	defer s.Mutex.Unlock()

	return testAttemptId, err
}

func (s *TestStateService) GetRemainTimeOfInprogressTest(testAttemptId int32) int {

	timer, ok := (s.InprogressTestAttemptMap[int(testAttemptId)]).(*helpers.Timer)
	if !ok {
		return -1
	}

	var remainTime datatypes.RemainTime = timer.RemainTime
	var remainTimeInSecond int = remainTime.Hour*3600 + remainTime.Minute*60 + remainTime.Second

	log.Printf("%v", remainTimeInSecond)
	return remainTimeInSecond
}
