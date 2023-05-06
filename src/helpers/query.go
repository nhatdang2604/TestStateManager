package helpers

import "strings"

func GetBatchInsertQuery(
	tableName string,
	columnName []string,
	insertData [](map[string]interface{}),
) string {

	var stringBuilder strings.Builder
	stringBuilder.WriteString("INSERT INTO " + tableName + "(" + strings.Join(columnName, ", ") + ") VALUES ")

	var questionArray []string
	var questionString string
	var questionStringArray []string
	var sql string

	//Build []string{"?", "?", "?", ....}
	for _, _ = range columnName {
		questionArray = append(questionArray, "?")
	}

	//Build "(?, ?, ?, ...., ?)"
	questionString = "(" + strings.Join(questionArray, ", ") + ")"

	//Build []string{"(?, ?, ?, ...., ?)", "(?, ?, ?, ...., ?)", "(?, ?, ?, ...., ?)", ..., "(?, ?, ?, ...., ?)"}
	for _, _ = range insertData {
		questionStringArray = append(questionStringArray, questionString)
	}

	//Build "(?, ?, ?, ...., ?), (?, ?, ?, ...., ?), (?, ?, ?, ...., ?), ... (?, ?, ?, ...., ?)"

	//Build "INSERT INTO tableName(col1, col2,...,coln) VALUES (?, ?, ?, ...., ?), (?, ?, ?, ...., ?), (?, ?, ?, ...., ?), ... (?, ?, ?, ...., ?)"
	stringBuilder.WriteString(strings.Join(questionStringArray, ", "))
	sql = stringBuilder.String()

	return sql
}
