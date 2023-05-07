package helpers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nhatdang2604/TestStateManager/src/datatypes"
)

type Timer struct {
	RemainTime datatypes.RemainTime
}

func NewTimer(remainTime datatypes.RemainTime) *Timer {
	var timer *Timer = new(Timer)
	timer.RemainTime = remainTime
	return timer
}

func (t *Timer) isStop() bool {
	var result bool = false
	if (0 == t.RemainTime.Hour) &&
		(0 == t.RemainTime.Minute) &&
		(0 == t.RemainTime.Second) {

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
	id := rand.Intn(1000)
	var ticker *time.Ticker = time.NewTicker(1 * time.Second)

	for {
		<-ticker.C
		t.ReduceOneSecond()
		fmt.Printf("Id: %v, Hour: %v, Min: %v, Second: %v \r\n", id, t.RemainTime.Hour, t.RemainTime.Minute, t.RemainTime.Second)
		if t.isStop() {
			ticker.Stop()
			break
		}
	}
}
