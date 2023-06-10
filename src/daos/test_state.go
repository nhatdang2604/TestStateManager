package daos

import (
	"database/sql"
	"log"
	"time"

	"github.com/nhatdang2604/TestStateManager/src/constants"
	"github.com/nhatdang2604/TestStateManager/src/helpers"
)

var (
	TEST_ATTEMP_TABLE_NAME = "test_attempt"
	TEST_ATTEMPT_COLUMNS   = []string{"test_attempt_id", "user_id", "original_test_id", "state", "started_at", "ended_at"}
)

type TestStateDao struct{}

func NewTestStateDAO() TestStateDao {
	dao := TestStateDao{}
	return dao
}

//create the test attemp, with the param is the map: userId => testId[]
func (d *TestStateDao) CreateTestAttempts(userTestAttemptMap map[int]([]int)) ([]int, error) {

	var tableName string = TEST_ATTEMP_TABLE_NAME
	var columnNames []string = TEST_ATTEMPT_COLUMNS[1:] //ignore the primary key column

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
				columnNames[4]: sql.NullTime{},
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
		log.Printf("Error on begining transaction: %v", err)
		return nil, err
	}

	queryResult, err := transaction.Exec(sql, bindingValues...)
	if nil != err {
		log.Printf("Error on executing query: %v", err)
		transaction.Rollback()
		return nil, err
	}

	defer transaction.Commit()
	rowEffected, err := queryResult.RowsAffected()
	if nil != err {
		log.Printf("DB doesn't support checking row effected: %v", err)
		return nil, err
	}

	lastInsertId, err := queryResult.LastInsertId()
	if nil != err {
		log.Printf("DB doesn't support checking last insert id: %v", err)
		return nil, err
	}

	var insertIds []int = make([]int, rowEffected)
	for i := int64(0); i < rowEffected; i++ {
		insertIds[i] = int(lastInsertId + i)
	}

	return insertIds, err
}

func (d *TestStateDao) UpdateTestAttempt(updateTestAttempt map[string]interface{}, idColumnName string) (errorCode int, err error) {
	errorCode = 0
	var tableName string = TEST_ATTEMP_TABLE_NAME

	sql, bindingValues := helpers.GetUpdateQuery(tableName, idColumnName, updateTestAttempt)

	var db *helpers.DB = helpers.GetDB()
	connection := db.Connection
	transaction, err := connection.Begin()
	if nil != err {
		log.Printf("Error on begining transaction: %v", err)
		errorCode = 1
		return
	}

	_, err = transaction.Exec(sql, bindingValues...)
	if nil != err {
		log.Printf("Error on executing query: %v", err)
		transaction.Rollback()
		errorCode = 2
		return
	}

	defer transaction.Commit()

	return errorCode, nil
}

func (d *TestStateDao) FindTestAttemptByIds(testAttemptIds []interface{}, isLimit bool) (testAttempts [](map[string]interface{}), err error) {
	var tableName string = TEST_ATTEMP_TABLE_NAME
	var selectColumns []string = TEST_ATTEMPT_COLUMNS
	var idColumnName = selectColumns[0]

	sql, bindingValues := helpers.GetFindByIdQuery(tableName, selectColumns, testAttemptIds, idColumnName, isLimit)

	var db *helpers.DB = helpers.GetDB()
	connection := db.Connection

	queryResult, err := connection.Query(sql, bindingValues...)
	if nil != err {
		log.Printf("Error on executing query: %v", err)
		return nil, err
	}

	testAttempts = [](map[string]interface{}){}

	//Make the buffer can be using in queryResult.Scan()
	var buffer = make([]interface{}, len(selectColumns))
	var timestampContainerSize int = 2
	var timestampContainers = make([]string, timestampContainerSize)
	var containerSize = len(selectColumns) - timestampContainerSize
	var containers = make([]interface{}, containerSize) //minux the timestamp column (started_at, ended_at), those column must be convert into string, not interface{}
	var idx int = 0
	for range containers {
		buffer[idx] = &containers[idx]
		idx++
	}
	for i := range timestampContainers {
		buffer[idx] = &timestampContainers[i]
		idx++
	}

	var rowData map[string]interface{} = map[string]interface{}{}

	for queryResult.Next() {
		if err = queryResult.Scan(buffer...); nil != err {
			log.Printf("Error on parsing query result: %v", err)
			return nil, err
		}

		for idx, columnName := range selectColumns {
			if idx < containerSize {
				rowData[columnName] = containers[idx]
			} else {
				rowData[columnName] = timestampContainers[idx-containerSize]
			}
		}

		testAttempts = append(testAttempts, rowData)
	}

	return testAttempts, err
}
