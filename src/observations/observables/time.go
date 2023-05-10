package observables

import (
	"log"
	"math/rand"
	"time"

	"github.com/nhatdang2604/TestStateManager/src/datatypes"
	"github.com/nhatdang2604/TestStateManager/src/observations/observers"
)

type Timer struct {
	Id         int
	RemainTime datatypes.RemainTime
	IsStop     bool
	Observers  []observers.Time
}

func NewTimer(remainTime datatypes.RemainTime, observers []observers.Time) *Timer {
	var timer *Timer = &Timer{
		Id:         rand.Intn(1000),
		RemainTime: remainTime,
		IsStop:     false,
		Observers:  observers,
	}
	return timer
}

func (t *Timer) Stop() {
	t.IsStop = true
}

func (t *Timer) isStop() bool {
	var result bool = false
	if t.IsStop ||
		((0 == t.RemainTime.Hour) &&
			(0 == t.RemainTime.Minute) &&
			(0 == t.RemainTime.Second)) {

		result = true
	}

	return result
}

func (t *Timer) ReduceOneSecond() {

	if t.RemainTime.Second > 0 {
		t.RemainTime.Second--
	} else {
		t.RemainTime.Second = 59
		if t.RemainTime.Minute > 0 {
			t.RemainTime.Minute--
		} else {
			t.RemainTime.Minute = 59
			if t.RemainTime.Hour > 0 {
				t.RemainTime.Hour--
			}
		}
	}
}

func (t *Timer) CountDown() {
	var ticker *time.Ticker = time.NewTicker(1 * time.Second)

	for {
		<-ticker.C
		t.ReduceOneSecond()
		log.Printf("Id: %v, Hour: %v, Min: %v, Second: %v \r\n", t.Id, t.RemainTime.Hour, t.RemainTime.Minute, t.RemainTime.Second)
		if t.isStop() {
			ticker.Stop()

			//Notify the observers
			for _, observer := range t.Observers {
				observer.AfterStopTime()
			}

			log.Printf("Timer with id = %v has been stopped", t.Id)
			break
		}
	}
}
