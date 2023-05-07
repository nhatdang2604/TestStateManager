package constants

import "github.com/nhatdang2604/TestStateManager/src/datatypes"

const (
	MARIADB_TIMESTAMP_FORMAT = "2006-01-02 15:04:05"
)

var (
	READING_TEST_DURATION = datatypes.RemainTime{
		Hour:   1,
		Minute: 0,
		Second: 0,
	}
)
