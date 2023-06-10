package services

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nhatdang2604/TestStateManager/src/constants"
	"github.com/nhatdang2604/TestStateManager/src/daos"
	"github.com/nhatdang2604/TestStateManager/src/datatypes"
	"github.com/nhatdang2604/TestStateManager/src/observations/observables"
	"github.com/nhatdang2604/TestStateManager/src/observations/observers"
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

func (s *TestStateService) AfterStopTime() {
	fmt.Println("The time has been stopped")
}

func (s *TestStateService) StartTest(userId int, testId int) (int, error) {

	userTestMap := map[int]([]int){
		userId: []int{testId},
	}
	testAttemptIds, err := s.TestStateDao.CreateTestAttempts(userTestMap)

	if nil != err {
		log.Printf("Service: Error while creating test attemp with user's id = %v and test's id = %v", userId, testId)
		return -1, err
	}

	var testAttemptId int = testAttemptIds[0]

	//Start the countdown timer
	s.Mutex.Lock()
	s.InprogressTestAttemptMap[testAttemptId] = observables.NewTimer(constants.READING_TEST_DURATION, []observers.Time{s})

	//Casting the value in the map into observables.Timer
	timer := (s.InprogressTestAttemptMap[testAttemptId]).(*observables.Timer)
	go timer.CountDown()
	defer s.Mutex.Unlock()

	return testAttemptId, err
}

func (s *TestStateService) GetRemainTimeOfInprogressTest(testAttemptId int32) int {

	timer, ok := (s.InprogressTestAttemptMap[int(testAttemptId)]).(*observables.Timer)
	if !ok {
		return -1
	}

	var remainTime datatypes.RemainTime = timer.RemainTime
	var remainTimeInSecond int = remainTime.Hour*3600 + remainTime.Minute*60 + remainTime.Second

	log.Printf("%v", remainTimeInSecond)
	return remainTimeInSecond
}

func (s *TestStateService) EndTest(testAttemptId int32) (errorCode int, err error) {

	errorCode = 0
	var testAttemptIds []interface{} = []interface{}{testAttemptId}
	testAttempts, err := s.TestStateDao.FindTestAttemptByIds(testAttemptIds, true)

	if nil != err {
		log.Printf("Error on checking if the test attempt is exists")
		errorCode = 1
		return errorCode, err
	}

	if 0 == len(testAttempts) {
		log.Printf("The test attempt with id = %v is not exists", testAttemptId)
		errorCode = 2
		return errorCode, err
	}

	//Update the test attempt's data after ending it
	var idColumnName string = "test_attempt_id"
	var endTestAttempt = map[string]interface{}{
		idColumnName: testAttempts[0][idColumnName],
		"state":      constants.TEST_STATE_END,
		"ended_at":   time.Now().Format(constants.MARIADB_TIMESTAMP_FORMAT),
	}

	//Execute the query from database to terminate the test attempt
	errorCode, err = s.TestStateDao.UpdateTestAttempt(endTestAttempt, idColumnName)
	if nil != err {
		log.Printf("Error on ending the test with id = %v", testAttemptId)
		return errorCode, err
	}

	//Terminate the timer, and delete the timer from our map
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	container, ok := s.InprogressTestAttemptMap[int(testAttemptId)]
	if !ok {
		var msg string = fmt.Sprintf("The test attempt wit id = %v is already ended", testAttemptId)
		errorCode = 3
		err = errors.New(msg)
		log.Println(msg)
		return errorCode, err
	}

	var timer *observables.Timer = container.(*observables.Timer)
	timer.Stop()
	delete(s.InprogressTestAttemptMap, int(testAttemptId))

	return errorCode, err
}
