package daos

import (
	"log"
	"time"

	"github.com/nhatdang2604/TestStateManager/src/constants"
	"github.com/nhatdang2604/TestStateManager/src/helpers"
)

var (
	TEST_ATTEMP_TABLE_NAME = "test_attempt"
	TEST_ATTEMPT_COLUMNS   = []string{"user_id", "test_id", "state", "started_at", "ended_at"}
)

type TestStateDao struct{}

func NewTestStateDAO() TestStateDao {
	dao := TestStateDao{}
	return dao
}

//create the test attemp, with the param is the map: userId => testId[]
func (d *TestStateDao) CreateTestAttempts(userTestAttemptMap map[int]([]int)) ([]int, error) {

	var tableName string = TEST_ATTEMP_TABLE_NAME
	var columnNames []string = TEST_ATTEMPT_COLUMNS

	var insertDataSize int = 0
	for _, testIds := range userTestAttemptMap {
		for range testIds {
			insertDataSize++
		}
	}

	insertData := make([](map[string]interface{}), insertDataSize)

	var index int = 0
	for userId, testIds := range userTestAttemptMap {
		for _, testId := range testIds {
			buffer := map[string]interface{}{
				columnNames[0]: userId,
				columnNames[1]: testId,
				columnNames[2]: constants.TEST_STATE_START,
				columnNames[3]: time.Now().Format(constants.MARIADB_TIMESTAMP_FORMAT),
				columnNames[4]: nil,
			}

			insertData[index] = buffer
			index++
		}
	}

	var db *helpers.DB = helpers.GetDB()
	var sql, bindingValues = helpers.GetBatchInsertQuery(
		tableName,
		columnNames,
		insertData,
	)

	connection := db.Connection
	transaction, err := connection.Begin()
	if nil != err {
		log.Fatalf("Error on begining transaction: %v", err)
		return nil, err
	}

	queryResult, err := transaction.Exec(sql, bindingValues...)
	if nil != err {
		log.Fatalf("Error on executing query: %v", err)
		transaction.Rollback()
		return nil, err
	}

	defer transaction.Commit()
	rowEffected, err := queryResult.RowsAffected()
	if nil != err {
		log.Fatalf("DB doesn't support checking row effected: %v", err)
		return nil, err
	}

	lastInsertId, err := queryResult.LastInsertId()
	if nil != err {
		log.Fatalf("DB doesn't support checking last insert id: %v", err)
		return nil, err
	}

	var insertIds []int = make([]int, rowEffected)
	for i := int64(0); i < rowEffected; i++ {
		insertIds[i] = int(lastInsertId + i)
	}

	return insertIds, err
}
