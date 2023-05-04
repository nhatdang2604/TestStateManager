package daos

type TestStateDao struct{}

func NewTestStateDAO() TestStateDao {
	dao := TestStateDao{}
	return dao
}

//create the test attemp, with the param is the map: userId => testId[]
func (d *TestStateDao) CreateTestAttempts(userTestAttemptMap map[int]([]int)) ([]int, error) {
	return []int{1, 2, 3}, nil
}
